package pages

import (
	"bloggo/util"
	"github.com/gomarkdown/markdown"
	"log"
	"path/filepath"
)

type ContentPage struct {
	isArticle      bool
	BucketName     string
	Title          string
	Description    string
	PublishedDate  string
	templatingConf TemplatingConf
}

const articleDirName string = `posts`

func CollectArticles(siteDir string) []ContentPage {
	articlesRootPath := filepath.Join(siteDir, articleDirName)
	var articles []ContentPage
	for _, bucketName := range util.ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bucketName.Name(), `article.md`)
		if util.Exists(articlePath) {
			articles = append(articles, assembleArticle(articlesRootPath, bucketName.Name(), util.ReadFileContent(articlePath)))
		} else {
			log.Printf("Skipping %s, no 'article.md' found in it.", bucketName)
		}
	}
	return articles
}

func assembleArticle(articlesRootPath, bucketName, rawPostableContent string) ContentPage {
	const structurePattern = `-{3}(?P<meta>[\-\s\w.:]+)-{3}(?P<content>[\s\w:.#,\-!\[\]\(\)\/]+)`
	const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
	const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
	const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`

	metadata := util.ExtractGroup(rawPostableContent, structurePattern, `meta`)
	mdContent := util.ExtractGroup(rawPostableContent, structurePattern, `content`)
	return ContentPage{
		isArticle:     true,
		BucketName:    bucketName,
		Title:         util.ExtractGroup(metadata, titlePattern, `value`),
		Description:   util.ExtractGroup(metadata, descriptionPattern, `value`),
		PublishedDate: util.ExtractGroup(metadata, publishedDatePattern, `value`),
		templatingConf: TemplatingConf{
			string(markdown.ToHTML([]byte(mdContent), nil, nil)),
			filepath.Join(articlesRootPath, bucketName),
			`blogpost`,
			`blogpost.html`,
			bucketName,
			`article.html`},
	}
}
