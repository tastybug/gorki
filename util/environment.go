package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func PrepareTargetFolder(dir string) {
	log.Printf("Preparing target folder '%s'.\n", dir)
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

// https://golangcode.com/print-the-current-memory-usage/
func PrintMemUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(memStats.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(memStats.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(memStats.Sys))
	fmt.Printf("\tNumGC = %v\n", memStats.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
