package templating

import (
	"bloggo/postable"
	"bloggo/util"
	"bufio"
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Page struct {
	HtmlContent string
	Postable    postable.Postable
}

func CreateBlogPostPage(postable postable.Postable, templatesFolder string) Page {

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
	tmpFile, err := ioutil.TempFile(os.TempDir(), "tmp-template-")
	util.PanicOnError(err)

	fileWriter := bufio.NewWriter(tmpFile)
	_, err = fileWriter.Write([]byte("{{define \"content\"}}" + content + "{{end}}"))
	util.PanicOnError(err)
	err = fileWriter.Flush()
	util.PanicOnError(err)

	return tmpFile
}
