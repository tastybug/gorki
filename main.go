package main

import (
	"bloggo/postable"
	"bloggo/templating"
	"fmt"
	"path/filepath"
)

const workDir string = "testdata"

func main() {
	for fileName, post := range postable.CollectPostables(workDir) {
		fmt.Printf("file %s -> %+v\n", fileName, post)
		page := templating.CreateBlogPostPage(
			post,
			filepath.Join(workDir, `templates`))
		templating.WritePage(workDir, page)
	}
}
