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

	pages.Build(settings)

	util.PrintMemUsage()
}
