package pages

import (
	"bloggo/util"
	"bufio"
	"log"
	"os"
	"path/filepath"
)

type WritableContent struct {
	Path        string // can be a file name or subpath in target
	HtmlContent string
	assets      map[string]Asset
}

type Asset struct {
	Filename     string
	Context      string
	CopyFromPath string
}

func WriteContent(targetDir string, page WritableContent) {
	f, err := os.Create(filepath.Join(targetDir, page.Path))
	util.PanicOnError(err)
	defer f.Close()
	fileWriter := bufio.NewWriter(f)
	_, err = fileWriter.Write([]byte(page.HtmlContent))
	util.PanicOnError(err)
	err = fileWriter.Flush()
	util.PanicOnError(err)

	if page.assets != nil {
		for _, asset := range page.assets {
			WriteAsset(targetDir, asset)
		}
	}
}

func WriteAsset(targetDir string, asset Asset) {
	log.Printf("Writing asset %s\n", asset.CopyFromPath)
	var targetPath string
	if asset.Context != `` {
		targetPath = filepath.Join(targetDir, asset.Context, asset.Filename)
		err := os.Mkdir(filepath.Join(targetDir, asset.Context), 0740)
		util.PanicOnError(err)
	} else {
		targetPath = filepath.Join(targetDir, asset.Filename)
	}
	util.CopyFile(asset.CopyFromPath, targetPath)
}
