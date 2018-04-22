// Package acronym is a Lengthy Name Abbreviating Utility Library (LNAUL)
package acronym

import (
	"strings"
)

// Abbreviate returns an acronym for the given parameter string
func Abbreviate(s string) string {
	abbr := ""
	s = strings.Replace(s, "-", " ", -1) // I have no idea what the -1 means at the end but it works hue hue
	for _, s := range strings.Split(s, " ") {
		abbr += strings.ToUpper(string(s[0]))
	}
	return abbr
}
