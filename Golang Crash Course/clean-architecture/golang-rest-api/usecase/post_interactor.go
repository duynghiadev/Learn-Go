package usecase

import (
	"context"
	"golang-clean-architecture/domain"
	"golang-clean-architecture/usecase/repository"
	"log"
)

type postInteractor struct {
	postRepo repository.PostRepository
}

func NewPostInteractor(repo repository.PostRepository) PostUseCase {
	return &postInteractor{
		postRepo: repo,
	}
}

func (pi *postInteractor) CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	log.Printf("Interactor: Creating post with title '%s'", post.Title)

	createdPost, err := pi.postRepo.Save(ctx, post)
	if err != nil {
		log.Printf("Interactor: Error saving post: %v", err)
		return nil, err
	}
	return createdPost, nil
}

func (pi *postInteractor) GetAllPosts(ctx context.Context) ([]domain.Post, error) {
	log.Println("Interactor: Getting all posts")
	posts, err := pi.postRepo.FindAll(ctx)
	if err != nil {
		log.Printf("Interactor: Error finding all posts: %v", err)
		return nil, err
	}
	return posts, nil
}
