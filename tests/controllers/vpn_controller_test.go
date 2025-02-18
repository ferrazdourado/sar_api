package controllers_test

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/ferrazdourado/sar_api/internal/controllers"
    "github.com/ferrazdourado/sar_api/internal/models"
    "github.com/ferrazdourado/sar_api/internal/services"
)

type MockVPNService struct {
    mock.Mock
}

func (m *MockVPNService) GetConfig(ctx context.Context, id string) (*models.VPNConfig, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*models.VPNConfig), args.Error(1)
}

func (m *MockVPNService) CreateConfig(ctx context.Context, config *models.VPNConfig) error {
    args := m.Called(ctx, config)
    return args.Error(0)
}

func (m *MockVPNService) ListConfigs(ctx context.Context, page, limit int) ([]*models.VPNConfig, int64, error) {
    args := m.Called(ctx, page, limit)
    return args.Get(0).([]*models.VPNConfig), args.Get(1).(int64), args.Error(2)
}

func (m *MockVPNService) UpdateConfig(ctx context.Context, config *models.VPNConfig) error {
    args := m.Called(ctx, config)
    return args.Error(0)
}

func (m *MockVPNService) DeleteConfig(ctx context.Context, id string) error {
    args := m.Called(ctx, id)
    return args.Error(0)
}

func (m *MockVPNService) GetStatus(ctx context.Context) (*models.VPNStatus, error) {
    args := m.Called(ctx)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*models.VPNStatus), args.Error(1)
}

func TestGetVPNConfig(t *testing.T) {
    gin.SetMode(gin.TestMode)
    
    tests := []struct {
        name       string
        id         string
        setupMock  func(*MockVPNService)
        wantStatus int
        wantConfig *models.VPNConfig
        wantErr    string
    }{
        {
            name: "Configuração encontrada",
            id:   "123",
            setupMock: func(m *MockVPNService) {
                m.On("GetConfig", mock.Anything, "123").Return(&models.VPNConfig{
                    ServerAddress: "vpn.example.com",
                    Port:         1194,
                    Protocol:     "udp",
                }, nil)
            },
            wantStatus: http.StatusOK,
            wantConfig: &models.VPNConfig{
                ServerAddress: "vpn.example.com",
                Port:         1194,
                Protocol:     "udp",
            },
        },
        {
            name: "Configuração não encontrada",
            id:   "456",
            setupMock: func(m *MockVPNService) {
                m.On("GetConfig", mock.Anything, "456").Return(nil, services.ErrConfigNotFound)
            },
            wantStatus: http.StatusNotFound,
            wantErr:    "configuração VPN não encontrada",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockService := new(MockVPNService)
            tt.setupMock(mockService)

            controller := controllers.NewVPNController(mockService)
            r := gin.New()
            r.GET("/vpn/config/:id", controller.GetVPNConfig)

            w := httptest.NewRecorder()
            req := httptest.NewRequest("GET", "/vpn/config/"+tt.id, nil)
            r.ServeHTTP(w, req)

            assert.Equal(t, tt.wantStatus, w.Code)

            if tt.wantConfig != nil {
                var response models.VPNConfig
                err := json.Unmarshal(w.Body.Bytes(), &response)
                assert.NoError(t, err)
                assert.Equal(t, tt.wantConfig.ServerAddress, response.ServerAddress)
                assert.Equal(t, tt.wantConfig.Port, response.Port)
                assert.Equal(t, tt.wantConfig.Protocol, response.Protocol)
            }

            if tt.wantErr != "" {
                var response map[string]string
                err := json.Unmarshal(w.Body.Bytes(), &response)
                assert.NoError(t, err)
                assert.Equal(t, tt.wantErr, response["error"])
            }

            mockService.AssertExpectations(t)
        })
    }
}