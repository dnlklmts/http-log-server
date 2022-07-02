package processors

import (
	"fmt"
	"golang-log-server/internals/app/storages"
	"os"
)

type LogProcessor struct {
	storage *storages.LogStorage
}

func NewLogProcessor(storage *storages.LogStorage) *LogProcessor {
	return &LogProcessor{storage: storage}
}

func (processor *LogProcessor) WriteLog(input string) error {
	res, err := processor.storage.Storage.WriteString(input)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "write %d bytes\n", res)
	return nil
}
