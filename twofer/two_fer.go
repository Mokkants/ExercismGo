package twofer

// ShareWith returns a string
func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}
	return "One for " + s + ", one for me."
}

// I kinda wish Exercism exercises were ordered by difficulty, this seems kinda silly