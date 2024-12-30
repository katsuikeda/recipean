package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/katsuikeda/recipean/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db       *database.Queries
	platform string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM environment variable is not set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database connection: %v\n", err)
	}
	defer dbConn.Close()
	dbQueries := database.New(dbConn)

	apiCfg := &apiConfig{
		db:       dbQueries,
		platform: platform,
	}

	addr := ":" + port

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/healthz", handlerReadiness)

	mux.HandleFunc("POST /api/v1/users", apiCfg.handlerCreateUser)

	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 10,
	}

	fmt.Printf("Listening to port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
