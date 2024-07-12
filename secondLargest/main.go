package main

import (
	"fmt"

	"github.com/soyHouston256/go-challenges/secondLargest/algorithm"
)

func main() {
	nums := []int{1, 8, 3, 4, 5, 6, 7, 2}
	fmt.Println(algorithm.SecondLargest(nums))
}
