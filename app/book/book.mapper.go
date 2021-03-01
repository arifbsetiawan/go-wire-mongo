package book

import (
	"go-wire-mongo/app/author"
	"go-wire-mongo/shared"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToIndexReponse(data []BookDTO) IndexReponse {
	return IndexReponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusOK, true, "Index Books"),
		Data:            data,
	}
}

func ToShowResponse(data BookDTO) ShowReponse {
	return ShowReponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusOK, true, "Show Book"),
		Data:            data,
	}
}

func ToStoreResponse() StoreResponse {
	return StoreResponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusCreated, true, "Book created"),
	}
}

func ToUpdateResponse() UpdateResponse {
	return UpdateResponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusOK, true, "Book updated"),
	}
}

func ToDestroyResponse() DestroyResponse {
	return DestroyResponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusOK, true, "Book deleted"),
	}
}

func ToBookDTO(data Book, authorData author.Author) BookDTO {
	return BookDTO{
		ID:    data.ID.Hex(),
		Title: data.Title,
		Year:  data.Year,
		Author: author.AuthorDTO{
			ID:      authorData.ID.Hex(),
			Name:    authorData.Name,
			Phone:   authorData.Phone,
			Address: authorData.Address,
		},
	}
}

func ToBookDTOs(data []Book, author author.Author) []BookDTO {
	books := make([]BookDTO, len(data))

	for i, v := range data {
		books[i] = ToBookDTO(v, author)
	}

	return books
}

func ToBook(data StoreDTO) Book {
	AuthorID, _ := primitive.ObjectIDFromHex(data.AuthorID)
	return Book{
		ID:        primitive.NewObjectID(),
		AuthorID:  AuthorID,
		Title:     data.Title,
		Year:      data.Year,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ToAuthorShowDTO(data string) author.ShowDTO {
	return author.ShowDTO{
		ID: data,
	}
}
