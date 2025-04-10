package usecase

import (
	"context"
	"golang-clean-architecture/domain"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error)
	GetAllPosts(ctx context.Context) ([]domain.Post, error)
}
