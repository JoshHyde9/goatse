package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"goatse/middleware"
	"goatse/routers"
)

func main() {
	r := chi.NewRouter()

	postRouter := routers.PostRouter()

	r.Use(middleware.SetJSONContentType)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Hello, World!",
		})
	})

	r.Mount("/posts", postRouter)

	log.Println("Starting server on :5000")
	http.ListenAndServe(":5000", r)
}
