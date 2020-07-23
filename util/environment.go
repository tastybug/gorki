package util

import (
	"log"
	"os"
	"path/filepath"
)

func PrepareTargetFolder(dir string) {
	log.Println("Preparing target folder..")
	if _, err := os.Stat(dir); err == nil {
		for _, toBeRemoved := range ListFilesWithSuffix(dir, ``) {
			name := toBeRemoved.Name()
			log.Println("Removing " + filepath.Join(dir, name) + " from target.")
			PanicOnError(os.RemoveAll(filepath.Join(dir, name)))
		}
	} else {
		log.Println("Creating non-existent target folder.")
		err := os.Mkdir(dir, os.FileMode(0740))
		PanicOnError(err)
	}
}
