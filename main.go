package main

import (
	"bloggo/post"
	"fmt"
)

const WORK_DIR string = "testdata"

func main() {
	var posts map[string]post.BlogPost = post.GetSitePosts(WORK_DIR)

	for key, elem := range posts {
		fmt.Printf("%+v (file name %s)\n", elem, key)
	}

}
