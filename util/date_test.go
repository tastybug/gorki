package util

import "testing"

func Test_iso8601ToRfc822Date(t *testing.T) {
	result := Iso8601ToRfc822Date(`2020-01-01`)

	if result != `01 Jan 20 00:00 UT` {
		t.Fail()
	}
}
