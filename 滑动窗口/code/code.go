package code

import "math"

// 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	sum, len, n := 0, math.MaxInt32, len(nums)
	left, right := 0, 0
	for right < n {
		sum += nums[right]
		for sum >= target {
			len = min(len, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if len == math.MaxInt32 {
		return 0
	}
	return len
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
