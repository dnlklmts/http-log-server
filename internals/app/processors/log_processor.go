package processors

import (
	"bufio"
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
	input := fmt.Sprintf("%v\n", log)
	buf := bufio.NewWriter(processor.storage.Storage)
	res, err := buf.WriteString(input)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "write %d bytes\n", res)
	buf.Flush()
	return nil
}
