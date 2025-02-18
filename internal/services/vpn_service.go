package services

import (
    "context"
    "errors"
    "github.com/ferrazdourado/sar_api/internal/models"
    "github.com/ferrazdourado/sar_api/internal/repository/interfaces"
)

type VPNService struct {
    vpnRepo interfaces.VPNRepository
}

func NewVPNService(vpnRepo interfaces.VPNRepository) *VPNService {
    return &VPNService{
        vpnRepo: vpnRepo,
    }
}

func (s *VPNService) CreateConfig(ctx context.Context, config *models.VPNConfig) error {
    // Validar configuração
    if err := s.validateConfig(config); err != nil {
        return err
    }

    return s.vpnRepo.CreateConfig(ctx, config)
}

func (s *VPNService) GetConfig(ctx context.Context, id string) (*models.VPNConfig, error) {
    config, err := s.vpnRepo.GetConfig(ctx, id)
    if err != nil {
        return nil, err
    }
    if config == nil {
        return nil, errors.New("configuração não encontrada")
    }
    return config, nil
}

func (s *VPNService) ListConfigs(ctx context.Context, page, limit int) ([]*models.VPNConfig, int64, error) {
    if page < 1 {
        page = 1
    }
    if limit < 1 || limit > 100 {
        limit = 10
    }
    return s.vpnRepo.ListConfigs(ctx, page, limit)
}

func (s *VPNService) GetStatus(ctx context.Context) (*models.VPNStatus, error) {
    return s.vpnRepo.GetStatus(ctx)
}

func (s *VPNService) validateConfig(config *models.VPNConfig) error {
    if config.Name == "" {
        return errors.New("nome é obrigatório")
    }
    if config.ServerAddress == "" {
        return errors.New("endereço do servidor é obrigatório")
    }
    if config.Port < 1 || config.Port > 65535 {
        return errors.New("porta inválida")
    }
    if config.Protocol != "udp" && config.Protocol != "tcp" {
        return errors.New("protocolo deve ser 'udp' ou 'tcp'")
    }
    return nil
}