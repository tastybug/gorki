package pages

import "testing"

func TestIsDraft(t *testing.T) {

	if IsDraft(`---\nDraft: false\n---`) != false {
		t.Error("Did not see that draft was declared FALSE")
	}
	if IsDraft(`---\nDraft: true\n---`) != true {
		t.Error("Did not see that draft was declared TRUE")
	}
	if IsDraft(``) != true {
		t.Error("Unless specifically set otherwise, an article is supposed to be a draft")
	}
}
