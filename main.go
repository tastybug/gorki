package main

import (
	"bloggo/postable"
	"bloggo/templating"
	"fmt"
	"path/filepath"
)

const workDir string = "testdata"

func main() {
	// write post files
	for fileName, post := range postable.CollectPostables(workDir) {
		fmt.Printf("file %s -> %+v\n", fileName, post.Title)
		page := templating.CreateBlogPostPage(
			post,
			filepath.Join(workDir, `templates`))
		templating.WritePage(workDir, page)
	}

	// write main files
	templating.CreateMainPages(workDir)
}
