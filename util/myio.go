package util

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
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

// TODO clean this up
func ListFilesWithoutSuffix(dir, suffix string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)

	onlyWithSuffix := func(file os.FileInfo) bool { return !strings.HasSuffix(file.Name(), suffix) }
	return filter(allFiles, onlyWithSuffix)
}

func ListFilesWithSuffix(dir, suffix string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)

	onlyWithSuffix := func(file os.FileInfo) bool { return strings.HasSuffix(file.Name(), suffix) }
	return filter(allFiles, onlyWithSuffix)
}

func ListDirectories(dir string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)

	isDir := func(file os.FileInfo) bool { return file.IsDir() }
	return filter(allFiles, isDir)
}

func ReadFileContent(path string) string {
	file, err := os.Open(path)
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

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
