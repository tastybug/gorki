package main

import (
	"bloggo/postable"
	"testing"
)

func TestExistingTitleIsExtracted(t *testing.T) {
	var fileContent = `---
Title:The Bible
Description: So God said ..
---
`
	result := postable.CreatePostable(fileContent)
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
	result := postable.CreatePostable(fileContent)
	expectedTitle := ""
	if result.Title != expectedTitle {
		t.Errorf("Title, got: %s, want: %s.", result.Title, expectedTitle)
	}
}

func TestMissingMetadataLeadsToEmptyTitle(t *testing.T) {
	var fileContent = `# Header
Lorem ipsum..
`
	result := postable.CreatePostable(fileContent)
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
	result := postable.CreatePostable(fileContent)
	exptected := `
# Header
Lorem ipsum..
`
	if result.Content != exptected {
		t.Errorf("Content, got: %s, want: %s.", result.Content, exptected)
	}
}
