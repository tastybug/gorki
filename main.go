package main

import (
	"bloggo/util"
	"fmt"
	"os"
	"path/filepath"
)

const workDir string = "testdata"
const targetDir string = "target"

func main() {
	prepareTarget(targetDir)

	// write post files
	for fileName, post := range CollectPostables(workDir) {
		fmt.Printf("file %s -> %+v\n", fileName, post.Title)
		page := PublishPost(
			post,
			filepath.Join(workDir, `templates`))
		WritePage(targetDir, page)
	}

	// write main files
	PublishOtherPages(workDir, targetDir)
}

func prepareTarget(dir string) {
	if _, err := os.Stat(dir); err == nil {
		err := os.RemoveAll(dir)
		util.PanicOnError(err)
	}

	err := os.Mkdir(dir, os.FileMode(0740))
	util.PanicOnError(err)
}
