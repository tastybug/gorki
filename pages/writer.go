package pages

import (
	"bufio"
	"gorki/util"
	"log"
	"os"
	"path/filepath"
)

func WriteContentPack(settings util.Settings, pack renderedPage) {
	targetDir := settings.TargetRoot
	if pack.FolderName != `` {
		util.CreateDirIfNotExisting(filepath.Join(targetDir, pack.FolderName))
	}
	f, err := os.Create(filepath.Join(targetDir, pack.FolderName, pack.FileName))
	util.PanicOnError(err)
	defer util.CloseFile(*f)
	fileWriter := bufio.NewWriter(f)
	_, err = fileWriter.Write([]byte(pack.HtmlContent))
	util.PanicOnError(err)
	err = fileWriter.Flush()
	util.PanicOnError(err)

	for _, asset := range pack.assets {
		writeAsset(targetDir, asset)
	}
}

func writeAsset(targetRoot string, asset asset) {
	log.Printf("Writing asset %s/%s\n", asset.FolderName, asset.FileName)
	var writeToPath string
	if asset.FolderName != `` {
		writeToPath = filepath.Join(targetRoot, asset.FolderName, asset.FileName)
		util.CreateDirIfNotExisting(filepath.Join(targetRoot, asset.FolderName))
	} else {
		writeToPath = filepath.Join(targetRoot, asset.FileName)
	}
	util.CopyFile(asset.CopyFromPath, writeToPath)
}
