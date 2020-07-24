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
	BucketName    string
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
	for _, bucketName := range util.ListDirectories(postsDir) {
		articlePath := filepath.Join(postsDir, bucketName.Name(), `article.md`)
		if util.Exists(articlePath) {
			postables = append(postables, AssemblePostable(bucketName.Name(), util.ReadFileContent(articlePath)))
		} else {
			log.Printf("Skipping %s, no 'article.md' found in it.", bucketName)
		}
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
		Folders:     postable.BucketName,
		HtmlContent: htmlString.String(),
		FileName:    "article.html",
		assets:      collectAssetsForArticle(postsDir, postable),
	}
}

func createContentTemplate(content string) *os.File {
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func collectAssetsForArticle(postsDir string, postable Postable) []Asset {
	assetFolderPath := filepath.Join(postsDir, postable.BucketName)
	var assets []Asset
	if util.Exists(assetFolderPath) {
		for _, assetFile := range util.ListFilesWithoutSuffix(assetFolderPath, `.md`) {
			log.Println("Article %s has asset %s "+postable.BucketName, assetFile)
			assets = append(assets,
				Asset{
					Context:      postable.BucketName,
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

func AssemblePostable(bucketName, rawPostableContent string) (post Postable) {
	const structurePattern = `-{3}(?P<meta>[\-\s\w.:]+)-{3}(?P<content>[\s\w:.#,\-!\[\]\(\)\/]+)`
	const titlePattern = `[t|T]itle: ?(?P<value>[\w. ]*)`
	const publishedDatePattern = `[p|P]ublishedDate: ?(?P<value>[\-\:\w. ]*)`
	const descriptionPattern = `[d|D]escription: ?(?P<value>[\w. ]*)`

	metadata := extractGroup(rawPostableContent, structurePattern, `meta`)
	mdContent := extractGroup(rawPostableContent, structurePattern, `content`)
	return Postable{
		BucketName:    bucketName,
		Title:         extractGroup(metadata, titlePattern, `value`),
		Description:   extractGroup(metadata, descriptionPattern, `value`),
		PublishedDate: extractGroup(metadata, publishedDatePattern, `value`),
		ContentAsMd:   mdContent,
		ContentAsHtml: string(markdown.ToHTML([]byte(mdContent), nil, nil)),
	}
}

// TODO move to util
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
