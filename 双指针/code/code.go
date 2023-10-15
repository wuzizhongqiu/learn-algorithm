package code

import "sort"

// 移动零
func moveZeroes(nums []int) {
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

// 复写零
func duplicateZeros(arr []int) {
	left, right := 0, 0
	// 找到一共经过了几个 0, 调整好位置
	for right < len(arr) {
		if arr[left] == 0 { // 注意这里用的是 left
			right++
		}
		left++
		right++
	}
	left--
	right--
	// 反向遍历
	for left >= 0 {
		if right < len(arr) {
			arr[right] = arr[left]
		}
		if arr[left] == 0 { // 复写 0 的操作
			right--
			arr[right] = 0
		}
		left--
		right--
	}
}

// 快乐数
func isHappy(n int) bool {
	Sum := func(n int) int { // 进行一次快乐数的计算
		sum := 0
		for n > 0 {
			tmp := n % 10
			sum += tmp * tmp
			n /= 10
		}
		return sum
	}
	fast, slow := n, n
	for {
		slow = Sum(slow)
		fast = Sum(Sum(fast))
		if fast == slow {
			break
		}
	}
	return fast == 1
}

// 盛最多水的容器
func maxArea(height []int) int {
	left, right, max := 0, len(height)-1, 0
	for left < right {
		tmp := Min(height[left], height[right]) * (right - left) // 计算当前容量
		max = Max(max, tmp)                                      // 迭代出最大容量
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 有效三角形的个数
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		left, right := 0, i-1
		for left < right {
			if (nums[left] + nums[right]) > nums[i] {
				ans += (right - left)
				right--
			} else {
				left++
			}
		}
	}
	return ans
}

// 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums) - 1
	for i, v := range nums[:n-1] {
		if i != 0 && v == nums[i-1] { // 跳过重复数字
			continue
		}
		if v+nums[i+1]+nums[i+2] > 0 { // 三数相加无论如何都会 > 0, 不需要再遍历
			break
		}
		if v+nums[n]+nums[n-1] < 0 { // 三数最大值 < 0, 让 i 继续遍历
			continue
		}
		// 双指针
		left, right := i+1, n
		for left < right {
			s := v + nums[left] + nums[right]
			if s > 0 {
				right--
			} else if s < 0 {
				left++
			} else {
				ans = append(ans, []int{v, nums[left], nums[right]})
				// 跳过重复数字
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return ans
}

// 四数之和
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums)-1
	for a := 0; a < n-2; a++ {
		value1 := nums[a]
		if a != 0 && value1 == nums[a-1] { // 跳过重复数字
			continue
		}
		if value1+nums[a+1]+nums[a+2]+nums[a+3] > target { // 四数之和 > target
			break
		}
		if value1+nums[n-2]+nums[n-1]+nums[n] < target { // 四数最大和 < target
			continue
		}
		for b := a+1; b < n-1; b++ {
			value2 := nums[b]
			if b != a+1 && value2 == nums[b-1] { // 跳过重复数字
				continue
			}
			if value1+value2+nums[b+1]+nums[b+2] > target { // 四数之和 > target
				break
			}
			if value1+value2+nums[n-1]+nums[n] < target { // 四数最大和 < target
				continue
			}
			left, right := b+1, n
			for left < right {
				sum := value1+value2+nums[left]+nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					ans = append(ans, []int{value1, value2, nums[left], nums[right]})
					// 跳过重复数字
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return ans
}
