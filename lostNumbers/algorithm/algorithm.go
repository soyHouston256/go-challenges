package algorithm

func LostNumber(nums []int) []int {
	mapNums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mapNums[nums[i]] = nums[i]
	}
	numLost := []int{}
	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		_, ok := mapNums[i]
		if !ok {
			numLost = append(numLost, i)
		}
	}
	return numLost
}
