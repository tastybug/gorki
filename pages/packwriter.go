package pages

import (
	"bloggo/util"
	"bufio"
	"log"
	"os"
	"path/filepath"
)

type ContentPack struct {
	Folders     string
	FileName    string
	HtmlContent string
	assets      []Asset
}

type Asset struct {
	Filename     string
	Context      string
	CopyFromPath string
}

func WriteContentPack(targetDir string, writable ContentPack) {
	if writable.Folders != `` {
		util.PanicOnError(os.MkdirAll(filepath.Join(targetDir, writable.Folders), 0740))
	}
	f, err := os.Create(filepath.Join(targetDir, writable.Folders, writable.FileName))
	util.PanicOnError(err)
	defer f.Close()
	fileWriter := bufio.NewWriter(f)
	_, err = fileWriter.Write([]byte(writable.HtmlContent))
	util.PanicOnError(err)
	err = fileWriter.Flush()
	util.PanicOnError(err)

	for _, asset := range writable.assets {
		writeAsset(targetDir, asset)
	}
}

func writeAsset(targetDir string, asset Asset) {
	log.Printf("Writing asset %s\n", asset.CopyFromPath)
	var targetPath string
	if asset.Context != `` {
		targetPath = filepath.Join(targetDir, asset.Context, asset.Filename)
		util.CreateDirIfNotExisting(filepath.Join(targetDir, asset.Context))
	} else {
		targetPath = filepath.Join(targetDir, asset.Filename)
	}
	util.CopyFile(asset.CopyFromPath, targetPath)
}
