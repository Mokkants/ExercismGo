// Package bob contains the voice of your inner teenager
package bob

import "regexp"

// Hey returns Bob's answer to what you told him
var question = regexp.MustCompile(`\?$`)
var yelling = regexp.MustCompile(`[^a-z]+\!?$`) // BUG: triggers from "1,2,3"
var yellingQuestion = regexp.MustCompile(`[A-Z]+\?$`)
var nothingToSay = regexp.MustCompile(`^$`)

func Hey(remark string) string {
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
