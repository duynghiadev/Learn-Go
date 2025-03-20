package handler

import (
	"net/http"

	"github.com/duynghiadev/backend-github-trending/model"
	"github.com/duynghiadev/backend-github-trending/repository"
	"github.com/labstack/echo"
)

type RepoHandler struct {
	GithubRepo repository.GithubRepo
}

func (r RepoHandler) RepoTrending(c echo.Context) error {
	repos, _ := r.GithubRepo.SelectRepos(c.Request().Context(), 25)
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       repos,
	})
}
