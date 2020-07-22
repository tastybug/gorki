package main

import (
	"bloggo/templating"
	"testing"
)

func TestBla(t *testing.T) {

	// given
	postable := Postable{
		Title:       `The Bible`,
		Description: `God said that..`,
		ContentAsMd: `# Header\nLorem Ipsum..`,
	}

	// when
	result := templating.PublishPost(postable, `testdata/templates`)

	// then
	expectedHtmlContent := `<html>
    <head>
        <title>The Bible</title>
    </head>
    <body>
    # Header\nLorem Ipsum..
    </body>
</html>`

	if result.HtmlContent != expectedHtmlContent {
		t.Errorf("Html content mismatch, got: %s, want: %s.", result.HtmlContent, expectedHtmlContent)
	}
}
