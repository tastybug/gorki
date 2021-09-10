package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	settings, err := newSettings()
	if err != nil {
		log.Fatalf("Fatal: %v", err)
	}

	gorkify(settings)
	printMemUsage()
}

// https://golangcode.com/print-the-current-memory-usage/
func printMemUsage() {
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

func gorkify(settings Settings) {
	publishablePages := collectAllBundles(settings)

	for _, pack := range renderPages(settings, publishablePages) {
		log.Printf("Writing bundle %s/%s\n", pack.FolderName, pack.FileName)
		writeContentPack(settings, pack)
	}

	log.Println("Finished generation.")
}
