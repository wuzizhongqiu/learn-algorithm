package code

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
			tmp := n%10
			sum += tmp*tmp
			n /= 10
		}
		return sum
	}
	fast, slow := n, n
	for {
		slow = Sum(slow)
		fast = Sum(Sum(fast))
		if fast == slow {
			break;
		}
	}
	return fast == 1
}