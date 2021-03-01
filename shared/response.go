package shared

import (
	"encoding/json"
	"net/http"
)

type Response struct{}

func (r *Response) ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (r *Response) ResponseError(w http.ResponseWriter, code int, errors ...error) {
	var errMessage []string
	for _, v := range errors {
		errMessage = append(errMessage, v.Error())
	}

	response := ToErrorResponse(ToDefaultResponse(code, false, ""), errMessage)
	r.ResponseJSON(w, code, response)
}

func (r *Response) ResponseErrValidation(w http.ResponseWriter, errors []error) {
	r.ResponseError(w, http.StatusBadRequest, errors...)
}
