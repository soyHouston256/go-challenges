package algorithm

import "time"

func BisiestoGenerator(n int) []int {
	count := 0
	t := time.Now()
	currentYear := t.Year()
	bisiestos := []int{}
	for count < n {
		currentYear++
		if (currentYear%4 == 0 && currentYear%100 != 0) || (currentYear%400 == 0) {
			bisiestos = append(bisiestos, currentYear)
			count++
		}
	}
	return bisiestos
}
