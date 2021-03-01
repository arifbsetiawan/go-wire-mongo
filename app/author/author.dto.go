package author

import "go-wire-mongo/shared"

type ShowDTO struct {
	ID string `json:"id,omitempty" validate:"required,len=24"`
}

type StoreDTO struct {
	AuthorDTO
}

type UpdateDTO struct {
	ID string `json:"id,omitempty" validate:"required,len=24"`
	AuthorDTO
}

type DestroyDTO struct {
	ID string `json:"id,omitempty" validate:"required,len=24"`
}

type IndexReponse struct {
	shared.DefaultResponse
	Data []AuthorDTO `json:"data"`
}

type ShowReponse struct {
	shared.DefaultResponse
	Data AuthorDTO `json:"data"`
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

type AuthorDTO struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}
