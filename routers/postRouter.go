package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"goatse/models"
)

func PostRouter() chi.Router {
	r := chi.NewRouter()

	posts := []models.Post{
		{Id: uuid.NewString(), Title: "Post 1", Author: "Your Mum", CreatedAt: time.Now()},
		{Id: uuid.NewString(), Title: "Post 2", Author: "Your Dad", CreatedAt: time.Now()},
		{Id: uuid.NewString(), Title: "Post 3", Author: "Your Fat Hog", CreatedAt: time.Now()},
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(posts)
	})

	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		var newPost models.Post
		err := json.NewDecoder(r.Body).Decode(&newPost)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "You dun fucked",
			})
			return
		}

		posts := append(posts, newPost)

		json.NewEncoder(w).Encode(posts)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Missing `id` param",
			})
		}

		for _, post := range posts {
			if post.Id == id {
				json.NewEncoder(w).Encode(post)
			}
		}
	})

	r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Missing `id` param",
			})
		}

		var updatePost models.UpdatePost

		err := json.NewDecoder(r.Body).Decode(&updatePost)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "You dun fucked",
			})
			return
		}

		for i := range posts {
			if posts[i].Id == id {

				if updatePost.Title != nil {
					posts[i].Title = *updatePost.Title
				}

				if updatePost.Author != nil {
					posts[i].Author = *updatePost.Author
				}
				json.NewEncoder(w).Encode(posts[i])
				break
			}

		}
	})

	return r
}
