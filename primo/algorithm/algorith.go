package algorithm

func IsPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrimoRange(n int) []bool {
	values := []bool{}
	for i := 1; i <= n; i++ {
		val := IsPrimo(i)
		values = append(values, val)
	}
	return values
}
