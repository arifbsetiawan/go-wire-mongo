package book

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	MongoDB *mongo.Database
}

func ProvideBookRepository(db *mongo.Database) BookRepository {
	return BookRepository{
		MongoDB: db,
	}
}

func (r *BookRepository) collection() *mongo.Collection {
	return r.MongoDB.Collection("books")
}

func (r *BookRepository) Find(ctx context.Context, filter interface{}) (*[]Book, error) {
	var books []Book
	data, err := r.collection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for data.Next(ctx) {
		var book Book
		data.Decode(&book)
		books = append(books, book)
	}

	return &books, nil
}

func (r *BookRepository) FindOne(ctx context.Context, filter interface{}) (*Book, error) {
	var book Book
	if err := r.collection().FindOne(ctx, filter).Decode(&book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) Create(ctx context.Context, data Book) error {
	if _, err := r.collection().InsertOne(ctx, data); err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Update(ctx context.Context, data interface{}, filter interface{}) error {
	if _, err := r.collection().UpdateOne(ctx, filter, data); err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Delete(ctx context.Context, filter interface{}) error {
	if _, err := r.collection().DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
