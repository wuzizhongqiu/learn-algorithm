package code

// 2023_12_1 找出叠涂元素（读题/数组/哈希）
func firstCompleteIndex(arr []int, mat [][]int) int {
	mp := map[int][2]int{}
	n, m := len(mat), len(mat[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// mp 存的键值对是 arr 数组的值 : 在 mat 中对应的下标
			mp[mat[i][j]] = [2]int{i, j}
		}
	}
	rowCnt, colCnt := make([]int, n), make([]int, m)
	for i, v := range arr {
		t := mp[v]
		rowCnt[t[0]]++
		colCnt[t[1]]++
		// 只要有一行或者一列填满了, 就直接返回 arr 下标
		if rowCnt[t[0]] == m || colCnt[t[1]] == n {
			return i
		}
	}
	return -1
}

// 2023_12_2 拼车（模拟/差分）
func carPooling(trips [][]int, capacity int) bool {
	var numPeople [1001]int
	for _, v := range trips {
		n, a, b := v[0], v[1], v[2]
		numPeople[a] += n
		numPeople[b] -= n
	}
	curCap := 0
	for _, v := range numPeople {
		curCap += v
		if curCap > capacity {
			return false
		}
	}
	return true
}

// 【LeetCode】每日一题 2023_12_3 可获得的最大点数（前缀和/滑动窗口/贪心）
func maxScore(cardPoints []int, k int) int {
	front := 0
	for i := 0; i < k; i++ {
		front += cardPoints[i]
	}
	ans := front
	for i := 1; i <= k; i++ {
		front += cardPoints[len(cardPoints)-i] - cardPoints[k-i]
		ans = max(ans, front)
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 【LeetCode】每日一题 2023_12_4 从二叉搜索树到更大和树（二叉树）
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var convert func(*TreeNode)
	convert = func(node *TreeNode) {
		if node == nil {
			return
		}
		convert(node.Right) // 右
		sum += node.Val     // 累加
		node.Val = sum      // 根
		convert(node.Left)  // 左
	}
	convert(root)
	return root
}
