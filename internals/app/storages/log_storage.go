package storages

import (
	"log"
	"os"
)

type LogStorage struct {
	Storage *os.File
}

func NewLogStorage(path string) *LogStorage {
	storage := new(LogStorage)
	var err error

	storage.Storage, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}

	return storage
}
