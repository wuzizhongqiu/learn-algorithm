package code

// 移动零
func moveZeroes(nums []int)  {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			tmp := nums[fast]
			nums[fast] = nums[slow]
			nums[slow] = tmp
			slow++
		}
		fast++
	}
}