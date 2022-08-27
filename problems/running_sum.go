package problems

// 1480. Running Sum of 1d Array 
func RunningSum(nums []int) []int {
	respond := []int{}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			respond = append(respond, nums[0])
			continue
		}
		respond = append(respond, respond[i-1] + nums[i])
	}
	return respond
}