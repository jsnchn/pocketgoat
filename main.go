package main

import (
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"pocketgoat/templates"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Home route - specific route first
		se.Router.GET("/", func(e *core.RequestEvent) error {
			component := templates.Home()
			templ.Handler(component).ServeHTTP(e.Response, e.Request)
			return nil
		})

		// API routes
		se.Router.GET("/api/hello", func(e *core.RequestEvent) error {
			return e.JSON(http.StatusOK, map[string]string{"message": "Hello from PocketBase!"})
		})

		// Static files - only for paths that don't match other routes
		se.Router.GET("/static/{path...}", apis.Static(os.DirFS("./pb_public"), true))

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
