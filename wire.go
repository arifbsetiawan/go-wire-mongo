// +build wireinject

package main

import (
	"go-wire-mongo/app"
	"go-wire-mongo/app/author"
	"go-wire-mongo/app/book"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func AppModule() app.AppController {
	wire.Build(AppSet)

	return app.AppController{}
}

func AuthorModule(db *mongo.Database) author.AuthorController {
	wire.Build(AuthorSet)

	return author.AuthorController{}
}

func BookModule(db *mongo.Database) book.BookController {
	wire.Build(BookSet)

	return book.BookController{}
}
