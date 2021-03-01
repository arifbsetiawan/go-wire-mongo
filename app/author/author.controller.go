package author

import (
	"encoding/json"
	"go-wire-mongo/shared"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorController struct {
	authorService IAuthorService
	response      *shared.Response
}

func ProvideAuthorController(AuthorService AuthorService) AuthorController {
	return AuthorController{
		authorService: &AuthorService,
	}
}

func (c *AuthorController) GetIndex(w http.ResponseWriter, r *http.Request) {
	authors, err := c.authorService.Index()
	if err != nil {
		c.response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result := ToIndexResponse(*authors)

	c.response.ResponseJSON(w, http.StatusOK, result)
}

func (c *AuthorController) GetShow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	dto := ShowDTO{ID: ID}
	if err := shared.Validator(dto, ""); err != nil {
		c.response.ResponseErrValidation(w, err)
		return
	}

	author, err := c.authorService.Show(dto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.response.ResponseError(w, http.StatusNotFound, shared.ErrNilAuthor)
		} else {
			c.response.ResponseError(w, http.StatusInternalServerError, err)
		}
		return
	}

	result := ToShowResponse(*author)

	c.response.ResponseJSON(w, http.StatusOK, result)
}

func (c *AuthorController) PostStore(w http.ResponseWriter, r *http.Request) {
	dto := StoreDTO{}
	json.NewDecoder(r.Body).Decode(&dto)
	if err := shared.Validator(dto, ""); err != nil {
		c.response.ResponseErrValidation(w, err)
		return
	}

	if err := c.authorService.Store(dto); err != nil {
		c.response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result := ToStoreResponse()

	c.response.ResponseJSON(w, http.StatusCreated, result)
}

func (c *AuthorController) PutUpdate(w http.ResponseWriter, r *http.Request) {
	dto := UpdateDTO{}
	json.NewDecoder(r.Body).Decode(&dto)
	if err := shared.Validator(dto, ""); err != nil {
		c.response.ResponseErrValidation(w, err)
		return
	}

	if err := c.authorService.Update(dto); err != nil {
		c.response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result := ToUpdateResponse()

	c.response.ResponseJSON(w, http.StatusOK, result)
}

func (c *AuthorController) DeleteDestroy(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	dto := DestroyDTO{ID: ID}
	if err := shared.Validator(dto, ""); err != nil {
		c.response.ResponseErrValidation(w, err)
		return
	}

	if err := c.authorService.Destroy(dto); err != nil {
		c.response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result := ToDestroyResponse()

	c.response.ResponseJSON(w, http.StatusOK, result)
}
