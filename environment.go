package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func CreateOrPurgeTargetFolder(dir string) {
	if _, err := os.Stat(dir); err == nil {
		log.Printf("Emptying target folder '%s'.\n", dir)
		for _, toBeRemoved := range ListFilesAndDirs(dir) {
			name := toBeRemoved.Name()
			PanicOnError(os.RemoveAll(filepath.Join(dir, name)))
		}
	} else {
		log.Printf("Creating non-existent target folder '%s'.\n", dir)
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
