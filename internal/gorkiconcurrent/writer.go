package gorkiconcurrent

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

func write(targetDir string, inChan <-chan renderedBundle, outChan chan<- renderedBundle) {
	for renderedBundle := range inChan {
		writeImpl(targetDir, renderedBundle)
		outChan <- renderedBundle
	}
	close(outChan)
}

func writeImpl(targetDir string, page renderedBundle) {
	if page.FolderName != `` {
		ensureDir(filepath.Join(targetDir, page.FolderName))
	}
	if f, err := os.Create(filepath.Join(targetDir, page.FolderName, page.FileName)); err != nil {
		panic(err)
	} else {
		defer closeFile(*f)
		fileWriter := bufio.NewWriter(f)
		if _, err = fileWriter.Write([]byte(page.HtmlContent)); err != nil {
			panic(err)
		}
		if err = fileWriter.Flush(); err != nil {
			panic(err)
		}
		for _, asset := range page.assets {
			writeAsset(targetDir, asset)
		}
	}
}

func writeAsset(targetRoot string, asset asset) {
	var writeToPath string
	if asset.FolderName != `` {
		writeToPath = filepath.Join(targetRoot, asset.FolderName, asset.FileName)
		ensureDir(filepath.Join(targetRoot, asset.FolderName))
	} else {
		writeToPath = filepath.Join(targetRoot, asset.FileName)
	}
	copyFile(asset.CopyFromPath, writeToPath)
}

func copyFile(sourcePath, destinationPath string) {
	if in, err := os.Open(sourcePath); err != nil {
		panic(err)
	} else {
		defer closeFile(*in)
		if out, err := os.Create(destinationPath); err != nil {
			panic(err)
		} else {
			defer func() {
				if err := out.Close(); err != nil {
					panic(err)
				}
			}()
			if _, err = io.Copy(out, in); err != nil {
				panic(err)
			}
		}
	}
}

func ensureDir(path string) {
	if !FileExists(path) {
		if err := os.MkdirAll(path, 0740); err != nil {
			panic(err)
		}
	}
}
