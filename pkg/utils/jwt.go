package utils

import (
    "time"
    "github.com/golang-jwt/jwt"
    "github.com/ferrazdourado/sar_api/pkg/config"
)

type Claims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

func GenerateToken(claims Claims, config *config.JWTConfig) (string, error) {
    expirationTime := time.Now().Add(time.Hour * time.Duration(config.ExpiresIn))
    claims.StandardClaims = jwt.StandardClaims{
        ExpiresAt: expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.SigningKey))
}

func ValidateToken(tokenStr string, config *config.JWTConfig) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.SigningKey), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }

    return claims, nil
}