package mongodb

import (
    "context"
    "time"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/ferrazdourado/sar_api/internal/models"
)

var (
    ErrNotFound = errors.New("registro não encontrado")
)

type UserRepository struct {
    db         *MongoDB
    collection *mongo.Collection
}

func NewUserRepository(db *MongoDB) *UserRepository {
    return &UserRepository{
        db:         db,
        collection: db.client.Database(db.database).Collection("users"),
    }
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    
    result, err := r.collection.InsertOne(ctx, user)
    if err != nil {
        return err
    }
    
    user.ID = result.InsertedID.(primitive.ObjectID)
    return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var user models.User
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }

    return &user, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
    var user models.User
    err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) List(ctx context.Context, page, limit int) ([]*models.User, int64, error) {
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

    var users []*models.User
    if err = cursor.All(ctx, &users); err != nil {
        return nil, 0, err
    }

    return users, total, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
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

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
    user.UpdatedAt = time.Now()

    filter := bson.M{"_id": user.ID}
    update := bson.M{"$set": user}

    result, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return ErrNotFound
    }

    return nil
}