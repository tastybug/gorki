package main

import (
	"bloggo/pages"
	"testing"
)

func TestBla(t *testing.T) {

	// given
	postable := pages.Postable{
		Title:       `The Bible`,
		Description: `God said that..`,
		ContentAsMd: `# Header\nLorem Ipsum..`,
	}

	// when
	result := pages.ToWritableContent(postable, `testdata/templates`)

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
