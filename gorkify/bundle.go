package gorkify

import (
	"gorki/util"
	"log"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

const structurePattern = "-{3}(?P<meta>[*&?/'\\-\\s\\w.,;:\\(\\)\\[\\]!\\-\"]+)-{3}(?P<content>[$|^\\s\\w.;=&{}\\\\%:_\"'\\*.#,\\-!\\[\\]\\(\\)\\/<>?`~-]+)"
const titlePattern = `[t|T]itle: ?(?P<value>[\w.,; &?*"-]*)`
const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
const descriptionPattern = `[d|D]escription: ?(?P<value>[&!?\/'\(\)\[\]\w.,; *\"-]*)`
const isDraftPattern = `[d|D]raft: ?(?P<value>(?:true|false)*)`

func newBundle(articlesRootPath, bundleName, rawContent string) bundle {

	metadata := readMetadataBlock(rawContent)
	mdContent := readContentBlock(rawContent)
	title := readTitle(metadata)
	description := readDescription(metadata)
	publishedDate := readPublishedDate(metadata)
	isDraft := isDraft(metadata)

	log.Printf("Found bundle '%s':\n title: '%s',\n description: '%s',\n published on: '%s',\n draft: %t",
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
	return util.ExtractGroupOrFailOnMismatch(input, publishedDatePattern, `value`)
}

func readDescription(input string) string {
	return util.ExtractGroupOrFailOnMismatch(input, descriptionPattern, `value`)
}

func readTitle(input string) string {
	return util.ExtractGroupOrFailOnMismatch(input, titlePattern, `value`)
}

func readContentBlock(input string) string {
	return util.ExtractGroupOrFailOnMismatch(input, structurePattern, `content`)
}

func readMetadataBlock(input string) string {
	return util.ExtractGroupOrFailOnMismatch(input, structurePattern, `meta`)
}

func isDraft(input string) bool {
	value := util.ExtractGroupOrFailOnMismatch(input, isDraftPattern, `value`)
	if value == `false` {
		return false
	} else if value == `true` {
		return true
	} else {
		return true
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
