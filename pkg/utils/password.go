package utils

import (
    "crypto/rand"
    "encoding/base64"
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedBytes), nil
}

func ComparePasswords(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateRandomKey(length int) (string, error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(bytes), nil
}