package mongodb

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
    client   *mongo.Client
    database string
    uri      string
}

func NewMongoDB(uri, database string) *MongoDB {
    return &MongoDB{
        database: database,
        uri:      uri,
    }
}

func (m *MongoDB) Connect(ctx context.Context) error {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
    if err != nil {
        return err
    }
    m.client = client
    return nil
}

func (m *MongoDB) Disconnect(ctx context.Context) error {
    return m.client.Disconnect(ctx)
}

func (m *MongoDB) Ping(ctx context.Context) error {
    return m.client.Ping(ctx, nil)
}

func (m *MongoDB) Transaction(ctx context.Context, fn func(context.Context) error) error {
    session, err := m.client.StartSession()
    if err != nil {
        return err
    }
    defer session.EndSession(ctx)

    return mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
        if err := session.StartTransaction(); err != nil {
            return err
        }
        if err := fn(sc); err != nil {
            return session.AbortTransaction(sc)
        }
        return session.CommitTransaction(sc)
    })
}