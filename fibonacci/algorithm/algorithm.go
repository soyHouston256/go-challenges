package algorithm

func Fibonacci(n int) []int {
	serie := []int{}
	n0 := 0
	n1 := 1
	for i := 1; i <= n; i++ {
		serie = append(serie, n0)
		fibonacci := n0 + n1
		n0 = n1
		n1 = fibonacci

	}

	return serie
}
