package main

import (
	"bloggo/pages"
	"bloggo/util"
	"log"
	"os"
	"path/filepath"
)

const defaultSiteDir string = "site" // relative to CWD
const defaultTargetDirName string = "target"

func main() {
	siteDir := getSiteDirectory()
	targetDir := filepath.Join(siteDir, defaultTargetDirName)
	log.Printf("Using target directory '%s'.", targetDir)
	util.PrepareTargetFolder(targetDir)

	for _, article := range pages.CollectArticlePages(siteDir) {
		log.Printf("Writing article %s\n", article.FileName)
		pages.WriteContent(targetDir, article)
		log.Println("Done")
	}

	for _, mainPage := range pages.CollectMainPages(siteDir) {
		log.Printf("Writing main page %s\n", mainPage.FileName)
		pages.WriteContent(targetDir, mainPage)
		log.Println("Done")
	}

	for _, globalAsset := range pages.CollectGlobalAssets(siteDir) {
		log.Printf("Writing shared asset %s\n", globalAsset.CopyFromPath)
		pages.WriteAsset(targetDir, globalAsset)
		log.Println("Done")
	}

	log.Println("Finished generation.")
	util.PrintMemUsage()
}

func getSiteDirectory() string {
	siteDir := defaultSiteDir
	if len(os.Args) == 2 {
		siteDir = os.Args[1]
		log.Printf("Using site directory '%s'.", siteDir)
	} else if len(os.Args) > 2 {
		log.Printf("Usage: bloggo [path-to-site-directory]\n\nIf omitted, site is expected at $CWD/%s", defaultSiteDir)
		os.Exit(0)
	} else {
		log.Printf("Using default site directory '%s'.", defaultSiteDir)
	}
	return siteDir
}
