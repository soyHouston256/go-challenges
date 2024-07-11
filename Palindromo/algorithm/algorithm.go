package algorithm

func IsPalindromo(word string) bool {
	reverseWord := ""
	for _, leter := range word {
		reverseWord = string(leter) + reverseWord
	}
	if word == reverseWord {
		return true
	} else {
		return false
	}
}
