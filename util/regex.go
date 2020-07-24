package util

import "regexp"

func ExtractGroup(data string, pattern string, groupName string) string {

	r := regexp.MustCompile(pattern)
	result := r.FindStringSubmatch(data)

	for index, value := range r.SubexpNames() {
		if value == groupName && len(result) >= index {
			return result[index]
		}
	}
	return ``
}

func matches(data, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(data)
}
