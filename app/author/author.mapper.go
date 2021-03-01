package author

import (
	"go-wire-mongo/shared"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToIndexResponse(data []Author) IndexReponse {
	return IndexReponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusOK, true, "Index Authors"),
		Data:            ToAuthorDTOs(data),
	}
}

func ToShowResponse(data Author) ShowReponse {
	return ShowReponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusOK, true, "Show Author"),
		Data:            ToAuthorDTO(data),
	}
}

func ToStoreResponse() StoreResponse {
	return StoreResponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusCreated, true, "Author created"),
	}
}

func ToUpdateResponse() UpdateResponse {
	return UpdateResponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusOK, true, "Author updated"),
	}
}

func ToDestroyResponse() DestroyResponse {
	return DestroyResponse{
		DefaultResponse: shared.ToDefaultResponse(http.StatusOK, true, "Author deleted"),
	}
}

func ToAuthorDTOs(data []Author) []AuthorDTO {
	result := make([]AuthorDTO, len(data))

	for i, v := range data {
		result[i] = ToAuthorDTO(v)
	}

	return result
}

func ToAuthorDTO(data Author) AuthorDTO {
	return AuthorDTO{
		ID:      data.ID.Hex(),
		Name:    data.Name,
		Phone:   data.Phone,
		Address: data.Address,
	}
}

func ToAuthor(data StoreDTO) Author {
	return Author{
		ID:        primitive.NewObjectID(),
		Name:      data.Name,
		Phone:     data.Phone,
		Address:   data.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
