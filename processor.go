package main

import "strings"

func processor(s string) string {
	s = strings.TrimSpace(s)
	word := strings.Split(s, "\n")
	for i := 0; i < len(word); i++ {
		word[i] = applycase(word[i])
	}
	return strings.Join(word, "\n")
}
