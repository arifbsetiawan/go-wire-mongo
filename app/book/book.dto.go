package book

import (
	"go-wire-mongo/app/author"
	"go-wire-mongo/shared"
)

type IndexDTO struct {
	AuthorID string `json:"authorId" validate:"required,len=24"`
}

type ShowDTO struct {
	ID       string `json:"id" validate:"required,len=24"`
	AuthorID string `json:"authorId" validate:"required,len=24"`
}

type StoreDTO struct {
	AuthorID string `json:"authorId" validate:"required,len=24"`
	Title    string `json:"title" validate:"required"`
	Year     string `json:"year" validate:"required"`
}

type UpdateDTO struct {
	ID       string `json:"id" validate:"required,len=24"`
	AuthorID string `json:"authorId" validate:"required,len=24"`
	Title    string `json:"title" validate:"required"`
	Year     string `json:"year" validate:"required"`
}

type DestroyDTO struct {
	ID       string `json:"id" validate:"required,len=24"`
	AuthorID string `json:"authorId" validate:"required,len=24"`
}

type IndexReponse struct {
	shared.DefaultResponse
	Data []BookDTO `json:"data"`
}

type ShowReponse struct {
	shared.DefaultResponse
	Data BookDTO `json:"data"`
}

type StoreResponse struct {
	shared.DefaultResponse
}

type UpdateResponse struct {
	shared.DefaultResponse
}

type DestroyResponse struct {
	shared.DefaultResponse
}

type BookDTO struct {
	ID     string           `json:"id,omitempty"`
	Title  string           `json:"title" validate:"required"`
	Year   string           `json:"year" validate:"required"`
	Author author.AuthorDTO `json:"author,omitempty" validate:""`
}
