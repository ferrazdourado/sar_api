package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/ferrazdourad/sar_api/internal/controllers"
)

func TestGetVPNConfig(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/vpn/config", controllers.GetVPNConfig)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/vpn/config", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    var response map[string]interface{}
    err := json.Unmarshal(w.Body.package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/ferrazdourado/sar_api/internal/controllers"
)

func TestGetVPNConfig(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/vpn/config", controllers.GetVPNConfig)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/vpn/config", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    var response map[string]interface{}
    err := json.Unmarshal(w.Body.