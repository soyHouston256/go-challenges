package algorithm

import (
	"strconv"
)

func IsArmstrong(n int) bool {
	str := strconv.Itoa(n)
	armstrong := 0
	for i := 0; i < len(str); i++ {
		number, _ := strconv.Atoi(string(str[i]))
		value := Pow(number, len(str))
		armstrong = armstrong + value
		if n == armstrong {
			return true
		}
	}
	return false
}

func Pow(n int, pow int) int {
	result := 1
	for i := 1; i <= pow; i++ {
		result = result * n
	}
	return result
}
