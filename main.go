package main

import (
	"bloggo/pages"
	"bloggo/util"
	"fmt"
	"path/filepath"
)

const workDir string = "testdata"
const targetDir string = "target"

func main() {
	util.PrepareTargetFolder(targetDir)
	templatesFolder := filepath.Join(workDir, `templates`)

	for _, article := range pages.CollectArticlePages(workDir, templatesFolder) {
		fmt.Printf("Writing article %s\n", article.Path)
		pages.WriteContent(targetDir, article)
	}

	for _, mainPage := range pages.CollectMainPages(templatesFolder) {
		fmt.Printf("Writing main page %s\n", mainPage.Path)
		pages.WriteContent(targetDir, mainPage)
	}

	for _, asset := range pages.CollectAssets(workDir, targetDir) {
		fmt.Printf("Writing asset %s\n", asset.TargetPath)
		pages.WriteAsset(asset)
	}
}
