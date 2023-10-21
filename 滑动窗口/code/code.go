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

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	win := [128]bool{}
	left, len := 0, 0
	for right, v := range s {
		for win[v] == true { // 出现了重复的字符，开始循环去重（代码的核心）
			win[s[left]] = false
			left++
		}
		win[v] = true
		len = max(len, right-left+1)
	}
	return len
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// 最大连续1的个数 III
func longestOnes(nums []int, k int) int {
	left, cnt0, len := 0, 0, 0
	for right, v := range nums {
		if v == 0 {
			cnt0++
		}
		for cnt0 > k {
			if nums[left] == 0 {
				cnt0--
			}
			left++
		}
		len = max(len, right-left+1)
	}
	return len
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 将 x 减到 0 的最小操作数
func minOperations(nums []int, x int) int {
	left, right, sum, lenth, target := 0, 0, 0, math.MaxInt32, -x
	for _, v := range nums {
		target += v
	}
	if target < 0 { // 如果全加上都达不到要求就直接返回
		return -1
	}
	for right < len(nums) {
		sum += nums[right]
		right++
		for sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target {
			lenth = min(lenth, len(nums)-(right-left))
		}
	}
	if lenth == math.MaxInt32 {
		return -1
	}
	return lenth
}

//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}

// 水果成篮
func totalFruit(fruits []int) int {
	win := map[int]int{}
	lenth, left := 0, 0
	for right, v := range fruits {
		win[v]++
		for len(win) > 2 {
			win[fruits[left]]--
			if win[fruits[left]] == 0 {
				delete(win, fruits[left])
			}
			left++
		}
		lenth = max(lenth, right-left+1)
	}
	return lenth
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
