package templating

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func WritePage(workDir string, page Page) {
	f, err := os.Create(filepath.Join(workDir, `target`, getSafeFileName(page)))
	panicOnError(err)
	defer f.Close()
	fileWriter := bufio.NewWriter(f)
	_, err = fileWriter.Write([]byte(page.HtmlContent))
	fileWriter.Flush()
	panicOnError(err)
}

func getSafeFileName(page Page) string {
	return strings.ReplaceAll(page.Postable.Title, " ", "-") + ".html"
}

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}
