// Package bob contains the voice of your inner teenager
package bob

import (
	"regexp"
	"strings"
)

var question = regexp.MustCompile(`\?$`)
var yelling = regexp.MustCompile(`^[^a-z]*[A-Z]+[A-Z0-9\W\s]*\!?$`)
var yellingQuestion = regexp.MustCompile(`^[A-Z\s]+\?$`)
var nothingToSay = regexp.MustCompile(`^[\s]*$`)

// Hey returns Bob's answer to what you told him
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	answer := "Whatever."
		switch  {
		case yellingQuestion.MatchString(remark):
			answer = "Calm down, I know what I'm doing!"
		case question.MatchString(remark):
			answer = "Sure."
		case yelling.MatchString(remark):
			answer = "Whoa, chill out!"
		case nothingToSay.MatchString(remark):
			answer = "Fine. Be that way!"
		}

	return answer
}
