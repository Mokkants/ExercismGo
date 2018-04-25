//Package scrabble contains functions related to the popular board-game Scrabble
package scrabble

import "unicode"

func letterScore(c byte) int {

	value := 0

	switch unicode.ToUpper(rune(c)) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		value = 1
	case 'D', 'G':
		value = 2
	case 'B', 'C', 'M', 'P':
		value = 3
	case 'F', 'H', 'V', 'W', 'Y':
		value = 4
	case 'K':
		value = 5
	case 'J', 'X':
		value = 8
	case 'Q', 'Z':
		value = 10
	}

	return value

}

//Score returns the points a given word is worth
func Score(s string) int {
	score := 0

	for i := 0; i < len(s); i++ {
		score += letterScore(s[i])
	}

	return score
}
