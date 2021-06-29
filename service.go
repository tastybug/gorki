package main

import (
	"log"
)

func Gorkify(settings Settings) {
	publishablePages := collectAllBundles(settings)

	for _, pack := range renderPages(settings, publishablePages) {
		log.Printf("Writing bundle %s/%s\n", pack.FolderName, pack.FileName)
		writeContentPack(settings, pack)
	}

	log.Println("Finished generation.")
}
