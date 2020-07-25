package pages

import (
	"bloggo/util"
	"github.com/gomarkdown/markdown"
	"log"
	"path/filepath"
)

const structurePattern = `-{3}(?P<meta>[\-\s\w.:\(\)\[\]!]+)-{3}(?P<content>[\s\w:.#,\-!\[\]\(\)\/]+)`
const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`
const isDraftPattern = `[d|D]raft: ?(?P<value>(?:true|false)*)`

func CollectArticles() []ContentPage {
	articlesRootPath := GetArticlesRootDirectory()
	var articles []ContentPage
	for _, bucket := range util.ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bucket.Name(), `article.md`)
		if util.Exists(articlePath) {
			article := assembleArticle(articlesRootPath, bucket.Name(), util.ReadFileContent(articlePath))
			if !article.isDraft {
				articles = append(articles, article)
			}
		} else {
			log.Printf("Skipping article %s, no 'article.md' found in it.", bucket.Name())
		}
	}
	return articles
}

func assembleArticle(articlesRootPath, bucketName, rawContent string) ContentPage {

	metadata := util.ExtractGroup(rawContent, structurePattern, `meta`)
	mdContent := util.ExtractGroup(rawContent, structurePattern, `content`)
	title := util.ExtractGroup(metadata, titlePattern, `value`)
	description := util.ExtractGroup(metadata, descriptionPattern, `value`)
	publishedDate := util.ExtractGroup(metadata, publishedDatePattern, `value`)
	isDraft := IsDraft(bucketName, metadata)

	return ContentPage{
		isDraft:       isDraft,
		isArticle:     true,
		BucketName:    bucketName,
		Title:         title,
		Description:   description,
		PublishedDate: publishedDate,
		TemplatingConf: TemplatingConf{
			string(markdown.ToHTML([]byte(mdContent), nil, nil)),
			filepath.Join(articlesRootPath, bucketName),
			`blogpost`,
			`blogpost.html`,
			bucketName,
			`article.html`},
	}
}

func IsDraft(bucketName, metadata string) bool {
	value := util.ExtractGroup(metadata, isDraftPattern, `value`)
	if value == `false` {
		return false
	} else if value == `true` {
		return true
	} else {
		log.Printf("WARNING: Bucket %s has missing or broken draft flag, assuming it is a draft.\n", bucketName)
		return true
	}
}
