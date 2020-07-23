package main

import (
	"bloggo/pages"
	"bloggo/util"
	"log"
)

const workDir string = "testdata"
const targetDir string = "target"

func main() {

	util.PrepareTargetFolder(targetDir)

	for _, article := range pages.CollectArticlePages(workDir) {
		log.Printf("Writing article %s\n", article.Path)
		pages.WriteContent(targetDir, article)
		log.Println("Done")
	}

	for _, mainPage := range pages.CollectMainPages(workDir) {
		log.Printf("Writing main page %s\n", mainPage.Path)
		pages.WriteContent(targetDir, mainPage)
		log.Println("Done")
	}

	log.Println("Finished generation.")
	util.PrintMemUsage()
}
