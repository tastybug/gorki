package pages

import (
	"gorki/util"
	"log"
)

func Build(settings util.Settings) {
	publishablePages := collectPages(settings)

	for _, pack := range CreatePacks(settings, publishablePages) {
		log.Printf("Writing page %s\n", pack.FileName)
		WriteContentPack(settings, pack)
		log.Println("Done")
	}

	log.Println("Finished generation.")
}
