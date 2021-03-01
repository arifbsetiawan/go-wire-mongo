package book

import (
	"encoding/json"
	"go-wire-mongo/shared"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type BookController struct {
	bookService IBookService
	response    *shared.Response
}

func ProvideBookController(BookService BookService) BookController {
	return BookController{
		bookService: &BookService,
	}
}

func (c *BookController) GetIndex(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("authorId")
	dto := IndexDTO{AuthorID: ID}
	books, err := c.bookService.Index(dto)
	if err != nil {
		if err == shared.ErrNilAuthor {
			c.response.ResponseError(w, http.StatusNotFound, err)
		} else {
			c.response.ResponseError(w, http.StatusInternalServerError, err)
		}
		return
	}

	result := ToIndexReponse(*books)

	c.response.ResponseJSON(w, http.StatusOK, result)
}

func (c *BookController) GetShow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	AuthorID := r.URL.Query().Get("authorId")
	dto := ShowDTO{ID: ID, AuthorID: AuthorID}
	if err := shared.Validator(dto, ""); err != nil {
		c.response.ResponseErrValidation(w, err)
		return
	}

	author, err := c.bookService.Show(dto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.response.ResponseError(w, http.StatusNotFound, shared.ErrNilBook)
		} else if err == shared.ErrNilAuthor {
			c.response.ResponseError(w, http.StatusNotFound, err)
		} else {
			c.response.ResponseError(w, http.StatusInternalServerError, err)
		}
		return
	}

	result := ToShowResponse(*author)

	c.response.ResponseJSON(w, http.StatusOK, result)
}

func (c *BookController) PostStore(w http.ResponseWriter, r *http.Request) {
	dto := StoreDTO{}
	json.NewDecoder(r.Body).Decode(&dto)
	if err := shared.Validator(dto, ""); err != nil {
		c.response.ResponseErrValidation(w, err)
		return
	}

	if err := c.bookService.Store(dto); err != nil {
		c.response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result := ToStoreResponse()

	c.response.ResponseJSON(w, http.StatusCreated, result)
}

func (c *BookController) PutUpdate(w http.ResponseWriter, r *http.Request) {
	dto := UpdateDTO{}
	json.NewDecoder(r.Body).Decode(&dto)
	if err := shared.Validator(dto, ""); err != nil {
		c.response.ResponseErrValidation(w, err)
		return
	}

	if err := c.bookService.Update(dto); err != nil {
		c.response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result := ToUpdateResponse()

	c.response.ResponseJSON(w, http.StatusOK, result)
}

func (c *BookController) DeleteDestroy(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	AuthorID := r.URL.Query().Get("authorId")
	dto := DestroyDTO{ID: ID, AuthorID: AuthorID}
	if err := shared.Validator(dto, ""); err != nil {
		c.response.ResponseErrValidation(w, err)
		return
	}

	if err := c.bookService.Destroy(dto); err != nil {
		c.response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	result := ToDestroyResponse()

	c.response.ResponseJSON(w, http.StatusOK, result)
}
