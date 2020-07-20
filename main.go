package main

import (
	"bloggo/postable"
	"fmt"
)

const workDir string = "testdata"

func main() {
	for fileName, post := range postable.CollectPostables(workDir) {
		fmt.Printf("file %s -> %+v\n", fileName, post)
	}
}
