package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ferrazdourado/sar_api/internal/models"
    "github.com/ferrazdourado/sar_api/internal/services"
)

type VPNController struct {
    vpnService *services.VPNService
}

func NewVPNController(service *services.VPNService) *VPNController {
    return &VPNController{
        vpnService: service,
    }
}

func (c *VPNController) GetVPNConfig(ctx *gin.Context) {
    config, err := c.vpnService.GetConfig()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, config)
}

func (c *VPNController) CreateVPNConfig(ctx *gin.Context) {
    var config models.VPNConfig
    if err := ctx.ShouldBindJSON(&config); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.vpnService.CreateConfig(&config); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, config)
}

func (c *VPNController) GetVPNStatus(ctx *gin.Context) {
    status, err := c.vpnService.GetStatus()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, status)
}