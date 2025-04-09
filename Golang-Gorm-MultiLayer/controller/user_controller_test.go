package controller_test

import (
	"bytes"
	"encoding/json"
	"golang-transaction/controller"
	"golang-transaction/model"
	"golang-transaction/repository"
	"golang-transaction/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	db, _ := model.DBConnection()
	repo := repository.NewUserRepository(db)
	repo.Migrate()
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)

	r := gin.Default()
	r.POST("/users", ctrl.AddUser)
	r.GET("/users", ctrl.GetAllUser)
	return r
}

func TestAddUserAPI(t *testing.T) {
	r := setupRouter()

	user := model.User{
		Email:  "test3@gmail.com",
		Wallet: 300,
	}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllUsersAPI(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
