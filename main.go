package main

import (
	"bloggo/postable"
	"fmt"
)

const WORK_DIR string = "testdata"

func main() {
	var posts = postable.CollectPostables(WORK_DIR)

	for key, elem := range posts {
		fmt.Printf("%+v (file name %s)\n", elem, key)
	}

}
