package main

import "strings"

func SortStrings(unsorted []string, desc bool) []string {
	var sorted []string
	var bestMatch string
	var bestMatchIndex int

	for len(unsorted) > 0 {
		for index, candidate := range unsorted {
			if bestMatch == `` {
				bestMatch = candidate
				bestMatchIndex = index
			} else {
				result := strings.Compare(candidate, bestMatch)
				if (desc && result == 1) || (!desc && result == -1) {
					bestMatch = candidate
					bestMatchIndex = index
				}
			}
		}
		sorted = append(sorted, bestMatch)
		unsorted = RemoveIndex(unsorted, bestMatchIndex)
		bestMatch = ``
		bestMatchIndex = -1
	}

	return sorted
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
