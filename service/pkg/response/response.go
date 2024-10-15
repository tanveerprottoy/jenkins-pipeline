package response

import (
	"encoding/json"
	"net/http"
	"github.com/tanveerprottoy/jenkins-pipeline/service/pkg/constant"
)

type Response[T any] struct {
	Data T `json:"data"`
}

type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Errors any `json:"errors"`
}

func BuildData[T any](payload T) Response[T] {
	return Response[T]{Data: payload}
}

func RespondError(w http.ResponseWriter, code int, payload any) (int, error) {
	w.WriteHeader(code)
	res, errs := json.Marshal(payload)
	if errs != nil {
		// log failed to marshal
		return w.Write([]byte(constant.InternalServerError))
	}
	return w.Write(res)
}

func Respond(w http.ResponseWriter, code int, payload any) (int, error) {
	res, err := json.Marshal(payload)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, ErrorResponse{Errors: []any{"an error occured"}})
		return -1, err
	}
	w.WriteHeader(code)
	return w.Write(res)
}
