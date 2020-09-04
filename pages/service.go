package pages

import (
	"gorki/util"
	"log"
)

func Build(settings util.Settings) {
	publishablePages := collectPages(settings)

	for _, pack := range RenderPages(settings, publishablePages) {
		log.Printf("Writing page %s/%s\n", pack.FolderName, pack.FileName)
		WriteContentPack(settings, pack)
	}

	log.Println("Finished generation.")
}
