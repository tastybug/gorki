package main

import (
	"log"
)

func main() {
	settings := GetSettings()
	log.Printf("Reading from '%s', writing to '%s'.", settings.SiteRoot, settings.TargetRoot)
	CreateOrPurgeTargetFolder(settings.TargetRoot)

	Gorkify(settings)

	PrintMemUsage()
}
