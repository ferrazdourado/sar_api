package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Username  string            `bson:"username" json:"username" binding:"required"`
    Password  string            `bson:"password" json:"password,omitempty" binding:"required"`
    Email     string            `bson:"email" json:"email" binding:"required,email"`
    Role      string            `bson:"role" json:"role"`
    CreatedAt time.Time         `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time         `bson:"updated_at" json:"updated_at"`
}

type LoginCredentials struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}