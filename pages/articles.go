package pages

import (
	"bloggo/util"
	"github.com/gomarkdown/markdown"
	"log"
	"path/filepath"
	"regexp"
)

type Article struct {
	BucketName    string
	Title         string // From metadata
	Description   string // From metadata
	PublishedDate string
	ContentAsHtml string
}

type Articles struct {
	Articles     []Article
	ArticleCount int
}

func CreateOrderListOfArticles(siteDir string) Articles {
	postsDir := filepath.Join(siteDir, "posts")

	var postables []Article
	for _, postable := range collectArticles(postsDir) {
		postables = append(postables, postable)
	}
	return Articles{Articles: postables, ArticleCount: len(postables)}
}

func collectArticles(postsDir string) []Article {
	var postables []Article
	for _, bucketName := range util.ListDirectories(postsDir) {
		articlePath := filepath.Join(postsDir, bucketName.Name(), `article.md`)
		if util.Exists(articlePath) {
			postables = append(postables, assembleArticle(bucketName.Name(), util.ReadFileContent(articlePath)))
		} else {
			log.Printf("Skipping %s, no 'article.md' found in it.", bucketName)
		}
	}
	return postables
}

func assembleArticle(bucketName, rawPostableContent string) Article {
	const structurePattern = `-{3}(?P<meta>[\-\s\w.:]+)-{3}(?P<content>[\s\w:.#,\-!\[\]\(\)\/]+)`
	const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
	const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
	const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`

	metadata := extractGroup(rawPostableContent, structurePattern, `meta`)
	mdContent := extractGroup(rawPostableContent, structurePattern, `content`)
	return Article{
		BucketName:    bucketName,
		Title:         extractGroup(metadata, titlePattern, `value`),
		Description:   extractGroup(metadata, descriptionPattern, `value`),
		PublishedDate: extractGroup(metadata, publishedDatePattern, `value`),
		ContentAsHtml: string(markdown.ToHTML([]byte(mdContent), nil, nil)),
	}
}

// TODO move to util
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
