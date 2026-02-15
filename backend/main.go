package main

import (
    "log"
    "net/http"
    "os"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()

    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        // Register your custom API route
        se.Router.GET("/api/test", func(e *core.RequestEvent) error {
            log.Println("test api called")

            return e.JSON(http.StatusOK, map[string]any{
                "status":  "success",
                "message": "hello! this is JSON from GO",
                "data":    12345,
            })
        }).Bind(apis.CORS(apis.CORSConfig{})) // Keep CORS if you're calling this from a frontend on a different port

        // Serve static files from pb_public
        // Note: Ensure the pb_public directory exists in your project root
        se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

        return se.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}