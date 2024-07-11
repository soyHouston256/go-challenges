package algorithm

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	val := n * Factorial(n-1)

	return val
}
