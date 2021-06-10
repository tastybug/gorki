package util

import (
	"errors"
	"fmt"
	"regexp"
)

func ExtractGroupOrFailOnMismatch(data string, pattern string, groupName string) string {

	r := regexp.MustCompile(pattern)
	result := r.FindStringSubmatch(data)

	for index, value := range r.SubexpNames() {
		if value == groupName && len(result) >= index {
			return result[index]
		}
	}
	panic(errors.New(fmt.Sprintf("No match for group '%s' in pattern '%s':\n%s.", groupName, pattern, data)))
}

func matches(data, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(data)
}
