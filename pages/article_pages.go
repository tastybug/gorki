package pages

import (
	"bloggo/util"
	"github.com/gomarkdown/markdown"
	"log"
	"path/filepath"
)

type ArticlePage struct {
	BucketName    string
	Title         string
	Description   string
	PublishedDate string
	ContentAsHtml string
}

type Articles struct {
	Articles     []ArticlePage
	ArticleCount int
}

const articleDirName string = `posts`

func CreateOrderListOfArticles(siteDir string) Articles {
	articlesRootPath := filepath.Join(siteDir, articleDirName)

	var postables []ArticlePage
	for _, postable := range collectArticles(articlesRootPath) {
		postables = append(postables, postable)
	}
	return Articles{Articles: postables, ArticleCount: len(postables)}
}

func collectArticles(articlesRootPath string) []ArticlePage {
	var articles []ArticlePage
	for _, bucketName := range util.ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bucketName.Name(), `article.md`)
		if util.Exists(articlePath) {
			articles = append(articles, assembleArticle(bucketName.Name(), util.ReadFileContent(articlePath)))
		} else {
			log.Printf("Skipping %s, no 'article.md' found in it.", bucketName)
		}
	}
	return articles
}

func assembleArticle(bucketName, rawPostableContent string) ArticlePage {
	const structurePattern = `-{3}(?P<meta>[\-\s\w.:]+)-{3}(?P<content>[\s\w:.#,\-!\[\]\(\)\/]+)`
	const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
	const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
	const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`

	metadata := util.ExtractGroup(rawPostableContent, structurePattern, `meta`)
	mdContent := util.ExtractGroup(rawPostableContent, structurePattern, `content`)
	return ArticlePage{
		BucketName:    bucketName,
		Title:         util.ExtractGroup(metadata, titlePattern, `value`),
		Description:   util.ExtractGroup(metadata, descriptionPattern, `value`),
		PublishedDate: util.ExtractGroup(metadata, publishedDatePattern, `value`),
		ContentAsHtml: string(markdown.ToHTML([]byte(mdContent), nil, nil)),
	}
}
