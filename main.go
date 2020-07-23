package main

import (
	"bloggo/util"
	"fmt"
	"path/filepath"
)

const workDir string = "testdata"
const targetDir string = "target"

func main() {
	util.PrepareTargetFolder(targetDir)

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
