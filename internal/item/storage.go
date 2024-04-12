package item

import "context"

type Storage interface {
	Create(ctx context.Context, item ItemDTO) (string, error)
	FindOne(ctx context.Context, id string) (*Item, error)
	FindAll(ctx context.Context, page, perPage int) ([]*Item, error)
}
