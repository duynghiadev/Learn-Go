package firestore

import (
	"context"
	"golang-clean-architecture/domain"
	repo "golang-clean-architecture/usecase/repository"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	collectionName string = "posts"
)

type firestorePostRepository struct {
	client *firestore.Client
}

func NewFirestorePostRepository(client *firestore.Client) repo.PostRepository {
	return &firestorePostRepository{
		client: client,
	}
}

func (r *firestorePostRepository) Save(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	ref := r.client.Collection(collectionName).NewDoc()
	post.ID = ref.ID

	_, err := ref.Set(ctx, post)
	if err != nil {
		log.Printf("Firestore Repo: Failed adding post with ID %s: %v", post.ID, err)
		return nil, err
	}
	log.Printf("Firestore Repo: Successfully saved post with ID %s", post.ID)
	return post, nil
}

func (r *firestorePostRepository) FindAll(ctx context.Context) ([]domain.Post, error) {
	var posts []domain.Post
	iter := r.client.Collection(collectionName).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Firestore Repo: Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		var post domain.Post
		if err := doc.DataTo(&post); err != nil {
			log.Printf("Firestore Repo: Failed to convert document %s to Post: %v", doc.Ref.ID, err)
			continue
		}

		if post.ID == "" {
			post.ID = doc.Ref.ID
		}
		posts = append(posts, post)
	}
	log.Printf("Firestore Repo: Found %d posts", len(posts))
	return posts, nil
}
