package main

import (
	"gorki/pages"
	"gorki/util"
	"log"
)

func main() {
	siteDir := pages.GetSiteRootDirectory()
	targetDir := pages.GetTargetRootDirectory()

	log.Printf("Using site at '%s', target will be '%s'.", siteDir, targetDir)
	util.PrepareTargetFolder(targetDir)

	publishablePages := pages.CollectArticles()
	publishablePages = append(publishablePages, pages.CollectMains()...)

	for _, pack := range pages.CreatePacks(publishablePages) {
		log.Printf("Writing page %s\n", pack.FileName)
		pages.WriteContentPack(pack)
		log.Println("Done")
	}

	log.Println("Finished generation.")
	util.PrintMemUsage()
}
