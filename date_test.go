package main

import "testing"

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
