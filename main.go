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

	publishablePages := pages.CollectArticles(siteDir)
	publishablePages = append(publishablePages, pages.CollectMains(filepath.Join(siteDir, `templates`))...)

	//for _, article := range pages.TurnArticlesIntoContentPack(articles, siteDir) {
	//	log.Printf("Writing article %s\n", article.FileName)
	//	pages.WriteContentPack(targetDir, article)
	//	log.Println("Done")
	//}

	for _, pack := range pages.CreatePacks(publishablePages, siteDir) {
		log.Printf("Writing page %s\n", pack.FileName)
		pages.WriteContentPack(targetDir, pack)
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
