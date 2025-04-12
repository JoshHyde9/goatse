package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"goatse/models"
)

func PostRouter(db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var posts []models.Post

		db.Order("created_at desc").Find(&posts)

		json.NewEncoder(w).Encode(posts)
	})

	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		var newPostInput models.CreatePost
		err := json.NewDecoder(r.Body).Decode(&newPostInput)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
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

		if err := db.Create(newPost).Error; err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Something ain't right here",
			})
			return
		}

		json.NewEncoder(w).Encode(newPost)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		var existingPost models.Post
		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Missing `id` param",
			})
			return
		}

		result := db.First(&existingPost, "id = ?", id)

		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Something ain't right here",
			})
			return
		}
		json.NewEncoder(w).Encode(existingPost)
	})

	r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "Missing `id` param",
			})
			return
		}

		var updatePost models.UpdatePost

		err := json.NewDecoder(r.Body).Decode(&updatePost)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "You dun fucked",
			})
			return
		}

		updates := map[string]any{}

		if updatePost.Title != nil {
			updates["title"] = *updatePost.Title
		}
		if updatePost.Author != nil {
			updates["author"] = *updatePost.Author
		}

		result := db.Model(models.Post{}).Where("id = ?", id).Updates(updates)

		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "You dun fucked",
			})
			return
		}

		json.NewEncoder(w).Encode(updatePost)
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "You dun fucked",
			})
			return
		}

		result := db.Delete(&models.Post{}, "id = ?", id)

		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "You dun fucked",
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"success": true,
		})
	})

	return r
}
