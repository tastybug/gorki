package main

import (
	"os"
)

func filter(files []os.FileInfo, test func(os.FileInfo) bool) (ret []os.FileInfo) {
	for _, f := range files {
		if test(f) {
			ret = append(ret, f)
		}
	}
	return
}

func PathExists(path string) bool {
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
	if err := f.Close(); err != nil {
		panic(err)
	}
}
