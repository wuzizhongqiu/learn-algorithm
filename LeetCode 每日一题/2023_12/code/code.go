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
