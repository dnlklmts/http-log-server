package handlers

import (
	"errors"
	"net/http"
)

func NotFound(wrt http.ResponseWriter, req *http.Request) {
	CreateRespError(wrt, errors.New("not found"), http.StatusNotFound)
}
