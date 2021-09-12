package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	g "github.com/tastybug/gorki/internal/gorkiconcurrent"
)

func main() {
	runGorkify()
	printMemUsage()
}

func runGorkify() error {

	var siteDir = flag.String("s", `site`, "site directory")
	var targetDir = flag.String("t", `target`, "output directory name; will be purged if already existing")

	flag.Parse()

	settings := g.Settings{
		SiteRoot:      *siteDir,
		TargetRoot:    filepath.Join(*siteDir, *targetDir),
		TemplatesRoot: filepath.Join(*siteDir, `templates`),
		ArticlesRoot:  filepath.Join(*siteDir, `posts`),
	}

	if !g.FileExists(settings.SiteRoot) {
		return errors.New("Site root expected at '" + settings.SiteRoot + "' but path does not exist.")
	}
	if !g.FileExists(settings.TemplatesRoot) {
		return errors.New("Templates expected at '" + settings.TemplatesRoot + "' but path does not exist.")
	}
	if !g.FileExists(settings.ArticlesRoot) {
		return errors.New("Articles expected at '" + settings.ArticlesRoot + "' but path does not exist.")
	}

	createOrPurgeTargetFolder(settings.TargetRoot)

	g.Gorkify(settings)

	return nil
}

func createOrPurgeTargetFolder(dir string) {
	if _, err := os.Stat(dir); err == nil {
		log.Printf("Emptying target folder '%s'.\n", dir)
		for _, toBeRemoved := range ListFilesAndDirs(dir) {
			name := toBeRemoved.Name()
			if err := os.RemoveAll(filepath.Join(dir, name)); err != nil {
				panic(err)
			}
		}
	} else {
		log.Printf("Creating non-existent target folder '%s'.\n", dir)
		if err := os.Mkdir(dir, os.FileMode(0740)); err != nil {
			panic(err)
		}
	}
}

func ListFilesAndDirs(dir string) []os.FileInfo {
	if allFiles, err := ioutil.ReadDir(dir); err != nil {
		panic(err)
	} else {
		return allFiles
	}
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
