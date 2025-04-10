package usecase

import (
	"context"
	"golang-clean-architecture/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPostRepository is a mock type for repository.PostRepository
type MockPostRepository struct {
	mock.Mock
}

// Save is a mocked method
func (m *MockPostRepository) Save(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	args := m.Called(ctx, post)
	return args.Get(0).(*domain.Post), args.Error(1)
}

// FindAll is a mocked method
func (m *MockPostRepository) FindAll(ctx context.Context) ([]domain.Post, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Post), args.Error(1)
}

func TestCreatePost(t *testing.T) {
	mockRepo := new(MockPostRepository)
	postUseCase := NewPostInteractor(mockRepo)

	ctx := context.Background()
	post := &domain.Post{
		Title: "Test Post",
		Text:  "Test Content",
	}

	// Setup expectations
	mockRepo.On("Save", ctx, post).Return(post, nil)

	// Test the CreatePost method
	result, err := postUseCase.CreatePost(ctx, post)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, post.Title, result.Title)
	assert.Equal(t, post.Text, result.Text)

	// Verify that expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetAllPosts(t *testing.T) {
	mockRepo := new(MockPostRepository)
	postUseCase := NewPostInteractor(mockRepo)

	ctx := context.Background()
	expectedPosts := []domain.Post{
		{
			Title: "Test Post 1",
			Text:  "Test Text 1",
		},
		{
			Title: "Test Post 2",
			Text:  "Test Text 2",
		},
	}

	// Setup expectations
	mockRepo.On("FindAll", ctx).Return(expectedPosts, nil)

	// Test the GetAllPosts method
	results, err := postUseCase.GetAllPosts(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Len(t, results, 2)
	assert.Equal(t, expectedPosts[0].Title, results[0].Title)
	assert.Equal(t, expectedPosts[1].Title, results[1].Title)

	// Verify that expectations were met
	mockRepo.AssertExpectations(t)
}
