package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/ferrazdourado/sar_api/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
            c.Abort()
            return
        }

        // Remove o prefixo "Bearer "
        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token mal formatado"})
            c.Abort()
            return
        }

        token := bearerToken[1]
        claims, err := utils.ValidateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            c.Abort()
            return
        }

        // Adiciona as claims ao contexto
        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}