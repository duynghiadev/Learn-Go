package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errEmailRequired = errors.New("email is required")
var errFirstNameRequired = errors.New("first name is required")
var errLastNameRequired = errors.New("last name is required")
var errPasswordRequired = errors.New("password is required")

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/users/login", s.handleUserLogin).Methods("POST")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var payload *User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if validationErr := validateUserPayload(payload); validationErr != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: validationErr.Error()})
		return
	}

	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}
	payload.Password = hashedPassword

	u, err := s.store.CreateUser(payload)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}

	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}

	WriteJSON(w, http.StatusCreated, token)
}

func (s *UserService) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if unmarshalErr := json.Unmarshal(body, &loginData); unmarshalErr != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	// Step 1: Get user from DB
	user, err := s.store.GetUserByEmail(loginData.Email)
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "Invalid email or password"})
		return
	}

	// Step 2: Check password
	if checkErr := CheckPasswordHash(loginData.Password, user.Password); checkErr != nil {
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "Invalid email or password"})
		return
	}

	// Step 3: Create JWT + Set cookie
	token, err := createAndSetAuthCookie(user.ID, w)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to create auth token"})
		return
	}

	// Step 4: Return token in response
	WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func validateUserPayload(user *User) error {
	if user.Email == "" {
		return errEmailRequired
	}

	if user.FirstName == "" {
		return errFirstNameRequired
	}

	if user.LastName == "" {
		return errLastNameRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func createAndSetAuthCookie(userID int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, userID)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
