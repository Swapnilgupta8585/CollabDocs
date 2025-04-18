package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Swapnilgupta8585/CollabDocs/config"
	"github.com/Swapnilgupta8585/CollabDocs/handlers"
	"github.com/Swapnilgupta8585/CollabDocs/internal/database"
	"github.com/Swapnilgupta8585/CollabDocs/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}

	// Get database URL from environment variable
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set")
	}

	// Connect to database
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer dbConn.Close()

	// Test database connection
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Connected to database")

	// create database query
	dbQueries := database.New(dbConn)

	// create apiconfig to handle state of the program
	cfg := config.ApiConfig{
		Db: dbQueries,
		SecretToken:os.Getenv("SECRET_TOKEN"),
	}

	// Create a multiplexer
	mux := http.NewServeMux()

	//
	handler := handlers.Handler{Cfg: cfg}

	// Register routes
	mux.HandleFunc("GET /api/healthz", handlers.HandleHealth)
	mux.HandleFunc("DELETE /admin/reset", handler.HandleReset)

	mux.HandleFunc("POST /api/users", handler.HandleCreateUsers)
	mux.HandleFunc("POST /api/login", handler.HandleLogin)
	mux.HandleFunc("PUT /api/users", handler.HandleUpdateUserCredentials)


	mux.HandleFunc("POST /api/refresh", handler.HandleRefresh)
	mux.HandleFunc("POST /api/revoke", handler.HandleRevoke)
	
	mux.HandleFunc("POST /api/docs", handler.HandleCreateDocs)
	mux.HandleFunc("GET /api/docs", handler.HandleGetDocsForUser)
	mux.HandleFunc("GET /api/docs/{DocID}", handler.HandleGetDocs)
	mux.HandleFunc("PUT /api/docs/{DocID}", handler.HandleUpdateDocs)
	mux.HandleFunc("DELETE /api/docs/{DocID}", handler.HandleDeleteDocs)
	mux.HandleFunc("POST /api/docs/{DocID}/share", handler.HandleDocShare)


	// add cors using cors middleware
	handlerWithCORS := middleware.CORS()(mux)

	// Create a server
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: handlerWithCORS,
	}

	// Start the server
	log.Printf("Starting server on port %s", port)
	err = server.ListenAndServe()
	log.Fatal(err)

}


