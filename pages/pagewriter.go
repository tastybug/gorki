package pages

import (
	"bloggo/util"
	"bufio"
	"os"
	"path/filepath"
)

type WritableContent struct {
	HtmlContent string
	Path        string // can be a file name or subpath in target
}

type Asset struct {
	SourcePath string
	TargetPath string
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
}

func WriteAsset(asset Asset) {
	util.CopyFile(asset.SourcePath, asset.TargetPath)
}
