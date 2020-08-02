package main

import (
	"gorki/pages"
	"gorki/util"
	"log"
)

func main() {
	settings := util.GetSettings()
	log.Printf("Reading from '%s', writing to '%s'.", settings.SiteRoot, settings.TargetRoot)
	util.PrepareTargetFolder(settings.TargetRoot)

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
