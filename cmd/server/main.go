package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/emblems/pkg/server"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("emblems"))
	router.Handle("/emblems/{name}.png", http.StripPrefix("/emblems/", fs)).Methods("GET", "OPTIONS")
	router.Handle("/emblems/{make}", server.NewHandler()).Methods("GET", "OPTIONS")

	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})

	cors := handlers.CORS(origins, methods)(router)
	srv := http.Server{
		Addr:    ":8080",
		Handler: handlers.LoggingHandler(os.Stdout, cors),
	}

	log.Println("Listening on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
