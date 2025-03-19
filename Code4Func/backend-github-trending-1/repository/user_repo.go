package repository

import (
	"context"

	"github.com/duynghiadev/backend-github-trending/model"
	"github.com/duynghiadev/backend-github-trending/model/req"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)
	SelectUserById(context context.Context, userId string) (model.User, error)
	UpdateUser(context context.Context, user model.User) (model.User, error)
}
