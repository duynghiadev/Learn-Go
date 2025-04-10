package repository

import (
	"context"
	"golang-clean-architecture/domain"
)

type PostRepository interface {
	Save(ctx context.Context, post *domain.Post) (*domain.Post, error)
	FindAll(ctx context.Context) ([]domain.Post, error)
}
