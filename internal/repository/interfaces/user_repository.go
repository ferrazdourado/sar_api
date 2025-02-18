package interfaces

import (
    "context"
    "github.com/ferrazdourado/sar_api/internal/models"
)

type UserRepository interface {
    Create(ctx context.Context, user *models.User) error
    FindByID(ctx context.Context, id string) (*models.User, error)
    FindByUsername(ctx context.Context, username string) (*models.User, error)
    Update(ctx context.Context, user *models.User) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context, page, limit int) ([]*models.User, int64, error)
}