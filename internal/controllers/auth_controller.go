package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ferrazdourado/sar_api/internal/models"
    "github.com/ferrazdourado/sar_api/internal/services"
)

type AuthController struct {
    authService *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
    return &AuthController{
        authService: service,
    }
}

func (c *AuthController) Login(ctx *gin.Context) {
    var credentials models.LoginCredentials
    if err := ctx.ShouldBindJSON(&credentials); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := c.authService.Login(credentials)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *AuthController) Register(ctx *gin.Context) {
    var user models.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.authService.Register(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "Usuário registrado com sucesso"})
}