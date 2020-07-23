package proc

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type WritableContent struct {
	HtmlContent string
	Path        string // can be a file name or subpath in target
}

func ToWritableContent(postable Postable, templatesFolder string) WritableContent {
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
