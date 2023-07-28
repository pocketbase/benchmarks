package benchmarks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tokens"
	"github.com/pocketbase/pocketbase/tools/routine"
)

const (
	benchmarkStartedKey = "__benchmarkStarted"

	colOrganizations = "organizations"
	colPermissions   = "permissions"
	colUsers         = "users"
	colPosts10k      = "posts10k"
	colPosts25k      = "posts25k"
	colPosts50k      = "posts50k"
	colPosts100k     = "posts100k"
	colBenchmarks    = "benchmarks"
)

func MustRegister(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Server.WriteTimeout = 0

		// disable the logs to reduce the load and the required file descriptors
		app.Settings().Logs.MaxDays = 0

		// create a dummy admin if there is none
		if total, _ := app.Dao().TotalAdmins(); total == 0 {
			admin := &models.Admin{}
			admin.Email = "test@example.com"
			admin.SetPassword("1234567890")
			if err := app.Dao().SaveAdmin(admin); err != nil {
				return err
			}
		}

		// register the `GET /benchmarks` route
		e.Router.GET("/benchmarks", func(c echo.Context) error {
			if app.Cache().Has(benchmarkStartedKey) {
				return c.String(http.StatusOK, "Another benchmark is already running, please check later...")
			}
			app.Cache().Set(benchmarkStartedKey, true)

			toRunRaw := c.QueryParamDefault("run", "create,auth,search,custom,delete")

			toRun := strings.Split(toRunRaw, ",")

			buff := new(bytes.Buffer)

			r := runner{
				app:     app,
				baseUrl: "http://" + e.Server.Addr,
				writers: map[io.Writer]AfterRunFunc{
					// output the result in the console
					os.Stdout: nil,
					// output the result in a collection
					// (in case the console is not accessible or the deployment env doesn't allow persisting connections)
					buff: func(runErr error) {
						// write the result into the benchmark collection
						runResult := buff.String()
						collection, _ := app.Dao().FindCollectionByNameOrId(colBenchmarks)
						record := models.NewRecord(collection)
						record.Set("tests", toRunRaw)
						record.Set("result", runResult)
						if runErr != nil {
							record.Set("error", runErr.Error())
						}
						if err := app.Dao().SaveRecord(record); err != nil {
							log.Println("Failed to save benchmark record: ", err)
						}
					},
				},
			}

			// run tests in the background because some host providers
			// don't allow long persistence connections.
			routine.FireAndForget(func() {
				// the response was already commited, so we just log the error
				if err := r.run(toRun); err != nil {
					log.Println("Run error: ", err)
				}

				app.Cache().Remove(benchmarkStartedKey)
			})

			return c.String(http.StatusOK, "Benchmarks started - you can check the results later in the console or in the "+colBenchmarks+" collection.")
		})

		return nil
	})
}

type runWritter interface {
	io.Writer
	http.Flusher
}

type AfterRunFunc func(runErr error)

type runner struct {
	app     core.App
	baseUrl string
	writers map[io.Writer]AfterRunFunc
}

// write writes the provided string to the runner writers.
func (r *runner) write(format string, a ...any) error {
	for w := range r.writers {
		if _, err := fmt.Fprintf(w, format, a...); err != nil {
			log.Println("Write failure:", err)
		}

		// ensures that there is always a new line
		fmt.Fprintf(w, "\n")
	}

	return nil
}

// run executes the runner with the specified testNames.
func (r *runner) run(testNames []string) error {
	tests := map[string]func() error{
		"create": func() error {
			if err := r.resetSchema(true); err != nil {
				return err
			}

			if err := r.createOrganizations(); err != nil {
				return err
			}

			if err := r.createPermissions(); err != nil {
				return err
			}

			if err := r.createUsers(); err != nil {
				return err
			}

			if err := r.createPosts(); err != nil {
				return err
			}

			return nil
		},
		"auth": func() error {
			if err := r.authWithPassword(); err != nil {
				return err
			}

			if err := r.authRefresh(); err != nil {
				return err
			}

			return nil
		},
		"search": func() error {
			if err := r.listRecords(); err != nil {
				return err
			}

			return nil
		},
		"custom": func() error {
			if err := r.customRoute(); err != nil {
				return err
			}

			if err := r.customHook(); err != nil {
				return err
			}

			return nil
		},
		"delete": func() error {
			if err := r.deleteRecords(); err != nil {
				return err
			}

			return nil
		},
	}

	// run tests
	var runErr error
	for _, name := range testNames {
		name = strings.TrimSpace(name)

		exec, ok := tests[name]
		if !ok {
			return errors.New("missing benchmark test " + name)
		}

		if runErr = exec(); runErr != nil {
			break
		}
	}

	if runErr == nil {
		r.write("---------------------------------------------------")
		r.write("Completed!")
	}

	for _, afterRun := range r.writers {
		if afterRun != nil {
			afterRun(runErr)
		}
	}

	return nil
}

func (r *runner) cooldown() {
	time.Sleep(2 * time.Second)
}

func (r *runner) updateCollection(collectionIdOrName string, data map[string]any) error {
	collection, err := r.app.Dao().FindCollectionByNameOrId(collectionIdOrName)
	if err != nil {
		return err
	}

	rawData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rawData, collection); err != nil {
		return err
	}

	return r.app.Dao().SaveCollection(collection)
}

func (r *runner) totalRecords(collectionIdOrName string) (int, error) {
	collection, err := r.app.Dao().FindCollectionByNameOrId(collectionIdOrName)
	if err != nil {
		return 0, err
	}

	var total int

	if err := r.app.Dao().RecordQuery(collection).Select("count(id)").Row(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (r *runner) randomRecordIds(collectionIdOrName string, count int) ([]string, error) {
	collection, err := r.app.Dao().FindCollectionByNameOrId(collectionIdOrName)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, count)

	queryErr := r.app.Dao().RecordQuery(collection).
		Select("id").
		OrderBy("random()").
		Limit(int64(count)).
		Column(&ids)
	if queryErr != nil {
		return nil, queryErr
	}

	return ids, nil
}

func (r *runner) randomUserAuth() (*models.Record, string, error) {
	collection, err := r.app.Dao().FindCollectionByNameOrId(colUsers)
	if err != nil {
		return nil, "", err
	}

	user := &models.Record{}

	if err := r.app.Dao().RecordQuery(collection).Limit(1).One(user); err != nil {
		return nil, "", err
	}

	token, err := tokens.NewRecordAuthToken(r.app, user)

	return user, token, err
}

func (r *runner) randomAdmin() (*models.Admin, string, error) {
	admin := &models.Admin{}

	if err := r.app.Dao().AdminQuery().Limit(1).One(admin); err != nil {
		return nil, "", err
	}

	token, err := tokens.NewAdminAuthToken(r.app, admin)

	return admin, token, err
}
