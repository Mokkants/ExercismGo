package reverse

func String(s string) (result string) {
	for _, char := range s {
		result = string(char) + result
	}
	return
}
