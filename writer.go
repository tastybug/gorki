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
		ensureDir(filepath.Join(targetDir, pack.FolderName))
	}
	if f, err := os.Create(filepath.Join(targetDir, pack.FolderName, pack.FileName)); err != nil {
		panic(err)
	} else {
		defer closeFile(*f)
		fileWriter := bufio.NewWriter(f)
		if _, err = fileWriter.Write([]byte(pack.HtmlContent)); err != nil {
			panic(err)
		}
		if err = fileWriter.Flush(); err != nil {
			panic(err)
		}
		for _, asset := range pack.assets {
			writeAsset(targetDir, asset)
		}
	}
}

func writeAsset(targetRoot string, asset asset) {
	log.Printf("Writing asset %s/%s\n", asset.FolderName, asset.FileName)
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
	if !fileExists(path) {
		if err := os.MkdirAll(path, 0740); err != nil {
			panic(err)
		}
	}
}
