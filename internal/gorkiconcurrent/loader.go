package gorkiconcurrent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gomarkdown/markdown"
)

const structurePattern = "-{3}(?P<meta>[*&?/'\\-\\s\\w.,;:\\(\\)\\[\\]!\\-\"]+)-{3}(?P<content>[$|^\\s\\w.;=&{}\\\\%:_\"'\\*.#,\\-!\\[\\]\\(\\)\\/<>?€+`~-]+)"
const titlePattern = `[t|T]itle: ?(?P<value>[\w.,; &?*"-]*)`
const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
const descriptionPattern = `[d|D]escription: ?(?P<value>[&!?\/'\(\)\[\]\w.,; *\"-]*)`
const isDraftPattern = `[d|D]raft: ?(?P<value>(?:true|false)*)`

func load(articlesRootPage, templatesRoot string, inChan <-chan discovery, outChan chan<- bundle) {

	for in := range inChan {
		if in.kind == ARTICLE_BUNDLE {
			if bundle, err := newArticleBundle(articlesRootPage, in.fi.Name()); err != nil {
				log.Printf("skipping bundle: %s", err.Error())
				continue
			} else {
				log.Printf("Discovered article: %s\n", bundle.ArticleData.BucketName)
				outChan <- bundle
			}
		} else {
			bundle := newStaticBundle(templatesRoot, in.fi.Name())
			log.Printf("Discovered static page: %s\n", in.fi.Name())
			outChan <- bundle
		}
	}
	close(outChan)
}

func newStaticBundle(templatesRoot, name string) bundle {
	if name == `rss` { // ugly special treatment for rss bundle
		return bundle{
			name: name,
			kind: STATIC_BUNDLE,
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesRoot, name),
				name,
				`feed.tpl`,
				``,
				fmt.Sprintf(`%s.xml`, name)},
		}
	} else {
		return bundle{
			name: name,
			kind: STATIC_BUNDLE,
			TemplatingConf: templatingConf{
				``,
				filepath.Join(templatesRoot, name),
				name,
				fmt.Sprintf(`%s.html`, name),
				``,
				fmt.Sprintf(`%s.html`, name)},
		}
	}
}

func newArticleBundle(articlesRootPath, bundleName string) (bundle, error) {

	articlePath := filepath.Join(articlesRootPath, bundleName, `article.md`)
	if !FileExists(articlePath) {
		return bundle{}, fmt.Errorf("cannot build bundle '%s', no 'article.md' found", bundleName)
	}

	rawContent := read(articlePath)
	metadata := readMetadataBlock(rawContent)

	result := bundle{
		name: bundleName,
		kind: ARTICLE_BUNDLE,
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
			`_blogpost`,
			`blogpost.html`,
			bundleName,
			`article.html`},
	}

	return result, nil
}

func readPublishedDate(input string) string {
	return extractOrFail(input, publishedDatePattern, `value`)
}

func readDescription(input string) string {
	return extractOrFail(input, descriptionPattern, `value`)
}

func readTitle(input string) string {
	return extractOrFail(input, titlePattern, `value`)
}

func readContentBlock(input string) string {
	return extractOrFail(input, structurePattern, `content`)
}

func readMetadataBlock(input string) string {
	return extractOrFail(input, structurePattern, `meta`)
}

func isDraft(input string) bool {
	value := extractOrFail(input, isDraftPattern, `value`)
	return !(value == `false`)
}

func extractOrFail(data string, pattern string, groupName string) string {

	r := regexp.MustCompile(pattern)
	result := r.FindStringSubmatch(data)

	for index, value := range r.SubexpNames() {
		if value == groupName && len(result) >= index {
			return result[index]
		}
	}
	panic(fmt.Errorf("no match for group '%s' in pattern '%s': %s", groupName, pattern, data))
}

func (b *bundle) isToBeRendered() bool {
	return !b.ArticleData.IsDraft || b.kind == STATIC_BUNDLE
}

type bundle struct {
	name           string
	kind           int
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

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func closeFile(f os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}
}
