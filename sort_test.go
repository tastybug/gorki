package main

import (
	"testing"
)

func TestSortStringsAsc(t *testing.T) {

	input := []string{`a`, `c`, `b`, `f`}

	result := SortStrings(input, false)

	if !Equal(result, []string{`a`, `b`, `c`, `f`}) {
		t.Fatalf("result is not sorted asc: %s", result)
	}
}

func TestSortStringsWithRepeatingSimilarValues(t *testing.T) {
	input := []string{`a`, `a`, `b`, `a`, `b`, `a`}

	result := SortStrings(input, false)

	if !Equal(result, []string{`a`, `a`, `a`, `a`, `b`, `b`}) {
		t.Fatalf("result is not sorted asc: %s", result)
	}
}

func TestSortStringsDesc(t *testing.T) {

	input := []string{`a`, `c`, `b`, `f`}
	// true
	result := SortStrings(input, true)

	if !Equal(result, []string{`f`, `c`, `b`, `a`}) {
		t.Fatalf("result is not sorted desc: %s", result)
	}
}

func TestSortSingleString(t *testing.T) {

	input := []string{`a`}
	// true
	result := SortStrings(input, true)

	if !Equal(result, []string{`a`}) {
		t.Fatalf("result is not sorted desc: %s", result)
	}
}

func TestRemoveFromSlice(t *testing.T) {

	// remove last
	input := []string{`1`, `2`, `3`}
	result := RemoveIndex(input, 2)
	if !Equal(result, []string{`1`, `2`}) {
		t.Fatalf("result is not as expected: %s", result)
	}

	// remove middle
	input = []string{`1`, `2`, `3`}
	result = RemoveIndex(input, 1)
	if !Equal(result, []string{`1`, `3`}) {
		t.Fatalf("result is not as expected: %s", result)
	}

	// remove first
	input = []string{`1`, `2`, `3`}
	result = RemoveIndex(input, 0)
	if !Equal(result, []string{`2`, `3`}) {
		t.Fatalf("result is not as expected: %s", result)
	}

	// remove singular entry
	input = []string{`1`}
	result = RemoveIndex(input, 0)
	if !Equal(result, []string{}) {
		t.Fatalf("result is not as expected: %s", result)
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
