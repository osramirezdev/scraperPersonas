package utils

import "strings"

func ExtractValidStrings(input string) []string {
	substrings := strings.Split(input, ", ")

	for i, str := range substrings {
		substrings[i] = strings.TrimSpace(str)
	}

	return substrings
}
