package main

import "testing"

func Test_iso8601ToRfc822Date(t *testing.T) {
	result := ISODateToRSSDateTime(`2020-01-01`)

	expected := `Wed, 01 Jan 2020 00:00:00 +0000`
	if result != expected {
		t.Fatalf("Expected: '%s', but got '%s'", expected, result)
	}
}
