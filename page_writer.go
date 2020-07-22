package main

import (
	"bloggo/util"
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func WritePage(targetDir string, page Page) {
	f, err := os.Create(filepath.Join(targetDir, getSafeFileName(page)))
	util.PanicOnError(err)
	defer f.Close()
	fileWriter := bufio.NewWriter(f)
	_, err = fileWriter.Write([]byte(page.HtmlContent))
	fileWriter.Flush()
	util.PanicOnError(err)
}

func getSafeFileName(page Page) string {
	return strings.ReplaceAll(page.Postable.Title, " ", "-") + ".html"
}
