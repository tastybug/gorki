package post

import (
	"log"
	"regexp"
)

func CreateBlogPost(rawContent string) BlogPost {

	metadata := extractPattern(rawContent, `-{3}(?P<value>[\s\w:.]+)-{3}`)

	var post BlogPost
	post.Title = extractPattern(metadata, `[t|T]itle: ?(?P<value>[\w. ]*)`)
	post.Description = extractPattern(metadata, `[d|D]escription: ?(?P<value>[\w. ]*)`)
	post.Content = extractPattern(rawContent, `-{3}(?P<meta>[\s\w.:]+)-{3}(?P<value>[\s\w:.#]+)`)

	return post
}

func extractPattern(content string, pattern string) string {

	r := regexp.MustCompile(pattern)
	result := r.FindStringSubmatch(content)

	for index, value := range r.SubexpNames() {
		if value == "value" && len(result) >= index {
			log.Printf("For '%s' I found '%s'", pattern, result[index])
			return result[index]
		}
	}
	return ""
}

type BlogPost struct {
	Title       string
	Description string
	Content     string
}
