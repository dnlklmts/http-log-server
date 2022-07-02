package processors

import (
	"fmt"
	"golang-log-server/internals/app/models"
	"golang-log-server/internals/app/storages"
	"os"
)

type LogProcessor struct {
	storage *storages.LogStorage
}

func NewLogProcessor(storage *storages.LogStorage) *LogProcessor {
	return &LogProcessor{storage: storage}
}

func (processor *LogProcessor) WriteLog(log *models.Log) error {
	file := processor.storage.Storage
	input := []byte(fmt.Sprintf("%v\n", log.Log))
	res, err := file.Write(input)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "write %d bytes\n", res)
	file.Sync()
	return nil
}
