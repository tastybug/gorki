package main

import (
	"gorki/gorkify"
	"gorki/util"
	"log"
)

func main() {
	settings := util.GetSettings()
	log.Printf("Reading from '%s', writing to '%s'.", settings.SiteRoot, settings.TargetRoot)
	util.CreateOrPurgeTargetFolder(settings.TargetRoot)

	gorkify.Gorkify(settings)

	util.PrintMemUsage()
}
