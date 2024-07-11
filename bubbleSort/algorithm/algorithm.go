package algorithm

func BubbleSort(nums []int, order string) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if (order == "asc" && nums[j] > nums[j+1]) || (order == "desc" && nums[j] < nums[j+1]) {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	return nums
}
