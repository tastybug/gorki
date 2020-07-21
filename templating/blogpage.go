package templating

import (
	"bloggo/postable"
	"bytes"
	"html/template"
	"path/filepath"
)

type Page struct {
	HtmlContent string
	Postable    postable.Postable
}

func CreateBlogPostPage(postable postable.Postable, templatesFolder string) Page {

	var b bytes.Buffer

	t, _ := template.ParseFiles(filepath.Join(templatesFolder, `blogpost`, `blogpost.html`))
	t.Execute(&b, postable)

	return Page{HtmlContent: b.String(), Postable: postable}
}
