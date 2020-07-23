package main

import (
	"bloggo/proc"
	"bloggo/util"
	"fmt"
	"path/filepath"
)

const workDir string = "testdata"
const targetDir string = "target"

func main() {
	util.PrepareTargetFolder(targetDir)
	templatesFolder := filepath.Join(workDir, `templates`)

	for _, postable := range proc.CollectPostables(workDir) {
		writable := proc.ToWritableContent(postable, templatesFolder)
		fmt.Printf("Writing %s\n", writable.Path)
		proc.WriteContent(targetDir, writable)
	}

	for _, writable := range proc.CollectOtherContent(templatesFolder) {
		fmt.Printf("Writing %s\n", writable.Path)
		proc.WriteContent(targetDir, writable)
	}
	for _, asset := range proc.CollectAssets(workDir, targetDir) {
		fmt.Printf("Writing %s\n", asset.TargetPath)
		proc.WriteAsset(asset)
	}
}
