package main

import (
	"bloggo/pages"
	"testing"
)

func TestExistingTitleIsExtracted(t *testing.T) {
	var fileContent = `---
Title:The Bible
Description: So God said ..
---
`
	result := pages.CreatePostableFromRawString(fileContent)
	expectedTitle := "The Bible"
	if result.Title != expectedTitle {
		t.Errorf("Title, got: %s, want: %s.", result.Title, expectedTitle)
	}
}

func TestMissingTitleLeadsToEmpty(t *testing.T) {
	var fileContent = `---
Description: So God said ..
---
`
	result := pages.CreatePostableFromRawString(fileContent)
	expectedTitle := ""
	if result.Title != expectedTitle {
		t.Errorf("Title, got: %s, want: %s.", result.Title, expectedTitle)
	}
}

func TestMissingMetadataLeadsToEmptyTitle(t *testing.T) {
	var fileContent = `# Header
Lorem ipsum..
`
	result := pages.CreatePostableFromRawString(fileContent)
	expectedTitle := ""
	if result.Title != expectedTitle {
		t.Errorf("Title, got: %s, want: %s.", result.Title, expectedTitle)
	}
}

func TestContentIsExtracted(t *testing.T) {
	var fileContent = `---
Title:The Bible
Description: So God said ..
---
# Header
Lorem ipsum..
`
	result := pages.CreatePostableFromRawString(fileContent)
	exptected := `
# Header
Lorem ipsum..
`
	if result.ContentAsMd != exptected {
		t.Errorf("ContentAsMd, got: %s, want: %s.", result.ContentAsMd, exptected)
	}
}
