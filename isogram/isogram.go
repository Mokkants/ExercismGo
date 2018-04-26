package isogram

import "strings"

//IsIsogram returns true only if the input string has no repeating characters
func IsIsogram(s string) bool {

	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)
	letterMap := make(map[rune]int)

	for _, char := range s {

		_, ok := letterMap[char]
		if ok {
			return false
		}

		letterMap[char]++

	}

	return true

}
