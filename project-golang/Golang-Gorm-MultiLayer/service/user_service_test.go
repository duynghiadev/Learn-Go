package service_test

import (
	"golang-transaction/model"
	"golang-transaction/repository"
	"golang-transaction/service"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_SaveAndGetAll(t *testing.T) {
	db, err := model.DBConnection()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	repo := repository.NewUserRepository(db)
	repo.Migrate()
	svc := service.NewUserService(repo)

	user := model.User{
		Email:  "service_test@example.com",
		Wallet: 200,
	}

	savedUser, err := svc.Save(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, savedUser.Email)

	users, err := svc.GetAll()
	assert.NoError(t, err)
	assert.True(t, len(users) > 0)
}
