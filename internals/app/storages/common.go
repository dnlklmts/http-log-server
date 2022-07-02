package storages

import "os"

func fileExists(path string) bool {
	file, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}
