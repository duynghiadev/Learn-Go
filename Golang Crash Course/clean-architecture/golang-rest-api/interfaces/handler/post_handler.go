package handler

import (
	"encoding/json"
	"golang-clean-architecture/domain"
	"golang-clean-architecture/usecase"
	"log"
	"net/http"
)

type PostHandler struct {
	postUseCase usecase.PostUseCase
}

func NewPostHandler(uc usecase.PostUseCase) *PostHandler {
	return &PostHandler{postUseCase: uc}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Handler: Error marshalling JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal server error preparing response"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
	log.Printf("Handler: Responding with error - Code: %d, Message: %s", code, message)
	respondJSON(w, code, map[string]string{"error": message})
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := h.postUseCase.GetAllPosts(ctx)
	if err != nil {
		log.Printf("Handler: Error getting posts from use case: %v", err)
		respondError(w, http.StatusInternalServerError, "Error retrieving posts")
		return
	}

	log.Printf("Handler: Successfully retrieved %d posts", len(posts))
	respondJSON(w, http.StatusOK, posts)
}

func (h *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var post domain.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		log.Printf("Handler: Error decoding request body: %v", err)
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if post.Title == "" || post.Text == "" {
		respondError(w, http.StatusBadRequest, "Title and Text fields are required")
		return
	}

	createdPost, err := h.postUseCase.CreatePost(ctx, &post)
	if err != nil {
		log.Printf("Handler: Error creating post via use case: %v", err)
		respondError(w, http.StatusInternalServerError, "Error saving the post")
		return
	}

	log.Printf("Handler: Successfully created post with ID %s", createdPost.ID)
	respondJSON(w, http.StatusCreated, createdPost)
}
