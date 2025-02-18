package mongodb

import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/ferrazdourado/sar_api/internal/models"
)

type VPNRepository struct {
    db         *MongoDB
    collection *mongo.Collection
}

func NewVPNRepository(db *MongoDB) *VPNRepository {
    return &VPNRepository{
        db:         db,
        collection: db.client.Database(db.database).Collection("vpn_configs"),
    }
}

func (r *VPNRepository) CreateConfig(ctx context.Context, config *models.VPNConfig) error {
    config.CreatedAt = time.Now()
    config.UpdatedAt = time.Now()
    
    result, err := r.collection.InsertOne(ctx, config)
    if err != nil {
        return err
    }
    
    config.ID = result.InsertedID.(primitive.ObjectID)
    return nil
}

func (r *VPNRepository) GetConfig(ctx context.Context, id string) (*models.VPNConfig, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var config models.VPNConfig
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&config)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, ErrNotFound
        }
        return nil, err
    }

    return &config, nil
}

func (r *VPNRepository) ListConfigs(ctx context.Context, page, limit int) ([]*models.VPNConfig, int64, error) {
    skip := int64((page - 1) * limit)
    
    total, err := r.collection.CountDocuments(ctx, bson.M{})
    if err != nil {
        return nil, 0, err
    }

    opts := options.Find().
        SetSkip(skip).
        SetLimit(int64(limit)).
        SetSort(bson.D{{Key: "created_at", Value: -1}})

    cursor, err := r.collection.Find(ctx, bson.M{}, opts)
    if err != nil {
        return nil, 0, err
    }
    defer cursor.Close(ctx)

    var configs []*models.VPNConfig
    if err = cursor.All(ctx, &configs); err != nil {
        return nil, 0, err
    }

    return configs, total, nil
}

func (r *VPNRepository) UpdateConfig(ctx context.Context, config *models.VPNConfig) error {
    config.UpdatedAt = time.Now()

    filter := bson.M{"_id": config.ID}
    update := bson.M{"$set": config}

    result, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return ErrNotFound
    }

    return nil
}

func (r *VPNRepository) DeleteConfig(ctx context.Context, id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
    if err != nil {
        return err
    }

    if result.DeletedCount == 0 {
        return ErrNotFound
    }

    return nil
}

func (r *VPNRepository) GetStatus(ctx context.Context) (*models.VPNStatus, error) {
    // Implementação básica - você pode expandir conforme necessário
    return &models.VPNStatus{
        Status:         "connected",
        ConnectedSince: time.Now(),
        BytesSent:     0,
        BytesReceived: 0,
    }, nil
}