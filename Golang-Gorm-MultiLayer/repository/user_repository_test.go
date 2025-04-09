package repository_test

import (
	"golang-transaction/model"
	"golang-transaction/repository"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_SaveAndGetAll(t *testing.T) {
	db, err := model.DBConnection()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	repo := repository.NewUserRepository(db)
	repo.Migrate()

	// Save user
	newUser := model.User{
		Email:  "test1@example.com",
		Wallet: 100.0,
	}
	savedUser, err := repo.Save(newUser)
	assert.NoError(t, err)
	assert.Equal(t, newUser.Email, savedUser.Email)

	// Get all users
	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.True(t, len(users) > 0)
}

func TestUserRepository_IncrementAndDecrement(t *testing.T) {
	db, _ := model.DBConnection()
	repo := repository.NewUserRepository(db)

	// Save user and capture the returned ID
	user := model.User{Email: "test11@gmail.com", Wallet: 2302.4}
	savedUser, err := repo.Save(user)
	assert.NoError(t, err)

	// Increment
	err = repo.IncrementMoney(savedUser.ID, 21)
	assert.NoError(t, err)

	// Decrement with enough funds
	err = repo.DecrementMoney(savedUser.ID, 5)
	assert.NoError(t, err)

	// Decrement with not enough funds
	err = repo.DecrementMoney(savedUser.ID, 999999) // Ensure this is more than wallet
	assert.Error(t, err)
}
