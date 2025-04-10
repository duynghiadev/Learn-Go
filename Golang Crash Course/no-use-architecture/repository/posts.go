package repository

import (
	"context"
	entity "golang-rest-api/enitty"
	"log"

	"cloud.google.com/go/firestore"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct {
	client *firestore.Client
}

func NewPostRepository(client *firestore.Client) PostRepository {
	return &repo{client: client}
}

const (
	collectionName string = "posts"
)

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	ref := r.client.Collection(collectionName).NewDoc()
	post.ID = ref.ID

	_, err := ref.Set(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Printf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	var posts []entity.Post
	iterator := r.client.Collection(collectionName).Documents(ctx)
	defer iterator.Stop()

	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Printf("Failed to iterate the list of posts: %v", err)
			break
		}

		if doc == nil {
			break
		}

		var post entity.Post
		if err := doc.DataTo(&post); err != nil {
			log.Printf("Failed to convert document to Post: %v", err)
			continue
		}
		posts = append(posts, post)
	}
	return posts, nil
}
