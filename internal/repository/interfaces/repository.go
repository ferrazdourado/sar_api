package interfaces

import "context"

type Repository interface {
    Connect(ctx context.Context) error
    Disconnect(ctx context.Context) error
    Ping(ctx context.Context) error
    Transaction(ctx context.Context, fn func(context.Context) error) error
}