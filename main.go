package main

import (
	"bloggo/templating"
	"fmt"
	"path/filepath"
)

const workDir string = "testdata"

func main() {
	// write post files
	for fileName, post := range CollectPostables(workDir) {
		fmt.Printf("file %s -> %+v\n", fileName, post.Title)
		page := templating.PublishPost(
			post,
			filepath.Join(workDir, `templates`))
		templating.WritePage(workDir, page)
	}

	// write main files
	templating.PublishOtherPages(workDir)
}
