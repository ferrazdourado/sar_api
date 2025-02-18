package middleware_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/ferrazdourado/sar_api/internal/middleware"
    "github.com/ferrazdourado/sar_api/pkg/utils"
)

func TestAuthMiddleware(t *testing.T) {
    // Configurar Gin em modo teste
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name       string
        token      string
        wantStatus int
    }{
        {
            name:       "Token válido",
            token:      createValidToken(),
            wantStatus: http.StatusOK,
        },
        {
            name:       "Sem token",
            token:      "",
            wantStatus: http.StatusUnauthorized,
        },
        {
            name:       "Token inválido",
            token:      "invalid.token.here",
            wantStatus: http.StatusUnauthorized,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := gin.New()
            r.Use(middleware.AuthMiddleware())
            
            // Rota protegida para teste
            r.GET("/protected", func(c *gin.Context) {
                c.Status(http.StatusOK)
            })

            w := httptest.NewRecorder()
            req := httptest.NewRequest("GET", "/protected", nil)
            if tt.token != "" {
                req.Header.Set("Authorization", "Bearer "+tt.token)
            }

            r.ServeHTTP(w, req)
            assert.Equal(t, tt.wantStatus, w.Code)
        })
    }
}

func createValidToken() string {
    // Criar um token válido para testes
    claims := utils.Claims{
        UserID: "test-user",
        Role:   "admin",
    }
    token, _ := utils.GenerateToken(claims)
    return token
}