package gorkiconcurrent

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	ARTICLE_BUNDLE = iota + 1
	STATIC_BUNDLE
)

type discovery struct {
	fi   os.FileInfo
	kind int
}

func discover(articlesRoot, templatesRoot string, outChan chan<- discovery) {

	for _, bundleDir := range listDirs(articlesRoot) {
		outChan <- discovery{bundleDir, ARTICLE_BUNDLE}
	}
	for _, bundleDir := range listDirs(templatesRoot) {
		outChan <- discovery{bundleDir, STATIC_BUNDLE}
	}

	close(outChan)
}

func listDirs(dir string) []os.FileInfo {
	if allFiles, err := ioutil.ReadDir(dir); err != nil {
		panic(err)
	} else {
		return filter(
			allFiles,
			func(file os.FileInfo) bool {
				// must be a directory that does not start with an underscore
				return file.IsDir() && strings.IndexRune(file.Name(), '_') != 0
			})
	}
}

func filter(files []os.FileInfo, test func(os.FileInfo) bool) (result []os.FileInfo) {
	for _, f := range files {
		if test(f) {
			result = append(result, f)
		}
	}
	return
}
