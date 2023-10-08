package code

// 二分经典模板题目
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	left, right, begin, end := 0, len(nums)-1, -1, -1
	// 求左区间
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			begin = mid
			right--
		}
	}
	// 恢复 left 和 right
	left, right = 0, len(nums)-1
	// 求右区间
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			end = mid
			left++
		}
	}
	return []int{begin, end}
}

// 有效的完全平方数
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)/2
		sqrt := mid * mid
		if sqrt < num {
			left = mid + 1
		} else if sqrt > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

// 寻找峰值
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else if nums[mid] > nums[mid+1] {
			right = mid
		}
	}
	return right
}
