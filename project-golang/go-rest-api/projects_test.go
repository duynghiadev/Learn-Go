package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateProject(t *testing.T) {
	// Create a new project
	ms := &MockStore{}
	service := NewProjectService(ms)

	t.Run("should validate if the name is not empty", func(t *testing.T) {
		payload := &CreateProjectPayload{
			Name: "",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/projects", service.handleCreateProject)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

		var response ErrorResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		if response.Error != "Name is required" {
			t.Errorf("expected error message %s, got %s", "Name is required", response.Error)
		}
	})

	t.Run("should create a project", func(t *testing.T) {
		payload := &CreateProjectPayload{
			Name: "Super cool project",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/projects", service.handleCreateProject)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

func TestGetProject(t *testing.T) {
	// Create a new project
	ms := &MockStore{}
	service := NewProjectService(ms)

	t.Run("should return a project", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/projects", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/projects", service.handleGetProject)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}

		var response Project
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		if response.Name != "Super cool project" {
			t.Errorf("expected project name %s, got %s", "Super cool project", response.Name)
		}
	})
}

func TestDeleteProject(t *testing.T) {
	// Create a new project
	ms := &MockStore{}
	service := NewProjectService(ms)

	t.Run("should delete the project", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/projects", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/projects", service.handleDeleteProject)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("expected status code %d, got %d", http.StatusNoContent, rr.Code)
		}
	})
}

func TestGetAllProjects(t *testing.T) {
	// Setup
	ms := &MockStore{}
	service := NewProjectService(ms)

	t.Run("should return all projects", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/projects", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/projects", service.handleGetAllProjects)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}

		var projects []*Project
		err = json.NewDecoder(rr.Body).Decode(&projects)
		if err != nil {
			t.Fatal(err)
		}

		if len(projects) != 2 {
			t.Errorf("expected 2 projects, got %d", len(projects))
		}

		if projects[0].Name != "Test Project 1" || projects[1].Name != "Test Project 2" {
			t.Errorf("unexpected project names: %+v", projects)
		}
	})
}
