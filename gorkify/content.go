package gorkify

import (
	"github.com/gomarkdown/markdown"
	"gorki/util"
	"log"
	"path/filepath"
)

const structurePattern = "-{3}(?P<meta>[*&?/'\\-\\s\\w.,;:\\(\\)\\[\\]!\\-\"]+)-{3}(?P<content>[$|^\\s\\w.;=&{}\\\\%:_\"'\\*.#,\\-!\\[\\]\\(\\)\\/<>?`~-]+)"
const titlePattern = `[t|T]itle: ?(?P<value>[\w.,; &?*"-]*)`
const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
const descriptionPattern = `[d|D]escription: ?(?P<value>[&!?\/'\(\)\[\]\w.,; *\"-]*)`
const isDraftPattern = `[d|D]raft: ?(?P<value>(?:true|false)*)`

func collectPages(settings util.Settings) []bundle {
	pages := collectArticleBundles(settings)
	pages = append(pages, collectStaticBundles(settings)...)
	return pages
}

func collectArticleBundles(settings util.Settings) []bundle {
	articlesRootPath := settings.ArticlesRoot
	var bundles []bundle
	for _, bundle := range util.ListDirectories(articlesRootPath) {
		articlePath := filepath.Join(articlesRootPath, bundle.Name(), `article.md`)
		if util.Exists(articlePath) {
			page := assembleArticlePage(articlesRootPath, bundle.Name(), util.ReadFileContent(articlePath))
			if !page.ArticleData.IsDraft {
				log.Printf("Picking up article '%s' at '%s'", page.ArticleData.Title, articlePath)
				bundles = append(bundles, page)
			} else {
				log.Printf("Skipping draft article '%s.'", bundle.Name())
			}
		} else {
			log.Printf("Skipping bundle '%s', no 'article.md' found.", bundle.Name())
		}
	}
	return bundles
}

func assembleArticlePage(articlesRootPath, bundleName, rawContent string) bundle {

	metadata := readMetadataBlock(rawContent)
	mdContent := readContentBlock(rawContent)
	title := readTitle(metadata)
	description := readDescription(metadata)
	publishedDate := readPublishedDate(metadata)
	isDraft := isDraft(metadata)

	log.Printf("Found bundle '%s':\n title: '%s',\n description: '%s',\n published on: '%s',\n draft: '%t'",
		bundleName, title, description, publishedDate, isDraft)

	return bundle{
		ArticleData: articleData{
			IsDraft:       isDraft,
			BucketName:    bundleName,
			Title:         title,
			Description:   description,
			PublishedDate: publishedDate,
		},
		TemplatingConf: templatingConf{
			string(markdown.ToHTML([]byte(mdContent), nil, nil)),
			filepath.Join(articlesRootPath, bundleName),
			`blogpost`,
			`blogpost.html`,
			bundleName,
			`article.html`},
	}
}

func readPublishedDate(input string) string {
	return util.ExtractGroup(input, publishedDatePattern, `value`)
}

func readDescription(input string) string {
	return util.ExtractGroup(input, descriptionPattern, `value`)
}

func readTitle(input string) string {
	return util.ExtractGroup(input, titlePattern, `value`)
}

func readContentBlock(input string) string {
	return util.ExtractGroup(input, structurePattern, `content`)
}

func readMetadataBlock(input string) string {
	return util.ExtractGroup(input, structurePattern, `meta`)
}

func isDraft(input string) bool {
	value := util.ExtractGroup(input, isDraftPattern, `value`)
	if value == `false` {
		return false
	} else if value == `true` {
		return true
	} else {
		return true
	}
}

func collectStaticBundles(settings util.Settings) []bundle {
	templatesFolderPath := settings.TemplatesRoot
	return []bundle{
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

type bundle struct {
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
