package main

import (
	"go-wire-mongo/app"
	"go-wire-mongo/app/author"
	"go-wire-mongo/app/book"

	"github.com/google/wire"
)

var (
	AppSet = wire.NewSet(app.ProvideAppController)

	AuthorSet = wire.NewSet(
		author.ProvideAuthorRepository,
		author.ProvideAuthorService,
		author.ProvideAuthorController,
	)

	BookSet = wire.NewSet(
		AuthorSet,
		book.ProvideBookController,
		book.ProvideBookService,
		book.ProvideBookRepository,
	)
)
