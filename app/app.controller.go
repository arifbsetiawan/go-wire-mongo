package app

import (
	"go-wire-mongo/shared"
	"net/http"
)

type AppController struct {
	res shared.Response
}

func ProvideAppController() AppController {
	return AppController{}
}

func (c *AppController) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"status": "work"}
	c.res.ResponseJSON(w, http.StatusOK, data)
}
