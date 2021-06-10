package gorkify

import (
	"gorki/util"
	"log"
)

func Gorkify(settings util.Settings) {
	publishablePages := collectPages(settings)

	for _, pack := range RenderPages(settings, publishablePages) {
		log.Printf("Writing bundle %s/%s\n", pack.FolderName, pack.FileName)
		WriteContentPack(settings, pack)
	}

	log.Println("Finished generation.")
}
