package handlers

import (
	"encoding/json"
	"golang-log-server/internals/app/models"
	"golang-log-server/internals/app/processors"
	"net/http"
)

type LogHandler struct {
	processor *processors.LogProcessor
}

func NewLogHandler(processor *processors.LogProcessor) *LogHandler {
	return &LogHandler{processor: processor}
}

func (handler *LogHandler) WriteLog(wrt http.ResponseWriter, req *http.Request) {
	newLog := new(models.Log)

	err := json.NewDecoder(req.Body).Decode(newLog)
	if err != nil {
		RespError(wrt, err)
		return
	}

	err = handler.processor.WriteLog(newLog)
	if err != nil {
		RespError(wrt, err)
		return
	}

	answer := map[string]interface{}{
		"result": "OK",
		"data":   struct{}{},
	}
	RespOK(wrt, answer)
}
