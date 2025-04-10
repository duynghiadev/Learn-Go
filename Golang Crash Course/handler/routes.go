package handler

import (
	"encoding/json"
	entity "golang-rest-api/enitty"
	"golang-rest-api/repository"
	"net/http"
)

// PostHandler struct to hold the repository
type PostHandler struct {
	repo repository.PostRepository
}

// NewPostHandler creates a new handler with the injected repository
func NewPostHandler(repo repository.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

// GetPosts retrieves all posts
func (h *PostHandler) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := h.repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

// AddPost adds a new post
func (h *PostHandler) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshaling the request"}`))
		return
	}

	createdPost, err := h.repo.Save(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error saving the post"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(createdPost)
}
