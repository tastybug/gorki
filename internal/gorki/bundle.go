package gorki

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

const structurePattern = "-{3}(?P<meta>[*&?/'\\-\\s\\w.,;:\\(\\)\\[\\]!\\-\"]+)-{3}(?P<content>[$|^\\s\\w.;=&{}\\\\%:_\"'\\*.#,\\-!\\[\\]\\(\\)\\/<>?€+`~-]+)"
const titlePattern = `[t|T]itle: ?(?P<value>[\w.,; &?*"-]*)`
const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
const descriptionPattern = `[d|D]escription: ?(?P<value>[&!?\/'\(\)\[\]\w.,; *\"-]*)`
const isDraftPattern = `[d|D]raft: ?(?P<value>(?:true|false)*)`

func newBundle(articlesRootPath, bundleName string) (bundle, error) {

	articlePath := filepath.Join(articlesRootPath, bundleName, `article.md`)
	if !FileExists(articlePath) {
		return bundle{}, fmt.Errorf("cannot build bundle '%s', no 'article.md' found", bundleName)
	}

	rawContent := read(articlePath)
	metadata := readMetadataBlock(rawContent)

	result := bundle{
		ArticleData: articleData{
			IsDraft:       isDraft(metadata),
			BucketName:    bundleName,
			Title:         readTitle(metadata),
			Description:   readDescription(metadata),
			PublishedDate: readPublishedDate(metadata),
		},
		TemplatingConf: templatingConf{
			string(markdown.ToHTML([]byte(readContentBlock(rawContent)), nil, nil)),
			filepath.Join(articlesRootPath, bundleName),
			`blogpost`,
			`blogpost.html`,
			bundleName,
			`article.html`},
	}

	result.printSummary()
	return result, nil
}

func readPublishedDate(input string) string {
	return ExtractGroupOrFailOnMismatch(input, publishedDatePattern, `value`)
}

func readDescription(input string) string {
	return ExtractGroupOrFailOnMismatch(input, descriptionPattern, `value`)
}

func readTitle(input string) string {
	return ExtractGroupOrFailOnMismatch(input, titlePattern, `value`)
}

func readContentBlock(input string) string {
	return ExtractGroupOrFailOnMismatch(input, structurePattern, `content`)
}

func readMetadataBlock(input string) string {
	return ExtractGroupOrFailOnMismatch(input, structurePattern, `meta`)
}

func isDraft(input string) bool {
	value := ExtractGroupOrFailOnMismatch(input, isDraftPattern, `value`)
	return !(value == `false`)
}

func (b *bundle) isToBeRendered() bool {
	return !b.ArticleData.IsDraft
}

func (b *bundle) printSummary() {
	log.Println("Bundle--------------", b.ArticleData.BucketName)
	log.Println("    - title:        ", b.ArticleData.Title)
	log.Println("    - description:  ", b.ArticleData.Description)
	log.Println("    - published on: ", b.ArticleData.PublishedDate)
	log.Println("    - draft:        ", b.ArticleData.IsDraft)
	log.Println("    - article size  ", len(b.TemplatingConf.extraContent), "bytes")
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

func read(path string) string {
	if file, err := os.Open(path); err != nil {
		panic(err)
	} else {
		defer closeFile(*file)

		scanner := bufio.NewScanner(file)
		var content string
		for scanner.Scan() {
			content += scanner.Text() + "\n"
		}

		return content
	}
}
