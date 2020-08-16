package util

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func WriteToTempFile(content string) *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "gorki-")
	PanicOnError(err)

	fileWriter := bufio.NewWriter(tmpFile)
	_, err = fileWriter.Write([]byte(content))
	PanicOnError(err)
	err = fileWriter.Flush()
	PanicOnError(err)

	return tmpFile
}

func ListFilesMatching(dir, pattern string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)

	onlyWithSuffix := func(file os.FileInfo) bool {
		return matches(file.Name(), pattern) && !file.IsDir()
	}
	return filter(allFiles, onlyWithSuffix)
}

func ListFilesAndDirs(dir string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)
	return allFiles
}

func ListDirectories(dir string) []os.FileInfo {
	allFiles, err := ioutil.ReadDir(dir)
	PanicOnError(err)

	isDir := func(file os.FileInfo) bool { return file.IsDir() }
	return filter(allFiles, isDir)
}

func CreateDirIfNotExisting(path string) {
	if !Exists(path) {
		PanicOnError(os.MkdirAll(path, 0740))
	}
}

func ReadFileContent(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseFile(*file)

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

func CopyFile(sourcePath, destinationPath string) {
	in, err := os.Open(sourcePath)
	PanicOnError(err)
	defer CloseFile(*in)

	out, err := os.Create(destinationPath)
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

func CloseFile(f os.File) {
	err := f.Close()
	PanicOnError(err)
}
