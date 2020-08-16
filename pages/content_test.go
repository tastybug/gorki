package pages

import "testing"

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
