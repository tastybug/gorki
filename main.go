package main

import (
	"bloggo/content"
	"fmt"
)

const WORK_DIR string = "testdata"

func main() {
	var posts = content.CollectPostables(WORK_DIR)

	for key, elem := range posts {
		fmt.Printf("%+v (file name %s)\n", elem, key)
	}

}
