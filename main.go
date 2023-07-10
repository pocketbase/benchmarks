package main

import (
	"app/benchmarks"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
)

func customMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("total", 20)

		return next(c)
	}
}

func main() {
	app := pocketbase.New()

	jsvm.MustRegister(app, jsvm.Config{
		HooksPoolSize: 100,
	})

	benchmarks.MustRegister(app)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/go", func(c echo.Context) error {
			total, _ := c.Get("total").(int)

			records, err := app.Dao().FindRecordsByFilter("posts10k", "title != ''", "-created", total)
			if err != nil {
				return err
			}

			return c.JSON(200, records)
		}, apis.ActivityLogger(app), customMiddleware)

		return nil
	})

	app.OnRecordBeforeUpdateRequest("go").Add(func(e *core.RecordUpdateEvent) error {
		if e.Record.GetString("title") != "" {
			e.Record.Set("title", "go_update")
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
