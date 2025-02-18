package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type VPNConfig struct {
    ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name          string            `bson:"name" json:"name" binding:"required"`
    ServerAddress string            `bson:"server_address" json:"server_address" binding:"required"`
    Port          int               `bson:"port" json:"port" binding:"required"`
    Protocol      string            `bson:"protocol" json:"protocol" binding:"required"`
    Config        string            `bson:"config" json:"config" binding:"required"`
    CreatedAt     time.Time         `bson:"created_at" json:"created_at"`
    UpdatedAt     time.Time         `bson:"updated_at" json:"updated_at"`
    CreatedBy     primitive.ObjectID `bson:"created_by" json:"created_by"`
}

type VPNStatus struct {
    Status         string    `json:"status"`
    ConnectedSince time.Time `json:"connected_since,omitempty"`
    BytesSent      int64     `json:"bytes_sent"`
    BytesReceived  int64     `json:"bytes_received"`
    Error          string    `json:"error,omitempty"`
}