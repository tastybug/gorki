package main

import (
	"bloggo/pages"
	"bloggo/util"
	"fmt"
)

const workDir string = "testdata"
const targetDir string = "target"

func main() {
	util.PrepareTargetFolder(targetDir)

	for _, article := range pages.CollectArticlePages(workDir) {
		fmt.Printf("Writing article %s\n", article.Path)
		pages.WriteContent(targetDir, article)
		fmt.Println("Done")
	}

	for _, mainPage := range pages.CollectMainPages(workDir) {
		fmt.Printf("Writing main page %s\n", mainPage.Path)
		pages.WriteContent(targetDir, mainPage)
		fmt.Println("Done")
	}
}
