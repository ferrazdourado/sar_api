package routes_test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "bytes"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/ferrazdourado/sar_api/internal/routes"
    "github.com/ferrazdourado/sar_api/internal/controllers"
)

type MockAuthController struct {
    mock.Mock
}

type MockVPNController struct {
    mock.Mock
}

func TestPublicRoutes(t *testing.T) {
    mockAuth := new(MockAuthController)
    mockVPN := new(MockVPNController)

    router := routes.NewRouter(mockVPN, mockAuth)
    r := router.SetupRoutes()

    tests := []struct {
        name       string
        method     string
        path       string
        wantStatus int
    }{
        {
            name:       "Login route",
            method:     "POST",
            path:       "/api/v1/auth/login",
            wantStatus: http.StatusOK,
        },
        {
            name:       "Register route",
            method:     "POST",
            path:       "/api/v1/auth/register",
            wantStatus: http.StatusOK,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            w := httptest.NewRecorder()
            req := httptest.NewRequest(tt.method, tt.path, nil)
            r.ServeHTTP(w, req)

            assert.Equal(t, tt.wantStatus, w.Code)
        })
    }
}