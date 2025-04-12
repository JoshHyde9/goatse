package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"goatse/db"
	"goatse/middleware"
	"goatse/routers"
)

func main() {
	r := chi.NewRouter()

	db, err := db.Initialise()

	if err != nil {
		log.Fatalf("Could not start db: %v", err)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}))
	r.Use(middleware.SetJSONContentType)

	postRouter := routers.PostRouter(db)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"code":    http.StatusOK,
			"message": "Hello, World!",
		})
	})

	r.Mount("/posts", postRouter)

	log.Println("Starting server on :5000")
	http.ListenAndServe(":5000", r)
}
