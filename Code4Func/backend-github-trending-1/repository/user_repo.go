package repository

import (
	"context"

	"github.com/duynghiadev/backend-github-trending/model"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}
