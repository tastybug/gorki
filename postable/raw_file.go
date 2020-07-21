package postable

import (
	"bloggo/util"
	"bufio"
	. "io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func listMarkdownFiles(dir string) []os.FileInfo {
	allFiles, err := ReadDir(dir)
	util.PanicOnError(err)

	onlyMarkdown := func(file os.FileInfo) bool { return strings.HasSuffix(file.Name(), ".md") }
	return filter(allFiles, onlyMarkdown)
}

func readFileContent(dir string, fileName string) string {
	file, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	return content
}

func filter(files []os.FileInfo, test func(os.FileInfo) bool) (ret []os.FileInfo) {
	for _, f := range files {
		if test(f) {
			ret = append(ret, f)
		}
	}
	return
}
