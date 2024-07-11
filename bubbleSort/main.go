package main

import (
	"fmt"

	"github.com/soyHouston256/go-challenges/bubbleSort/algorithm"
)

func main() {
	fmt.Println("bubble sort")
	nums := []int{1, 8, 3, 4, 5, 6, 7, 2}
	fmt.Println(algorithm.BubbleSort(nums, "asc"))
	fmt.Println(algorithm.BubbleSort(nums, "desc"))
}
