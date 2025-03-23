package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IsaiahOden/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("hello world")
	godotenv.Load()                 // takes variables from .env file and pulls into current env
	portString := os.Getenv("PORT") // get value for variable from key
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal(("DB_URL is not found in the environment"))
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness) // connecting handlerReadiness function to /healthz path
	//v1Router.HandleFunc would handle all requests, changing it to .Get means its scope is only get requests

	v1Router.Get("/err", handlerErr) // hooking up to another handler, need to create an error handler file

	v1Router.Post("/users", apiCfg.handlerCreateUser) // look at how the method pointed to what you created here.

	router.Mount("/v1", v1Router) // nesting v1 router so full path will be /v1/healthz

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe() // code stops here and handles http requests
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString) // this almost works but you need to go into terminal and time "export PORT = 8000"
	// a package is needed to fix this "go get github.com/joho/godotenv"
	// go mod vendor added the copy to the vendor folder which holds modules.txt
	// go mod tidy if still getting import error and then go mod vendor again
	// routing is done by chi github.com/go-chi/chi/5

}

// sql done by  go install go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
