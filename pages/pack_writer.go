package pages

import (
	"bufio"
	"gorki/util"
	"log"
	"os"
	"path/filepath"
)

func WriteContentPack(pack ContentPack) {
	targetDir := GetTargetRootDirectory()
	if pack.FolderName != `` {
		util.CreateDirIfNotExisting(filepath.Join(targetDir, pack.FolderName))
	}
	f, err := os.Create(filepath.Join(targetDir, pack.FolderName, pack.FileName))
	util.PanicOnError(err)
	defer f.Close()
	fileWriter := bufio.NewWriter(f)
	_, err = fileWriter.Write([]byte(pack.HtmlContent))
	util.PanicOnError(err)
	err = fileWriter.Flush()
	util.PanicOnError(err)

	for _, asset := range pack.assets {
		writeAsset(targetDir, asset)
	}
}

func writeAsset(targetRoot string, asset Asset) {
	log.Printf("Writing asset %s\n", asset.CopyFromPath)
	var writeToPath string
	if asset.FolderName != `` {
		writeToPath = filepath.Join(targetRoot, asset.FolderName, asset.FileName)
		util.CreateDirIfNotExisting(filepath.Join(targetRoot, asset.FolderName))
	} else {
		writeToPath = filepath.Join(targetRoot, asset.FileName)
	}
	util.CopyFile(asset.CopyFromPath, writeToPath)
}
