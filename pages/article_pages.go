package pages

import (
	"bloggo/util"
	"github.com/gomarkdown/markdown"
	"log"
	"path/filepath"
)

func CollectArticles() []ContentPage {
	articlesRootPath := GetArticlesRootDirectory()
	var articles []ContentPage
	for _, bucket := range util.ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bucket.Name(), `article.md`)
		if util.Exists(articlePath) {
			articles = append(articles, assembleArticle(articlesRootPath, bucket.Name(), util.ReadFileContent(articlePath)))
		} else {
			log.Printf("Skipping article %s, no 'article.md' found in it.", bucket.Name())
		}
	}
	return articles
}

func assembleArticle(articlesRootPath, bucketName, rawContent string) ContentPage {
	const structurePattern = `-{3}(?P<meta>[\-\s\w.:]+)-{3}(?P<content>[\s\w:.#,\-!\[\]\(\)\/]+)`
	const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
	const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
	const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`

	metadata := util.ExtractGroup(rawContent, structurePattern, `meta`)
	mdContent := util.ExtractGroup(rawContent, structurePattern, `content`)
	title := util.ExtractGroup(metadata, titlePattern, `value`)
	description := util.ExtractGroup(metadata, descriptionPattern, `value`)
	publishedDate := util.ExtractGroup(metadata, publishedDatePattern, `value`)

	return ContentPage{
		isArticle:     true,
		BucketName:    bucketName,
		Title:         title,
		Description:   description,
		PublishedDate: publishedDate,
		templatingConf: TemplatingConf{
			string(markdown.ToHTML([]byte(mdContent), nil, nil)),
			filepath.Join(articlesRootPath, bucketName),
			`blogpost`,
			`blogpost.html`,
			bucketName,
			`article.html`},
	}
}
