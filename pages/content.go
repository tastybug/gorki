package pages

import (
	"github.com/gomarkdown/markdown"
	"gorki/util"
	"log"
	"path/filepath"
)

const structurePattern = "-{3}(?P<meta>[&?/'\\-\\s\\w.,;:\\(\\)\\[\\]!]+)-{3}(?P<content>[$|^\\s\\w.;=&{}\\\\%:_\"'\\*.#,\\-!\\[\\]\\(\\)\\/<>?`]+)"
const titlePattern = `[t|T]itle: ?(?P<value>[\w.,; &?]*)`
const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
const descriptionPattern = `[d|D]escription: ?(?P<value>[\/'\(\)\[\]\w.,; ]*)`
const isDraftPattern = `[d|D]raft: ?(?P<value>(?:true|false)*)`

func collectPages(settings util.Settings) []page {
	pages := collectArticles(settings)
	pages = append(pages, collectMains(settings)...)
	return pages
}

func collectArticles(settings util.Settings) []page {
	articlesRootPath := settings.ArticlesRoot
	var articles []page
	for _, bucket := range util.ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bucket.Name(), `article.md`)
		if util.Exists(articlePath) {
			page := assembleArticlePage(articlesRootPath, bucket.Name(), util.ReadFileContent(articlePath))
			if !page.ArticleData.IsDraft {
				log.Printf("Picking up article '%s' at '%s'", page.ArticleData.Title, articlePath)
				articles = append(articles, page)
			} else {
				log.Printf("Skipping draft article '%s.'", bucket.Name())
			}
		} else {
			log.Printf("Skipping article '%s', no 'article.md' found in it.", bucket.Name())
		}
	}
	return articles
}

func assembleArticlePage(articlesRootPath, bucketName, rawContent string) page {

	metadata := util.ExtractGroup(rawContent, structurePattern, `meta`)
	mdContent := util.ExtractGroup(rawContent, structurePattern, `content`)
	title := util.ExtractGroup(metadata, titlePattern, `value`)
	description := util.ExtractGroup(metadata, descriptionPattern, `value`)
	publishedDate := util.ExtractGroup(metadata, publishedDatePattern, `value`)
	isDraft := isDraft(metadata)

	log.Printf("Found article '%s':\n title: '%s',\n description: '%s',\n published on: '%s',\n draft: '%t'",
		bucketName, title, description, publishedDate, isDraft)

	return page{
		ArticleData: articleData{
			IsDraft:       isDraft,
			BucketName:    bucketName,
			Title:         title,
			Description:   description,
			PublishedDate: publishedDate,
		},
		TemplatingConf: templatingConf{
			string(markdown.ToHTML([]byte(mdContent), nil, nil)),
			filepath.Join(articlesRootPath, bucketName),
			`blogpost`,
			`blogpost.html`,
			bucketName,
			`article.html`},
	}
}

func isDraft(metadata string) bool {
	value := util.ExtractGroup(metadata, isDraftPattern, `value`)
	if value == `false` {
		return false
	} else if value == `true` {
		return true
	} else {
		return true
	}
}

func collectMains(settings util.Settings) []page {
	templatesFolderPath := settings.TemplatesRoot
	return []page{
		{
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesFolderPath, `index`),
				`index`,
				`index.html`,
				``,
				`index.html`},
		},
		{
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesFolderPath, `about`),
				`about`,
				`about.html`,
				`about`,
				`about.html`},
		},
		{
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesFolderPath, `privacy-imprint`),
				`privacy-imprint`,
				`privacy-imprint.html`,
				`privacy-imprint`,
				`privacy-imprint.html`},
		},
		{
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesFolderPath, `rss`),
				`rss`,
				`feed.tpl`,
				``,
				`rss.xml`},
		},
	}
}

type page struct {
	ArticleData    articleData    // used in template
	TemplatingConf templatingConf // used in template
}

type articleData struct {
	IsDraft       bool
	BucketName    string // used in template
	Title         string // used in template
	Description   string // used in template
	PublishedDate string // used in template
}

type templatingConf struct {
	extraContent     string
	assetFolderPath  string
	templateFolder   string
	templateFileName string
	resultFolderName string
	ResultFileName   string // used in template
}
