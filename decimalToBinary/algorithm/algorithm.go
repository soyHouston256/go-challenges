package algorithm

import "fmt"

func DecimalToBinary(n int) string {

	binary := ""
	for n > 0 {
		remainder := n % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		n = n / 2
	}
	return binary

}
