package pages

import "testing"
import "fmt"

func TestIsDraft(t *testing.T) {

	if isDraft(`---\nDraft: false\n---`) != false {
		t.Error("Did not see that draft was declared FALSE")
	}
	if isDraft(`---\nDraft: true\n---`) != true {
		t.Error("Did not see that draft was declared TRUE")
	}
	if isDraft(``) != true {
		t.Error("Unless specifically set otherwise, an article is supposed to be a draft")
	}
}

func TestReadTitle(t *testing.T) {
	expected := `Some title &?.,;*`
	input := fmt.Sprintf("---\nTitle: %s\n---\nContent", expected)

	actual := readTitle(input)

	if actual != expected {
		t.Error(fmt.Printf("Expecting '%s', got '%s'", expected, actual))
	}
}

func TestReadPublishedDate(t *testing.T) {
	expected := `2021-05-20`
	input := fmt.Sprintf("---\nPublishedDate: %s\n---\nContent", expected)

	actual := readPublishedDate(input)

	if actual != expected {
		t.Error(fmt.Printf("Expecting '%s', got '%s'", expected, actual))
	}
}

func TestReadDescription(t *testing.T) {
	expected := `You can't and won't and ? and & and * and !`
	input := fmt.Sprintf("---\nDescription: %s\n---\nContent", expected)

	actual := readDescription(input)

	if actual != expected {
		t.Error(fmt.Printf("Expecting '%s', got '%s'", expected, actual))
	}
}
