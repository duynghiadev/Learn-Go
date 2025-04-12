package firestore

import (
	"context"
	"encoding/json"
	"golang-clean-architecture/domain"
	repo "golang-clean-architecture/usecase/repository"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/iterator"
)

const (
	collectionName string = "posts"
)

type firestorePostRepository struct {
	client *firestore.Client
	redis  *redis.Client
}

func NewFirestorePostRepository(client *firestore.Client, redis *redis.Client) repo.PostRepository {
	return &firestorePostRepository{
		client: client,
		redis:  redis,
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
	// Try to get from Redis first
	cacheKey := "all_posts"
	cachedPosts, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var posts []domain.Post
		if err = json.Unmarshal([]byte(cachedPosts), &posts); err != nil {
			log.Printf("Firestore Repo: Failed to unmarshal cached posts: %v", err)
		} else {
			log.Printf("Firestore Repo: Retrieved %d posts from cache", len(posts))
			return posts, nil
		}
	} else if err != redis.Nil {
		log.Printf("Firestore Repo: Failed to get posts from Redis: %v", err)
	}

	// If not in cache, get from Firestore
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

	// Cache the results
	if len(posts) > 0 {
		postsJSON, err := json.Marshal(posts)
		if err == nil {
			err = r.redis.Set(ctx, cacheKey, postsJSON, 5000*time.Minute).Err()
			if err != nil {
				log.Printf("Firestore Repo: Failed to cache posts in Redis: %v", err)
			}
		} else {
			log.Printf("Firestore Repo: Failed to marshal posts for caching: %v", err)
		}
	}

	return posts, nil
}
