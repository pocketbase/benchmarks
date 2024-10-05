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

	"github.com/pocketbase/pocketbase/core"
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
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Server.WriteTimeout = 0

		// disable the logs to reduce the load and the required file descriptors
		app.Settings().Logs.MaxDays = 0

		// create a dummy superuser if there is none
		if total, _ := app.CountRecords(core.CollectionNameSuperusers); total == 0 {
			superusersCol, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
			if err != nil {
				return err
			}

			superuser := core.NewRecord(superusersCol)
			superuser.Set("email", "test@example.com")
			superuser.Set("password", "1234567890")
			err = app.Save(superuser)
			if err != nil {
				return err
			}
		}

		// register the `GET /benchmarks` route
		se.Router.GET("/benchmarks", func(e *core.RequestEvent) error {
			if e.App.Store().Has(benchmarkStartedKey) {
				return e.String(http.StatusOK, "Another benchmark is already running, please check later...")
			}
			e.App.Store().Set(benchmarkStartedKey, true)

			toRunRaw := e.Request.URL.Query().Get("run")
			if toRunRaw == "" {
				toRunRaw = "create,auth,search,custom,delete"
			}

			toRun := strings.Split(toRunRaw, ",")

			buff := new(bytes.Buffer)

			r := runner{
				app:     app,
				baseUrl: "http://" + se.Server.Addr,
				writers: map[io.Writer]AfterRunFunc{
					// output the result in the console
					os.Stdout: nil,
					// output the result in a collection
					// (in case the console is not accessible or the deployment env doesn't allow persisting connections)
					buff: func(runErr error) {
						collection, err := e.App.FindCollectionByNameOrId(colBenchmarks)
						if err != nil {
							log.Printf("Missing benchmarks collection probably due to failed schema import - %v (%v)\n", err, runErr)
							return
						}

						// write the result into the benchmark collection
						runResult := buff.String()
						record := core.NewRecord(collection)
						record.Set("tests", toRunRaw)
						record.Set("result", runResult)
						if runErr != nil {
							record.Set("error", runErr.Error())
						}
						if err := e.App.Save(record); err != nil {
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

				app.Store().Remove(benchmarkStartedKey)
			})

			return e.String(http.StatusOK, "Benchmarks started - you can check the results later in the console or in the "+colBenchmarks+" collection.")
		})

		return se.Next()
	})
}

type runWritter interface {
	io.Writer
	http.Flusher
}

type AfterRunFunc func(runErr error)

type runner struct {
	app     core.App
	writers map[io.Writer]AfterRunFunc
	baseUrl string
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
				return fmt.Errorf("resetSchema: %w", err)
			}

			if err := r.createOrganizations(); err != nil {
				return fmt.Errorf("createOrganizations: %w", err)
			}

			if err := r.createPermissions(); err != nil {
				return fmt.Errorf("createPermissions: %w", err)
			}

			if err := r.createUsers(); err != nil {
				return fmt.Errorf("createUsers: %w", err)
			}

			if err := r.createPosts(); err != nil {
				return fmt.Errorf("createPosts: %w", err)
			}

			return nil
		},
		"auth": func() error {
			if err := r.authWithPassword(); err != nil {
				return fmt.Errorf("authWithPassword: %w", err)
			}

			if err := r.authRefresh(); err != nil {
				return fmt.Errorf("authRefresh: %w", err)
			}

			return nil
		},
		"search": func() error {
			if err := r.listRecords(); err != nil {
				return fmt.Errorf("listRecords: %w", err)
			}

			return nil
		},
		"custom": func() error {
			if err := r.customRoute(); err != nil {
				return fmt.Errorf("customRoute: %w", err)
			}

			if err := r.customHook(); err != nil {
				return fmt.Errorf("customHook: %w", err)
			}

			return nil
		},
		"delete": func() error {
			if err := r.deleteRecords(); err != nil {
				return fmt.Errorf("deleteRecords: %w", err)
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

	return runErr
}

func (r *runner) cooldown() {
	time.Sleep(2 * time.Second)
}

func (r *runner) updateCollection(collectionIdOrName string, data map[string]any) error {
	collection, err := r.app.FindCollectionByNameOrId(collectionIdOrName)
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

	return r.app.Save(collection)
}

func (r *runner) randomRecordIds(collectionIdOrName string, count int) ([]string, error) {
	collection, err := r.app.FindCollectionByNameOrId(collectionIdOrName)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, count)

	queryErr := r.app.RecordQuery(collection).
		Select("id").
		OrderBy("random()").
		Limit(int64(count)).
		Column(&ids)
	if queryErr != nil {
		return nil, queryErr
	}

	return ids, nil
}

func (r *runner) randomUserAuth() (*core.Record, string, error) {
	user := &core.Record{}

	err := r.app.RecordQuery(colUsers).OrderBy("random()").Limit(1).One(user)
	if err != nil {
		return nil, "", err
	}

	token, err := user.NewAuthToken()

	return user, token, err
}

func (r *runner) randomSuperuserAuth() (*core.Record, string, error) {
	superuser := &core.Record{}

	err := r.app.RecordQuery(core.CollectionNameSuperusers).Limit(1).One(superuser)
	if err != nil {
		return nil, "", err
	}

	token, err := superuser.NewAuthToken()

	return superuser, token, err
}
