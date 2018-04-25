package raindrops

import "strconv"

func Convert(num int) string {
	raindrops := ""

	if num%3 == 0 {
		raindrops += "Pling"
	}
	if num%5 == 0 {
		raindrops += "Plang"
	}
	if num%7 == 0 {
		raindrops += "Plong"
	}

	if raindrops == "" {
		raindrops = strconv.Itoa(num)
	}

	return raindrops
}
