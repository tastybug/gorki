package templating

import (
	"bloggo/postable"
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

type Page struct {
	HtmlContent string
	Postable    postable.Postable
}

func CreateBlogPostPage(postable postable.Postable, templatesFolder string) Page {

	var b bytes.Buffer

	t, _ := template.ParseFiles(filepath.Join(templatesFolder, `blogpost.html`))
	t.Execute(&b, postable)

	fmt.Printf("Result %s", b.String())
	return Page{HtmlContent: b.String(), Postable: postable}
}
