package main

import (
	"bloggo/post"
	"fmt"
)

const WORK_DIR string = "."

func main() {
	const TITLE_REGEX string = "^.$"
	var posts []post.BlogPost = post.GetSitePosts(WORK_DIR)

	for index, elem := range posts {
		fmt.Printf("%d, %s\n", index, elem.Title)
	}

}
