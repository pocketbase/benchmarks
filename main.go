package main

import (
	"app/benchmarks"
	"log"
	"net/http"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
)

func customMiddleware(e *core.RequestEvent) error {
	e.Set("total", 20)

	return e.Next()
}

var dbConnect core.DBConnectFunc

func main() {
	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultQueryTimeout: 120 * time.Second,
		DBConnect:           dbConnect,
	})

	jsvm.MustRegister(app, jsvm.Config{
		HooksPoolSize: 50,
	})

	benchmarks.MustRegister(app)

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/go", func(e *core.RequestEvent) error {
			total, _ := e.Get("total").(int)

			records, err := e.App.FindRecordsByFilter("posts10k", "title != ''", "-created", total, 0)
			if err != nil {
				return err
			}

			return e.JSON(http.StatusOK, records)
		}).BindFunc(customMiddleware)

		return se.Next()
	})

	app.OnRecordUpdateRequest("go").BindFunc(func(e *core.RecordRequestEvent) error {
		if e.Record.GetString("title") != "" {
			e.Record.Set("title", "go_update")
		}

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
