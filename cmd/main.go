package main

import (
	"log"
	"net/http"
	"scaena/pkg/db"
	"scaena/pkg/handlers"
	"scaena/pkg/logger"
	"scaena/pkg/middleware"
)

func main() {
	//directory.Setup()
	logFile := logger.SetupLogger()
	defer logFile.Close()

	log.Println("Logger initialized")

	dbPool := db.InitDB()
	defer dbPool.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/test", handlers.TestHandler(dbPool))

	fileServer := http.StripPrefix("/api/data/abcd/efgh", http.FileServer(http.Dir("./data/storage/abcd/efgh")))
	mux.Handle("/api/data/abcd/efgh/", fileServer)

	loggedMux := middleware.Logging(mux)

	corsMux := middleware.CORS(loggedMux)

	//securedMux := middleware.Auth(loggedMux)

	rateLimitedMux := middleware.RateLimit(corsMux)

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", rateLimitedMux)
	if err != nil {
		log.Fatal(err)
	}

}
