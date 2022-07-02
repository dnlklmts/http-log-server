package storages

import (
	"fmt"
	"log"
	"os"
)

type LogStorage struct {
	Storage *os.File
}

func NewLogStorage(path string) *LogStorage {
	storage := new(LogStorage)
	var err error

	if !fileExists(path) {
		storage.Storage, err = os.Create(path)
		if err != nil {
			log.Fatalln("failed to create file:", err)
		}
	}

	storage.Storage, err = os.Open(path)
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}

	return storage
}

func (storage *LogStorage) WriteLog(input string) {
	res, err := storage.Storage.WriteString(input)
	if err != nil {
		log.Fatalln("failed to write into file:", err)
	}
	fmt.Fprintf(os.Stdout, "write %d bytes\n", res)
}
