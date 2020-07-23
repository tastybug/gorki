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
	"strings"
)

type Postable struct {
	fileName      string
	Title         string
	Description   string
	ContentAsMd   string
	ContentAsHtml string
	assets        []Asset
}

func CollectArticlePages(workDir, templatesFolder string) map[string]WritableContent {
	postsDir := filepath.Join(workDir, "posts")
	var resultMap = make(map[string]WritableContent)
	for _, postable := range collectPostables(postsDir) {
		resultMap[postable.Title] = toWritableContent(postable, templatesFolder)
		collectAssetsForArticle(postsDir, postable)
	}
	return resultMap
}

func collectPostables(postsDir string) map[string]Postable {
	var postableMap = make(map[string]Postable)

	for _, mdFile := range util.ListFilesWithSuffix(postsDir, ".md") {
		postableMap[mdFile.Name()] = AssemblePostable(mdFile.Name(), util.ReadFileContent(postsDir, mdFile.Name()))
	}
	return postableMap
}

func toWritableContent(postable Postable, templatesFolder string) WritableContent {
	var b bytes.Buffer

	contentTemplate := createContentTemplate(postable.ContentAsHtml)
	t, _ := template.ParseFiles(
		filepath.Join(templatesFolder, `blogpost`, `blogpost.html`),
		filepath.Join(contentTemplate.Name()),
		filepath.Join(templatesFolder, "footer.html"),
		filepath.Join(templatesFolder, "navigation.html"),
		filepath.Join(templatesFolder, "head.html"))
	err := t.Execute(&b, postable)
	util.PanicOnError(err)

	defer os.Remove(contentTemplate.Name())

	return WritableContent{HtmlContent: b.String(), Path: getSafeFileName(postable)}
}

func createContentTemplate(content string) *os.File {
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func getSafeFileName(postable Postable) string {
	return strings.ReplaceAll(postable.Title, " ", "-") + ".html"
}

func collectAssetsForArticle(postsDir string, postable Postable) {
	assetFolderName := extractGroup(postable.fileName, `(?P<name>.*).md`, `name`)
	assetFolderPath := filepath.Join(postsDir, assetFolderName)
	if util.Exists(assetFolderPath) {
		log.Println("Found assets for " + postable.Title)
	}

}

func AssemblePostable(fileName, rawPostableContent string) (post Postable) {
	const structurePattern = `-{3}(?P<meta>[\s\w.:]+)-{3}(?P<content>[\s\w:.#]+)`
	const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
	const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`

	metadata := extractGroup(rawPostableContent, structurePattern, `meta`)
	mdContent := extractGroup(rawPostableContent, structurePattern, `content`)
	return Postable{
		fileName:      fileName,
		Title:         extractGroup(metadata, titlePattern, `value`),
		Description:   extractGroup(metadata, descriptionPattern, `value`),
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