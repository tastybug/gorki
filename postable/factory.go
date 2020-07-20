package postable

import (
	"regexp"
)

func CreatePostableFromFile(rawPostableContent string) Postable {
	const structurePattern = `-{3}(?P<meta>[\s\w.:]+)-{3}(?P<content>[\s\w:.#]+)`
	const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
	const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`

	var post Postable
	metadata := extractGroup(rawPostableContent, structurePattern, `meta`)
	post.Content = extractGroup(rawPostableContent, structurePattern, `content`)
	post.Title = extractGroup(metadata, titlePattern, `value`)
	post.Description = extractGroup(metadata, descriptionPattern, `value`)

	return post
}

func extractGroup(content string, pattern string, groupAlias string) string {

	r := regexp.MustCompile(pattern)
	result := r.FindStringSubmatch(content)

	for index, value := range r.SubexpNames() {
		if value == groupAlias && len(result) >= index {
			return result[index]
		}
	}
	return ``
}

type Postable struct {
	Title       string
	Description string
	Content     string
}
