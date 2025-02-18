package controllers_test

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "bytes"

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

func (m *MockVPNService) GetConfig() (*models.VPNConfig, error) {
    args := m.Called()
    return args.Get(0).(*models.VPNConfig), args.Error(1)
}

func TestGetVPNConfig(t *testing.T) {
    // Configurar o mock
    mockService := new(MockVPNService)
    mockConfig := &models.VPNConfig{
        ServerAddress: "vpn.example.com",
        Port:         1194,
        Protocol:     "udp",
    }
    mockService.On("GetConfig").Return(mockConfig, nil)

    // Configurar o controller com o mock
    controller := controllers.NewVPNController(mockService)

    // Configurar o router Gin
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/vpn/config", controller.GetVPNConfig)

    // Criar request de teste
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/vpn/config", nil)
    r.ServeHTTP(w, req)

    // Verificar resultado
    assert.Equal(t, http.StatusOK, w.Code)
    
    var response models.VPNConfig
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Equal(t, mockConfig.ServerAddress, response.ServerAddress)
    assert.Equal(t, mockConfig.Port, response.Port)
    assert.Equal(t, mockConfig.Protocol, response.Protocol)
}