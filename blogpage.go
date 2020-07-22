package main

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

type Page struct {
	HtmlContent string
	Postable    Postable
}

func PublishPost(postable Postable, templatesFolder string) Page {

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

	return Page{HtmlContent: b.String(), Postable: postable}
}

func createContentTemplate(content string) *os.File {
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}
