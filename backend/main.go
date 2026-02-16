package main

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

// ===================== insert new user record with sessionId (pbid) into srauto_users collection ====================  //

func insertUsers(app *pocketbase.PocketBase , pbid string) error {


            // find collection defined in pocketbase admin UI
            collection , err := app.FindCollectionByNameOrId("srauto_users")


            // if collection not found , log error and return 
            if err != nil {
                log.Printf("Error finding collection: %v", err)
                return err
            }

            // create new record in that collection
            record := core.NewRecord(collection)

            // insert the pbid data into the sessionId field of the record
            record.Set("sessionId" , pbid);

            // save the record to the database
            // if there's an error during save (like validation issues), log it and return
            if err := app.Save(record); err != nil {
              log.Printf("SAVE FAILED: %v", err) // This will show validation errors
              return err
            }

            // if save is successful, log the new record ID
            log.Printf("User record created with ID: %s", record.Id)
            return nil
        }

// ===================== insert new sequence diagram record with user prompt, sessionId and file into srauto_sequenceDiagram collection ====================  //

func insertSequenceDiagram(app *pocketbase.PocketBase , prompt string , sessionId string , multipartFile *multipart.FileHeader) error {

      collection , err := app.FindCachedCollectionByNameOrId("srauto_sequenceDiagram")

      if err != nil {
        log.Printf("Error finding collection: %v", err)
        return err
      }

      // create a new file object from provided file path using pocketbase's filesystem library
      file , err := filesystem.NewFileFromMultipart(multipartFile)

      if err != nil {
        log.Printf("Error creating file from multipart: %v", err)
        return err
      }

    
      record := core.NewRecord(collection)

      record.Set("user_prompt" , prompt)
      record.Set("sessionID" , sessionId)
      record.Set("sequence_diagram" , file)

      if err := app.Save(record); err != nil {
        log.Printf("SAVE FAILED: %v", err) 
        return err
      }
      

      log.Printf("Diagram record created for ID: %s" , record.Id)
      return nil


}

// =============== helper to ensure we don't create duplicates if user already has a cookie and record ============= //

func ensureUserExists(app *pocketbase.PocketBase, pbid string) error {

    // find the collection defined in the pocketbase admin UI
    collection, err := app.FindCollectionByNameOrId("srauto_users")

    // if not found , log error and return 
    if err != nil {
        return err
    }

    // Check if sessionId already exists
    existing, _ := app.FindFirstRecordByData("srauto_users", "sessionId", pbid)
    if existing != nil {
        log.Printf("User %s already exists in DB, skipping insert.", pbid)
        return nil // Already there, nothing to do
    }

    // if not found , create a new record and save it
    record := core.NewRecord(collection)
    record.Set("sessionId", pbid)
    
    // log the new session ID before saving, so we can trace it in logs
    // Save the new record to the database
    log.Printf("Inserting new session to DB: %s", pbid)
    return app.Save(record)
}

// ============================================================================================================= //  

func main() {

        // create new pocketbase app instance
        app := pocketbase.New()

        // bind a function to the onserve event , which runs every time the server starts and is ready to handle requests
        app.OnServe().BindFunc(func(se *core.ServeEvent) error {

        // define a new GET endpoint at /api/test that will handle incoming requests from the frontend
        se.Router.GET("/api/test", func(e *core.RequestEvent) error {
        
        // check if the incoming request has a cookie named "client_id"
        cookie, err := e.Request.Cookie("client_id")

        // if cookie is not found (err != nil), we will create a new unique ID for the user and set it as a cookie in the response. 
        var currentID string

        if err != nil {
            // if a new user , create a new unique ID using uuid package and set it as a cookie in the response
            currentID = uuid.New().String()
            newCookie := &http.Cookie{
                Name:     "client_id",
                Value:    currentID,
                Path:     "/",
                HttpOnly: true, 
                MaxAge:   86400 * 365,
                SameSite: http.SameSiteLaxMode,
                Secure:   false, 
            }

            // Set the cookie in the response so the client can store it for future requests
            http.SetCookie(e.Response, newCookie)

            // insert the new record into the database with the new session ID (currentID)
             if err := insertUsers(app, currentID); err != nil {

                // if there's an error inserting the user into the database, log it
                log.Printf("Error inserting user: %v", err)

            }
            
            // log the new client ID for debugging purposes
            log.Printf("New client assigned ID: %s", currentID)
           
        } else {
            // if cookie is found, use the existing value as the current ID and log it for debugging purposes
            currentID = cookie.Value
            log.Printf("Client returned with ID: %s", currentID)

        }

        // ensure that the user record exists through a helper function
        if err := ensureUserExists(app, currentID); err != nil {
        log.Printf("DB Error: %v", err)

      }

        // return a json response to the client with the current session ID and success status
        return e.JSON(http.StatusOK, map[string]any{
            "status":  "success",
            "pbid":    currentID, // Vue will look for 'pbid'
            "message": "Connected to PocketBase",
        })

      // bind CORS middleware to allow requests from the frontend running on localhost:5173 with credentials (cookies) included
    }).Bind(apis.CORS(apis.CORSConfig{
        // Allow requests from the frontend development server
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowCredentials: true,
    }))


    se.Router.POST("/api/submit-diagram", func(e *core.RequestEvent) error {
        
        // check for the client_id cookie to identify the user session
        cookie , err := e.Request.Cookie("client_id")

        // if cookie is not found , log the error
        if err != nil {
            log.Printf("No client_id cookie found: %v", err)

        }

        // if cookie is found , extract the session ID from cookie value
        sessionId := cookie.Value


        // extract the user prompt 
        prompt := e.Request.FormValue("prompt")
        // extract the uploading file from multipart form 
        _, file, err := e.Request.FormFile("diagram")


        // if there's an error reading the file , log the error
        if err!= nil {
            log.Printf("Error reading file from request: %v", err)
            return e.BadRequestError("Failed to read file from request", err)
         }
        

        
       // launch the function record to database
       if err := insertSequenceDiagram(app, prompt, sessionId, file); err != nil {

        // if error , log the error and return internal server error to client
        log.Printf("Error inserting sequence diagram: %v", err)
        return e.InternalServerError("Failed to save sequence diagram", err)
       }

       // if successful , return a json response to client
       return e.JSON(http.StatusOK, map[string]any{
        "status": "success",
        "message": "Sequence diagram saved successfully",
       })

       // bind cors middleware to allow requests from frontend
    }).Bind(apis.CORS(apis.CORSConfig{
        AllowOrigins: []string{"http://localhost:5173"},
        AllowCredentials: true,
    }))

    // after setting up route and middleware , start the PocketBase server and log any errors that occur during startup
    return se.Next()
})

        // if the pocketbase server fails to start, log the error and exit the application
            if err := app.Start(); err != nil {
                log.Fatal(err)
            }
        }