package processors

import (
	"fmt"
	"golang-log-server/internals/app/models"
	"golang-log-server/internals/app/storages"
)

type LogProcessor struct {
	storage *storages.LogStorage
}

func NewLogProcessor(storage *storages.LogStorage) *LogProcessor {
	return &LogProcessor{storage: storage}
}

func (processor *LogProcessor) WriteLog(log *models.Log) error {
	file := processor.storage.Storage
	if _, err := file.Write([]byte(fmt.Sprintf("%v\n", log.Log))); err != nil {
		file.Close()
		return err
	}

	file.Sync()
	return nil
}
