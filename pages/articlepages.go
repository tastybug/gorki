package pages

import (
	"bloggo/util"
	"bytes"
	"github.com/gomarkdown/markdown"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Postable struct {
	CanonicalName string // Some-Blogpost, safe for FS
	SrcFileName   string // Some-Blogpost.md
	Title         string // From metadata
	Description   string // From metadata
	PublishedDate string
	ContentAsMd   string
	ContentAsHtml string
}

type Articles struct {
	Postables    []Postable
	ArticleCount int
}

func CollectArticlePages(siteDir string) []WritableContent {
	templatesFolder := filepath.Join(siteDir, `templates`)
	postsDir := filepath.Join(siteDir, "posts")

	var writables []WritableContent
	for _, postable := range collectPostables(postsDir) {
		writables = append(writables, toWritableContent(postable, postsDir, templatesFolder))
	}
	return writables
}

func CreateOrderListOfPreviewItems(siteDir string) Articles {
	postsDir := filepath.Join(siteDir, "posts")

	var postables []Postable
	for _, postable := range collectPostables(postsDir) {
		postables = append(postables, postable)
	}
	return Articles{Postables: postables, ArticleCount: len(postables)}
}

func collectPostables(postsDir string) []Postable {
	var postables []Postable
	for _, mdFile := range util.ListFilesWithSuffix(postsDir, ".md") {
		postables = append(postables, AssemblePostable(mdFile.Name(), util.ReadFileContent(postsDir, mdFile.Name())))
	}
	return postables
}

func toWritableContent(postable Postable, postsDir string, templatesFolder string) WritableContent {
	var htmlString bytes.Buffer

	contentTemplate := createContentTemplate(postable.ContentAsHtml)
	t, _ := template.ParseFiles(
		filepath.Join(templatesFolder, `blogpost.html`),
		filepath.Join(contentTemplate.Name()),
		filepath.Join(templatesFolder, "footer.html"),
		filepath.Join(templatesFolder, "navigation.html"),
		filepath.Join(templatesFolder, "head.html"))
	err := t.Execute(&htmlString, postable)
	util.PanicOnError(err)

	defer os.Remove(contentTemplate.Name())

	return WritableContent{
		HtmlContent:   htmlString.String(),
		PathToWriteTo: "/" + postable.CanonicalName + ".html",
		assets:        collectAssetsForArticle(postsDir, postable),
	}
}

func createContentTemplate(content string) *os.File {
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func collectAssetsForArticle(postsDir string, postable Postable) []Asset {
	baseName := extractGroup(postable.SrcFileName, `(?P<name>.*).md`, `name`)
	assetFolderPath := filepath.Join(postsDir, baseName)
	var assets []Asset
	if util.Exists(assetFolderPath) {
		log.Println("Found assets for " + postable.Title)
		for _, assetFile := range util.ListFilesWithSuffix(assetFolderPath, ``) {
			assets = append(assets,
				Asset{
					Context:      baseName,
					Filename:     assetFile.Name(),
					CopyFromPath: filepath.Join(assetFolderPath, assetFile.Name())})
		}
	}
	if len(assets) > 0 {
		return assets
	} else {
		return nil
	}
}

func AssemblePostable(fileName, rawPostableContent string) (post Postable) {
	const structurePattern = `-{3}(?P<meta>[\-\s\w.:]+)-{3}(?P<content>[\s\w:.#,\-!\[\]\(\)\/]+)`
	const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
	const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
	const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`

	metadata := extractGroup(rawPostableContent, structurePattern, `meta`)
	mdContent := extractGroup(rawPostableContent, structurePattern, `content`)
	canonicalName := filepath.Join(extractGroup(fileName, `(?P<name>.*).md`, `name`))
	return Postable{
		CanonicalName: canonicalName,
		SrcFileName:   fileName,
		Title:         extractGroup(metadata, titlePattern, `value`),
		Description:   extractGroup(metadata, descriptionPattern, `value`),
		PublishedDate: extractGroup(metadata, publishedDatePattern, `value`),
		ContentAsMd:   mdContent,
		ContentAsHtml: string(markdown.ToHTML([]byte(mdContent), nil, nil)),
	}
}

func extractGroup(content string, pattern string, groupAlias string) string {

	r := regexp.MustCompile(pattern)
	result := r.FindStringSubmatch(content)

	for index, value := range r.SubexpNames() {
		if value == groupAlias && len(result) >= index {
			return result[index]
		}
	}
	return ``
}
