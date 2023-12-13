package code

import "sort"

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

// 【LeetCode】每日一题 2023_12_5 到达首都的最少油耗（树，搜索）
func minimumFuelCost(roads [][]int, seats int) (ans int64) {
	g := make([][]int, len(roads)+1)
	for _, v := range roads { // g[x] 数组存的是与 x 相连的节点
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(cur, father int) int {
		size := 1
		for _, child := range g[cur] {
			// 只从父节点向子节点搜索
			if child != father {
				// v 从子节点变当前节点, cur 从当前节点变父节点, 统计子树大小
				size += dfs(child, cur)
			}
		}
		if cur > 0 { // cur 不是根节点了, 可以计算油耗了
			ans += int64((size-1)/seats + 1)
		}
		return size
	}
	dfs(0, -1)
	return ans
}

// 【LeetCode】每日一题 2023_12_7 重新规划路线（DFS/BFS）
func minReorder(n int, connections [][]int) (ans int) {
	g := make([][][]int, n)
	for _, v := range connections {
		x, y := v[0], v[1]
		// 第一个参数是正常存坐标, 第二个参数代表的是指向, 指向为 1 代表 x->y
		g[x] = append(g[x], []int{y, 1})
		g[y] = append(g[y], []int{x, 0})
	}
	var dfs func(int, int)
	dfs = func(cur, father int) {
		for _, v := range g[cur] {
			if v[0] != father { // 只向叶子节点 dfs
				if v[1] == 1 { // 如果是从 0 节点方向往外指, 就让 ans++
					ans++
				}
				dfs(v[0], cur)
			}
		}
	}
	dfs(0, -1)
	return ans
}

// 【LeetCode】每日一题 2023_12_7 出租车的最大盈利（动态规划）
func maxTaxiEarnings(n int, rides [][]int) int64 {
	type pair struct{ s, p int } // 一个存上车点, 一个存盈利
	group := make([][]pair, n+1)
	for _, r := range rides {
		start, end, tip := r[0], r[1], r[2]
		group[end] = append(group[end], pair{start, end - start + tip}) // 根据 end 存 pair
	}
	dp := make([]int64, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] // 填 dp 数组, 无论有没有更改, 都先把上一个最大盈利继承一下
		for _, v := range group[i] {
			dp[i] = max(dp[i], dp[v.s]+int64(v.p)) // 遍历所有 end 为 i 的情况的最大盈利
		}
	}
	return dp[n]
}

// 【LeetCode】每日一题 2023_12_9 下一个更大的数值平衡数（枚举/打表二分）
func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ { // 枚举
		cnt := [10]int{}
		for tmp := i; tmp > 0; tmp /= 10 {
			cnt[tmp%10]++
		}
		isBeautifulNumber := true
		for j := i; j > 0; j /= 10 { // 判断是不是最小数值平衡数
			if j%10 != cnt[j%10] {
				isBeautifulNumber = false
				break
			}
		}
		if isBeautifulNumber == true {
			return i
		}
	}
}

// 70. 爬楼梯 经典 dp
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 1631. 最小体力消耗路径
type unionFind struct {
	parent, size []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

type edge struct {
	v, w, diff int
}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	edges := []edge{}
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].diff < edges[j].diff })

	uf := newUnionFind(n * m)
	for _, e := range edges {
		uf.union(e.v, e.w)
		if uf.inSameSet(0, n*m-1) {
			return e.diff
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 【LeetCode】每日一题 2023_12_12 下一个更大元素 IV（堆，优先级队列/单调栈）
func secondGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	for i, _ := range nums { // 初始化成 -1
		ans[i] = -1
	}
	st1, st2 := []int{}, []int{}
	for i, v := range nums {
		for len(st2) > 0 && nums[st2[len(st2)-1]] < v { // 这次遍历的值比 st2 栈顶大
			ans[st2[len(st2)-1]] = v
			st2 = st2[:len(st2)-1]
		}
		j := len(st1) - 1
		for j >= 0 && nums[st1[j]] < v { // 找到 v 的第一大整数
			j--
		}
		st2 = append(st2, st1[j+1:]...) // 将 v 的第一大整数及其之前的区间直接从栈 1 移入栈 2
		st1 = append(st1[:j+1], i)      // 每次将当前位置的下标入栈 1
	}
	return ans
}

// 【LeetCode】每日一题 2023_12_13 字典序最小回文串（双指针，模拟）
func makeSmallestPalindrome(str string) string {
	left, right := 0, len(str)-1
	s := []byte(str)
	for left < right {
		if s[left] < s[right] {
			s[right] = s[left]
		} else {
			s[left] = s[right]
		}
		left++
		right--
	}
	return string(s)
}

