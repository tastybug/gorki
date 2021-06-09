package pages

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDraft(t *testing.T) {

	assert.False(t, isDraft(`---\nDraft: false\n---`), "Did not see that draft was declared FALSE")
	assert.True(t, isDraft(`---\nDraft: true\n---`), "Did not see that draft was declared TRUE")
	assert.True(t, isDraft(``), "Unless specifically set otherwise, an article is supposed to be a draft")
}

func TestReadTitle(t *testing.T) {
	expected := `Some title &?.,;*"`
	input := fmt.Sprintf("---\nTitle: %s\n---\nContent", expected)

	actual := readTitle(input)

	assert.Equal(t, expected, actual)
}

func TestReadPublishedDate(t *testing.T) {
	expected := `2021-05-20`
	input := fmt.Sprintf("---\nPublishedDate: %s\n---\nContent", expected)

	actual := readPublishedDate(input)

	assert.Equal(t, expected, actual)
}

func TestReadDescription(t *testing.T) {
	expected := `You can't and won't and ? and & and * and !""`
	input := fmt.Sprintf("---\nDescription: %s\n---\nContent", expected)

	actual := readDescription(input)

	assert.Equal(t, expected, actual)
}
