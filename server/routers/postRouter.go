package routers

import (
	"encoding/json"
	"net/http"
	"slices"
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
		var newPostInput models.CreatePost
		err := json.NewDecoder(r.Body).Decode(&newPostInput)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "You dun fucked",
			})
			return
		}

		newPost := models.Post{
			Id:        uuid.NewString(),
			Title:     newPostInput.Title,
			Author:    newPostInput.Author,
			CreatedAt: time.Now(),
		}

		posts = append(posts, newPost)

		json.NewEncoder(w).Encode(posts)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Missing `id` param",
			})
		}

		found := false
		for _, post := range posts {
			if post.Id == id {
				found = true
				json.NewEncoder(w).Encode(post)
				break
			}
		}

		if !found {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Missing `id` param",
			})
		}

	})

	r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Missing `id` param",
			})
		}

		var updatePost models.UpdatePost

		err := json.NewDecoder(r.Body).Decode(&updatePost)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "You dun fucked",
			})
			return
		}

		found := false
		for i := range posts {
			if posts[i].Id == id {
				found = true

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

		if !found {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "You dun fucked",
			})
		}
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "You dun fucked",
			})
			return
		}

		index := slices.IndexFunc(posts, func(post models.Post) bool {
			return post.Id == id
		})

		if index >= 0 {
			posts = slices.Delete(posts, index, index+1)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "You dun fucked",
			})
			return
		}

		json.NewEncoder(w).Encode(posts)
	})

	return r
}
