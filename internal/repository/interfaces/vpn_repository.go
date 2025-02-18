package interfaces

import (
    "context"
    "github.com/ferrazdourado/sar_api/internal/models"
)

type VPNRepository interface {
    CreateConfig(ctx context.Context, config *models.VPNConfig) error
    GetConfig(ctx context.Context, id string) (*models.VPNConfig, error)
    ListConfigs(ctx context.Context, page, limit int) ([]*models.VPNConfig, int64, error)
    UpdateConfig(ctx context.Context, config *models.VPNConfig) error
    DeleteConfig(ctx context.Context, id string) error
    GetStatus(ctx context.Context) (*models.VPNStatus, error)
}