package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/shynn12/biocad/internal/item"
	"github.com/shynn12/biocad/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

// Create implements item.Storage.
func (d *db) Create(ctx context.Context, item item.ItemDTO) (string, error) {
	d.logger.Debug("create item")
	result, err := d.collection.InsertOne(ctx, item)
	if err != nil {
		return "", fmt.Errorf("failed to create item due to error: %v", err)
	}

	d.logger.Debug("convert InsertedId to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(item)
	return "", fmt.Errorf("failed to convert objectid to hex")
}

func (d *db) FindAll(ctx context.Context, page, perPage int) (i []*item.Item, err error) {
	findOption := options.Find()

	findOption.SetSkip(int64((page - 1) * perPage))
	findOption.SetLimit(int64(perPage))

	cursor, err := d.collection.Find(ctx, bson.M{}, findOption)
	if cursor.Err() != nil {
		return i, fmt.Errorf("failed to find all items due to error: %v", err)
	}
	if err := cursor.All(ctx, &i); err != nil {
		return i, fmt.Errorf("failed to read all documents from cursor. error: %v", err)
	}

	return i, nil
}

// FindOne implements item.Storage.
func (d *db) FindOne(ctx context.Context, guid string) (i *item.Item, err error) {

	filter := bson.M{"unitguid": guid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return i, fmt.Errorf("not found")
		}
		return i, fmt.Errorf("failed to find one item by id: %s", guid)
	}
	if err = result.Decode(&i); err != nil {
		return i, fmt.Errorf("failde to decode item(id: %s) from DB due to error: %v", guid, err)
	}

	return i, nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) item.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
