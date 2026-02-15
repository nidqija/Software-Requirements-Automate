package main

import (
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)



func storageHandler(w http.ResponseWriter, r *http.Request) {
   cookie , err := r.Cookie("client_id")

   if err != nil {
	  newID := uuid.New().String() // Generate a random ID for the client

	  newCookie := &http.Cookie{
		
		 Name:  "client_id", // set the cookie name
		 Value: newID, // set the cookie value to the generated ID
		 Path :"/", // set the path to root so it's accessible across the entire site
		 HttpOnly: true, // make it inaccessible to javascript
		 MaxAge: 86400 * 365 * 10, // set it to 10 years in seconds method
		 SameSite: http.SameSiteLaxMode, // adjust samesite policy as needed
   }

   // set the cookie in response header
   http.SetCookie(w, newCookie)

   // log the new client connection with the assigned ID 
   log.Printf("New client connected, assigned ID: %s", newID)
   w.Write([]byte("New Client ID generated and stored in cookie."))

   // return a welcome message to the client with the id
   return
}

// if cookie exists , read the value and return to the client side
w.Write([]byte("Welcome back! Your ID is : " + cookie.Value))

}

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
        }).Bind(apis.CORS(apis.CORSConfig{
          
			AllowOrigins:     []string{"http://localhost:5173"}, // NO WILDCARD
			AllowCredentials: true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*"},
		
		})) // Keep CORS if you're calling this from a frontend on a different port

        // Serve static files from pb_public
        // Note: Ensure the pb_public directory exists in your project root
        se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

        return se.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}