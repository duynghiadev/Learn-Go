package repository

import (
	models "github.com/duynghiadev/go-jwt/model"
)

type UserRepo interface {
	FindUserByEmail(email string) (models.User, error)
	CheckLoginInfo(email string, password string) (models.User, error)
	Insert(u models.User) error
}
