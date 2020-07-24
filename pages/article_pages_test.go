package pages

import "testing"

func TestIsDraft(t *testing.T) {

	if IsDraft(`Draft: false`) != false {
		t.Error("Did not see that draft was declared FALSE")
	}
	if IsDraft(`Draft: true`) != true {
		t.Error("Did not see that draft was declared TRUE")
	}
	if IsDraft(``) != true {
		t.Error("Unless specifically set otherwise, an article is a draft")
	}
}
