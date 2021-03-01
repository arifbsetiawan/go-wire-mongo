package author

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorRepository struct {
	MongoDB *mongo.Database
}

func ProvideAuthorRepository(db *mongo.Database) AuthorRepository {
	return AuthorRepository{
		MongoDB: db,
	}
}

func (r *AuthorRepository) collection() *mongo.Collection {
	return r.MongoDB.Collection("authors")
}

func (r *AuthorRepository) Find(ctx context.Context) (*[]Author, error) {
	var authors []Author
	data, err := r.collection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for data.Next(ctx) {
		var author Author
		data.Decode(&author)
		authors = append(authors, author)
	}

	return &authors, nil
}

func (r *AuthorRepository) FindOne(ctx context.Context, filter interface{}) (*Author, error) {
	var author Author
	if err := r.collection().FindOne(ctx, filter).Decode(&author); err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *AuthorRepository) Create(ctx context.Context, data Author) error {
	if _, err := r.collection().InsertOne(ctx, data); err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepository) Update(ctx context.Context, data interface{}, filter interface{}) error {
	if _, err := r.collection().UpdateOne(ctx, filter, data); err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepository) Delete(ctx context.Context, filter interface{}) error {
	if _, err := r.collection().DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
