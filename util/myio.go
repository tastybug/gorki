package util

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func WriteToTempFile(content string) *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "bloggo-")
	PanicOnError(err)

	fileWriter := bufio.NewWriter(tmpFile)
	_, err = fileWriter.Write([]byte(content))
	PanicOnError(err)
	err = fileWriter.Flush()
	PanicOnError(err)

	return tmpFile
}

func ListFilesWithSuffix(dir, suffix string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)

	onlyMarkdown := func(file os.FileInfo) bool { return strings.HasSuffix(file.Name(), suffix) }
	return filter(allFiles, onlyMarkdown)
}

func ReadFileContent(dir string, fileName string) string {
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

func CopyFile(src, destination string) {
	in, err := os.Open(src)
	PanicOnError(err)
	defer in.Close()

	out, err := os.Create(destination)
	PanicOnError(err)

	_, err = io.Copy(out, in)
	PanicOnError(err)
	defer PanicOnError(out.Close())
}
