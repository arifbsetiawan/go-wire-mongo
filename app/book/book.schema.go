package book

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID  primitive.ObjectID `bson:"author_id"`
	Title     string             `bson:"title"`
	Year      string             `bson:"year"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
