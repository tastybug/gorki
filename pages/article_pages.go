package pages

import (
	"github.com/gomarkdown/markdown"
	"gorki/util"
	"log"
	"path/filepath"
)

const structurePattern = `-{3}(?P<meta>[\-\s\w.:\(\)\[\]!]+)-{3}(?P<content>[\s\w:.#,\-!\[\]\(\)\/]+)`
const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`
const isDraftPattern = `[d|D]raft: ?(?P<value>(?:true|false)*)`

func CollectArticles() []Page {
	articlesRootPath := GetArticlesRootDirectory()
	var articles []Page
	for _, bucket := range util.ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bucket.Name(), `article.md`)
		if util.Exists(articlePath) {
			page := assembleArticlePage(articlesRootPath, bucket.Name(), util.ReadFileContent(articlePath))
			if !page.ArticleData.isDraft {
				articles = append(articles, page)
			}
		} else {
			log.Printf("Skipping article %s, no 'article.md' found in it.", bucket.Name())
		}
	}
	return articles
}

func assembleArticlePage(articlesRootPath, bucketName, rawContent string) Page {

	metadata := util.ExtractGroup(rawContent, structurePattern, `meta`)
	mdContent := util.ExtractGroup(rawContent, structurePattern, `content`)
	title := util.ExtractGroup(metadata, titlePattern, `value`)
	description := util.ExtractGroup(metadata, descriptionPattern, `value`)
	publishedDate := util.ExtractGroup(metadata, publishedDatePattern, `value`)
	isDraft := IsDraft(metadata)

	return Page{
		ArticleData: ArticleData{
			isDraft:       isDraft,
			BucketName:    bucketName,
			Title:         title,
			Description:   description,
			PublishedDate: publishedDate,
		},
		TemplatingConf: TemplatingConf{
			string(markdown.ToHTML([]byte(mdContent), nil, nil)),
			filepath.Join(articlesRootPath, bucketName),
			`blogpost`,
			`blogpost.html`,
			bucketName,
			`article.html`},
	}
}

func IsDraft(metadata string) bool {
	value := util.ExtractGroup(metadata, isDraftPattern, `value`)
	if value == `false` {
		return false
	} else if value == `true` {
		return true
	} else {
		return true
	}
}
