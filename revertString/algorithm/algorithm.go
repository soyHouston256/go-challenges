package algorithm

func RevertString(text string) string {
	val := ""
	for i := 0; i < len(text); i++ {
		val = string(text[i]) + val
	}
	return val
}

func RevertWithRange(text string) string {
	resutl := ""
	for _, v := range text {
		resutl = string(v) + resutl
	}
	return resutl
}
