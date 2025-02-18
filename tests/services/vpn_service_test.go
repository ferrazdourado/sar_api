package services_test

import (
    "context"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/ferrazdourado/sar_api/internal/models"
    "github.com/ferrazdourado/sar_api/internal/services"
    "github.com/ferrazdourado/sar_api/pkg/config"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
    args := m.Called(ctx, username)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*models.User), args.Error(1)
}

func TestLogin(t *testing.T) {
    mockRepo := new(MockUserRepository)
    authService := services.NewAuthService(mockRepo)
    ctx := context.Background()

    t.Run("Login bem sucedido", func(t *testing.T) {
        mockUser := &models.User{
            Username: "testuser",
            Password: "$2a$10$...", // senha hasheada
        }
        
        mockRepo.On("FindByUsername", ctx, "testuser").Return(mockUser, nil)

        token, err := authService.Login(ctx, models.LoginCredentials{
            Username: "testuser",
            Password: "password123",
        })

        assert.NoError(t, err)
        assert.NotEmpty(t, token)
    })
}