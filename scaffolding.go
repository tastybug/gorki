package main

import (
	"bloggo/util"
	"bufio"
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func PublishOtherPages(workDir, targetDir string) {

	writeMainPage(
		targetDir,
		[]string{
			filepath.Join(workDir, `templates/about`, "_about.html"),
			filepath.Join(workDir, `templates`, "footer.html"),
			filepath.Join(workDir, `templates`, "navigation.html"),
			filepath.Join(workDir, `templates`, "head.html"),
		},
		"about.html")
	writeMainPage(
		targetDir,
		[]string{
			filepath.Join(workDir, `templates/index`, "_index.html"),
			filepath.Join(workDir, `templates`, "footer.html"),
			filepath.Join(workDir, `templates`, "navigation.html"),
			filepath.Join(workDir, `templates`, "head.html"),
		},
		"index.html")

	copyAssets(workDir, targetDir)
}

func writeMainPage(targetDir string, paths []string, mainFileName string) {
	tmpl := template.Must(template.ParseFiles(paths...))

	var b bytes.Buffer
	tmpl.Execute(&b, nil)

	f, err := os.Create(filepath.Join(targetDir, mainFileName))
	util.PanicOnError(err)
	defer f.Close()
	fileWriter := bufio.NewWriter(f)
	_, err = fileWriter.Write([]byte(b.String()))
	fileWriter.Flush()
	util.PanicOnError(err)
}

func copyAssets(workDir, targetDir string) {
	assetFolder := filepath.Join(workDir, `templates`, `assets`)

	allFiles, err := ioutil.ReadDir(assetFolder)
	util.PanicOnError(err)

	for _, fileInfo := range allFiles {
		copyFile(
			filepath.Join(assetFolder, fileInfo.Name()),
			filepath.Join(targetDir, fileInfo.Name()),
		)
	}
}

func copyFile(src, destination string) {
	in, err := os.Open(src)
	util.PanicOnError(err)
	defer in.Close()

	out, err := os.Create(destination)
	util.PanicOnError(err)

	_, err = io.Copy(out, in)
	util.PanicOnError(err)
	defer util.PanicOnError(out.Close())
}
