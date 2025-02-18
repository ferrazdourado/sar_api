package middleware_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/ferrazdourado/sar_api/internal/middleware"
    "github.com/ferrazdourado/sar_api/pkg/config"
    "github.com/ferrazdourado/sar_api/pkg/utils"
)

func TestAuthMiddleware(t *testing.T) {
    // Configurar o modo de teste do Gin
    gin.SetMode(gin.TestMode)

    // Criar configuração de teste
    cfg := &config.Config{
        JWT: config.JWTConfig{
            Secret:     "test_secret",
            ExpiresIn:  24,
            SigningKey: "test_signing_key",
        },
    }

    // Criar middleware
    authMiddleware := middleware.NewAuthMiddleware(cfg)

    tests := []struct {
        name       string
        token      string
        wantStatus int
        wantError  string
    }{
        {
            name:       "Token válido",
            token:      createValidToken(t, cfg),
            wantStatus: http.StatusOK,
            wantError:  "",
        },
        {
            name:       "Sem token",
            token:      "",
            wantStatus: http.StatusUnauthorized,
            wantError:  "Token não fornecido",
        },
        {
            name:       "Token inválido",
            token:      "invalid.token.here",
            wantStatus: http.StatusUnauthorized,
            wantError:  "Token inválido",
        },
        {
            name:       "Token mal formatado",
            token:      "malformed",
            wantStatus: http.StatusUnauthorized,
            wantError:  "Token mal formatado",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Configurar router de teste
            r := gin.New()
            r.Use(authMiddleware.Handler())
            
            // Adicionar rota protegida para teste
            r.GET("/protected", func(c *gin.Context) {
                userID, exists := c.Get("userID")
                assert.True(t, exists)
                assert.NotEmpty(t, userID)
                c.Status(http.StatusOK)
            })

            // Criar request de teste
            w := httptest.NewRecorder()
            req := httptest.NewRequest("GET", "/protected", nil)
            if tt.token != "" {
                req.Header.Set("Authorization", "Bearer "+tt.token)
            }

            // Executar request
            r.ServeHTTP(w, req)

            // Verificar resultados
            assert.Equal(t, tt.wantStatus, w.Code)
            
            if tt.wantError != "" {
                var response map[string]string
                err := json.NewDecoder(w.Body).Decode(&response)
                assert.NoError(t, err)
                assert.Equal(t, tt.wantError, response["error"])
            }
        })
    }
}

// Função auxiliar para criar um token válido para testes
func createValidToken(t *testing.T, cfg *config.Config) string {
    claims := utils.Claims{
        UserID: "test-user",
        Role:   "admin",
    }
    
    token, err := utils.GenerateToken(claims, &cfg.JWT)
    assert.NoError(t, err)
    return token
}

// Teste específico para verificação de roles
func TestAuthMiddlewareRoles(t *testing.T) {
    gin.SetMode(gin.TestMode)
    
    cfg := &config.Config{
        JWT: config.JWTConfig{
            Secret:     "test_secret",
            SigningKey: "test_signing_key",
        },
    }

    authMiddleware := middleware.NewAuthMiddleware(cfg)

    r := gin.New()
    r.Use(authMiddleware.Handler())

    // Rota que requer role específica
    r.GET("/admin", func(c *gin.Context) {
        role, exists := c.Get("role")
        assert.True(t, exists)
        assert.Equal(t, "admin", role)
        c.Status(http.StatusOK)
    })

    t.Run("Acesso com role correto", func(t *testing.T) {
        token := createTokenWithRole(t, cfg, "admin")
        w := httptest.NewRecorder()
        req := httptest.NewRequest("GET", "/admin", nil)
        req.Header.Set("Authorization", "Bearer "+token)

        r.ServeHTTP(w, req)
        assert.Equal(t, http.StatusOK, w.Code)
    })

    t.Run("Acesso com role incorreto", func(t *testing.T) {
        token := createTokenWithRole(t, cfg, "user")
        w := httptest.NewRecorder()
        req := httptest.NewRequest("GET", "/admin", nil)
        req.Header.Set("Authorization", "Bearer "+token)

        r.ServeHTTP(w, req)
        assert.Equal(t, http.StatusForbidden, w.Code)
    })
}

func createTokenWithRole(t *testing.T, cfg *config.Config, role string) string {
    claims := utils.Claims{
        UserID: "test-user",
        Role:   role,
    }
    
    token, err := utils.GenerateToken(claims, &cfg.JWT)
    assert.NoError(t, err)
    return token
}