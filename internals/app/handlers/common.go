package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespError(wrt http.ResponseWriter, err error) {
	CreateRespError(wrt, err, http.StatusBadRequest)
}

func CreateRespError(wrt http.ResponseWriter, err error, httpsStatus int) {
	answer := map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	resp, _ := json.Marshal(answer)
	wrt.Header().Set("Content-Type", "application/json; charset=utf-8")
	wrt.Header().Set("X-Content-Type-Options", "nosniff")
	wrt.WriteHeader(httpsStatus)

	fmt.Fprintln(wrt, string(resp))
}

func RespOK(wrt http.ResponseWriter, data map[string]interface{}) {
	resp, _ := json.Marshal(data)
	wrt.Header().Set("Content-Type", "application/json; charset=utf-8")
	wrt.WriteHeader(http.StatusOK)
	fmt.Fprintln(wrt, string(resp))

}
