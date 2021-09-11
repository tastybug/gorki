package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDraft(t *testing.T) {

	data := []struct {
		input    string
		expected bool
	}{
		{"---\nDraft: false\n---", false},
		{"---\nDraft: true\n---", true},
		{"---\nDraft: asdasd\n---", true},
	}

	for _, test := range data {
		if actual := isDraft(test.input); actual != test.expected {
			t.Errorf("input '%s' did not yield expected result %t", test.input, test.expected)
		}
	}
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
	expected := `To aid my Terraform learning sessions, I wanted the equivalent of a "Hello World" project. A basic canvas to experiment with. A canvas that is bare-bone, covers the whole lifecycle from setup to teardown and employs AWS to make it a bit more exciting that plain "Hello World".`
	input := fmt.Sprintf("---\nDescription: %s\n---\nContent", expected)

	actual := readDescription(input)

	assert.Equal(t, expected, actual)
}

func TestReadContent(t *testing.T) {
	expected := "Check out the project and have a look at the [README](https://github.com/tastybug/terraform-aws-playground/blob/master/README.md) for complete instructions. Make sure that your AWS credentials are available at `~/.aws/credentials`, which you get when you setup `awscli` as recommended. If you want to avoid this, simply provide the credentials as environment variables. +?€"
	input := fmt.Sprintf("---\nBla\n---%s", expected)

	actual := readContentBlock(input)

	assert.Equal(t, expected, actual)
}

func Test_iso8601ToRfc822Date(t *testing.T) {
	testdata := []struct {
		input    string
		expected string
	}{
		{input: `2020-01-01`, expected: `Wed, 01 Jan 2020 00:00:00 +0000`},
		{input: `2020-12-24`, expected: `Thu, 24 Dec 2020 00:00:00 +0000`},
	}

	for _, data := range testdata {
		if got := ISODateToRSSDateTime(data.input); got != data.expected {
			t.Fatalf("Expected: '%s', but got '%s'", data.expected, got)
		}
	}
}
