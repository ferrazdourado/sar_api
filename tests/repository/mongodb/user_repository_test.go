package mongodb_test

import (
    "context"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/ferrazdourado/sar_api/internal/repository/mongodb"
    "github.com/ferrazdourado/sar_api/internal/models"
)

func TestUserRepository(t *testing.T) {
    ctx := context.Background()
    db := mongodb.NewMongoDB("mongodb://localhost:27017", "sar_test")
    err := db.Connect(ctx)
    assert.NoError(t, err)
    defer db.Disconnect(ctx)

    repo := mongodb.NewUserRepository(db)

    t.Run("Create and Find User", func(t *testing.T) {
        user := &models.User{
            Username: "testuser",
            Email:    "test@example.com",
            Password: "password123",
        }

        err := repo.Create(ctx, user)
        assert.NoError(t, err)
        assert.NotEmpty(t, user.ID)

        found, err := repo.FindByID(ctx, user.ID.Hex())
        assert.NoError(t, err)
        assert.Equal(t, user.Username, found.Username)
    })
}