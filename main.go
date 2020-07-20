package main

import (
	"bloggo/post"
	"fmt"
)

const WORK_DIR string = "testdata"

func main() {
	var posts = post.GetPostables(WORK_DIR)

	for key, elem := range posts {
		fmt.Printf("%+v (file name %s)\n", elem, key)
	}

}
