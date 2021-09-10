package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

func writeContentPack(settings Settings, pack renderedPage) {
	targetDir := settings.TargetRoot
	if pack.FolderName != `` {
		CreateDirIfNotExisting(filepath.Join(targetDir, pack.FolderName))
	}
	f, err := os.Create(filepath.Join(targetDir, pack.FolderName, pack.FileName))
	PanicOnError(err)
	defer CloseFile(*f)
	fileWriter := bufio.NewWriter(f)
	_, err = fileWriter.Write([]byte(pack.HtmlContent))
	PanicOnError(err)
	err = fileWriter.Flush()
	PanicOnError(err)

	for _, asset := range pack.assets {
		writeAsset(targetDir, asset)
	}
}

func writeAsset(targetRoot string, asset asset) {
	log.Printf("Writing asset %s/%s\n", asset.FolderName, asset.FileName)
	var writeToPath string
	if asset.FolderName != `` {
		writeToPath = filepath.Join(targetRoot, asset.FolderName, asset.FileName)
		CreateDirIfNotExisting(filepath.Join(targetRoot, asset.FolderName))
	} else {
		writeToPath = filepath.Join(targetRoot, asset.FileName)
	}
	CopyFile(asset.CopyFromPath, writeToPath)
}

func CopyFile(sourcePath, destinationPath string) {
	in, err := os.Open(sourcePath)
	PanicOnError(err)
	defer CloseFile(*in)

	out, err := os.Create(destinationPath)
	PanicOnError(err)

	_, err = io.Copy(out, in)
	PanicOnError(err)
	defer PanicOnError(out.Close())
}

func CreateDirIfNotExisting(path string) {
	if !PathExists(path) {
		PanicOnError(os.MkdirAll(path, 0740))
	}
}
