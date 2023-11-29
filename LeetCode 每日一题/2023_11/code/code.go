package code

import (
	"sort"
	"strings"
)

// 11_1 参加会议的最多员工数
func maximumInvitations(favorite []int) int {
	n := len(favorite)
	deg := make([]int, n)
	for _, f := range favorite {
		deg[f]++ // 统计基环树每个节点的入度
	}

	rg := make([][]int, n) // 反图
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 { // 拓扑排序，剪掉图上所有树枝
		x := q[0]
		q = q[1:]
		y := favorite[x] // x 只有一条出边
		rg[y] = append(rg[y], x)
		if deg[y]--; deg[y] == 0 {
			q = append(q, y)
		}
	}

	// 通过反图 rg 寻找树枝上最深的链
	var rdfs func(int) int
	rdfs = func(x int) int {
		maxDepth := 1
		for _, son := range rg[x] {
			maxDepth = max(maxDepth, rdfs(son)+1)
		}
		return maxDepth
	}

	maxRingSize, sumChainSize := 0, 0
	for i, d := range deg {
		if d == 0 {
			continue
		}

		// 遍历基环上的点
		deg[i] = 0    // 将基环上的点的入度标记为 0，避免重复访问
		ringSize := 1 // 基环长度
		for x := favorite[i]; x != i; x = favorite[x] {
			deg[x] = 0 // 将基环上的点的入度标记为 0，避免重复访问
			ringSize++
		}

		if ringSize == 2 { // 基环长度为 2
			sumChainSize += rdfs(i) + rdfs(favorite[i]) // 累加两条最长链的长度
		} else {
			maxRingSize = max(maxRingSize, ringSize) // 取所有基环长度的最大值
		}
	}
	return max(maxRingSize, sumChainSize)
}

// 环和杆
func countPoints(rings string) (ans int) {
	num := [10][3]int{}
	ch := 0
	for i := 0; i < len(rings); i += 2 {
		if rings[i] == 'R' {
			ch = 0
		} else if rings[i] == 'G' {
			ch = 1
		} else if rings[i] == 'B' {
			ch = 2
		}
		num[rings[i+1]-'0'][ch]++
	}
	for _, v1 := range num {
		cnt := 0
		for _, v2 := range v1 {
			if v2 > 0 {
				cnt++
			}
		}
		if cnt == 3 { // 三个标志位都有数，证明集齐了所有颜色
			ans++
		}
	}
	return ans
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 填充每个节点的下一个右侧节点指针 II
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	Queue := []*Node{root}
	for Queue != nil {
		tmp := Queue
		Queue = nil
		for i, node := range tmp {
			if i > 0 { // 层序过程中，连接一下指针
				tmp[i-1].Next = node
			}
			if node.Left != nil {
				Queue = append(Queue, node.Left)
			}
			if node.Right != nil {
				Queue = append(Queue, node.Right)
			}
		}
	}
	return root
}

// 数组中两个数的最大异或值
func findMaximumXOR(nums []int) (ans int) {
	mask := 0
	for i := 30; i >= 0; i-- { // 从最高位开始判断
		mp := map[int]bool{}
		mask |= 1 << i           // 把第 i 位, 置为 1
		checkAns := ans | 1<<i   // 将 checkAns的 第 i 位, 置为 1
		for _, v := range nums { // 遍历 nums 数组
			v &= mask           // i 位之后全置为 0
			if mp[checkAns^v] { // 如果存在两个数异或等于 checkAns
				ans = checkAns // checkAns 成真，更新 ans
				break
			}
			mp[v] = true // 将 v 塞进 map
		}
	}
	return ans
}

// 重复的DNA序列
func findRepeatedDnaSequencesM(s string) (ans []string) {
	mp := map[string]int{}
	for i := 0; i <= len(s)-10; i++ {
		str := s[i : i+10]
		mp[str]++
	}
	for k, v := range mp {
		if v >= 2 {
			ans = append(ans, k)
		}
	}
	return ans
}

func findRepeatedDnaSequences(s string) (ans []string) {
	mp := map[string]int{}
	for i := 0; i <= len(s)-10; i++ {
		sub := s[i : i+10]
		mp[sub]++
		if mp[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return ans
}

// 最大单词长度乘积
func maxProduct(words []string) (ans int) {
	marks := [1000]int{}
	for i, v := range words {
		t := 0
		for j := 0; j < len(v); j++ { // 用 int 的低 26 位来代指字母 a-z 是否出现
			u := v[j] - 'a'
			t |= 1 << u
		}
		marks[i] = t
	}
	for i := 0; i < len(words); i++ {
		for j := 0; j < i; j++ {
			if (marks[i] & marks[j]) == 0 { // 每个字符串对应的两个 int 执行 & 操作
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

// 统计范围内的元音字符串数
func vowelStrings(words []string, left int, right int) (ans int) {
	mp := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	for i := left; i <= right; i++ {
		if mp[words[i][0]] == 1 && mp[words[i][len(words[i])-1]] == 1 {
			ans++
		}
	}
	return ans
}

// 最长平衡子字符串
func findTheLongestBalancedSubstring(s string) (ans int) {
	n := len(s)
	for i := 0; i < n; {
		a, b := 0, 0
		for i < n && s[i] == '0' {
			a++
			i++
		}
		for i < n && s[i] == '1' {
			b++
			i++
		}
		ans = max(ans, min(a, b)*2)
	}
	return ans
}

// 逃离火灾
type pair struct{ x, y int }

var dir = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// bfs 函数：取三个数, 到达安全屋/安全屋上边/安全屋左边
	bfs := func(q []pair) (int, int, int) {
		time := make([][]int, m)
		for i, _ := range time { // 给 time 数组赋值为 -1
			time[i] = make([]int, n)
			for j, _ := range time[i] {
				time[i][j] = -1 // 表示未访问该点
			}
		}
		for _, v := range q { // 设置起点
			time[v.x][v.y] = 0
		}
		for t := 1; len(q) > 0; t++ {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, v2 := range dir {
					// x, y 设置偏移量, 然后控制边界, grid == 0 表示能走(不是墙), time < 0 也就是 == -1 表示未访问该节点
					if x, y := v.x+v2.x, v.y+v2.y; x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 0 && time[x][y] < 0 {
						time[x][y] = t
						q = append(q, pair{x, y}) // 需要 bfs 的新坐标起点
					}
				}
			}
		}
		return time[m-1][n-1], time[m-1][n-2], time[m-2][n-1]
	}

	manToHouseTime, m1, m2 := bfs([]pair{{0, 0}})
	if manToHouseTime < 0 { // 人能否到安全屋
		return -1
	}

	firPos := []pair{}
	for i, row := range grid { // 收集火的位置
		for j, x := range row {
			if x == 1 {
				firPos = append(firPos, pair{i, j})
			}
		}
	}
	fireToHouseTime, f1, f2 := bfs(firPos)
	if fireToHouseTime < 0 { // 火能否到安全屋
		return 1_000_000_000
	}

	d := fireToHouseTime - manToHouseTime
	if d < 0 { // 火到安全屋的时间是否比人快
		return -1
	}

	// 如果人需要从上面或左边进入安全屋, 与火到这两个位置的时间进行比较
	if m1 != -1 && m1+d < f1 || m2 != -1 && m2+d < f2 {
		return d
	}
	return d - 1
}

// 咒语和药水的成功对数
func successfulPairs(spells []int, potions []int, success int64) (ans []int) {
	sort.Ints(potions)
	for _, v := range spells {
		left, right := 0, len(potions)-1
		for left < right { // 二分找出成功组合的最小下标
			mid := left + (right-left)/2
			sum := int64(v) * int64(potions[mid])
			if sum >= success {
				right = mid
			} else {
				left = mid + 1
			}
		}
		cmp := int64(v) * int64(potions[left])
		if cmp >= success {
			ans = append(ans, len(potions)-right)
		} else { // 没有一个组合成功
			ans = append(ans, 0)
		}
	}
	return ans
}

// K 个元素的最大和
func maximizeSum(nums []int, k int) int {
	maxV := 0
	for _, v := range nums {
		maxV = max(maxV, v)
	}
	return (maxV + (maxV + k - 1)) * k / 2
}

// 最长奇偶子数组
func longestAlternatingSubarray(nums []int, threshold int) (ans int) {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] > threshold || nums[i]%2 != 0 { // 寻找起始点
			i++
			continue
		}
		start := i
		i++ // 从第二个点开始找
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		ans = max(ans, i-start)
	}
	return ans
}

// 数位和相等数对的最大和
func maximumSum(nums []int) int {
	val, ans := make([]int, 100), -1
	for _, v := range nums {
		t, cur := v, 0
		for t > 0 {
			cur += t % 10
			t /= 10
		}
		if val[cur] != 0 {
			ans = max(ans, val[cur]+v)
		}
		val[cur] = max(val[cur], v)
	}
	return ans
}

// 最大子数组和
func maxSubArray(nums []int) int {
	ans := -100000
	dp := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		dp[i] = max(nums[i-1], nums[i-1]+dp[i-1])
		ans = max(ans, dp[i])
	}
	return ans
}

// 美化数组的最少删除数
func minDeletion(nums []int) (cnt int) {
	for i := 0; i < len(nums); i++ {
		for i%2 == 0 && i < len(nums)-1 && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			cnt++
		}
	}
	if len(nums)%2 != 0 {
		nums = nums[:len(nums)-1]
		cnt++
	}
	return cnt
}

// HTML 实体解析器
func entityParser(s string) (ans string) {
	return strings.NewReplacer(`&quot;`, `"`, `&apos;`, `'`, `&gt;`, `>`, `&lt;`, `<`, `&frasl;`, `/`, `&amp;`, `&`).Replace(s)
}

// 统计和小于目标的下标对数目（暴力/双指针）
func countPairs(nums []int, target int) (ans int) {
	sort.Ints(nums)
	for left, right := 0, len(nums)-1; left < right; left++ {
		for right >= 0 && nums[left]+nums[right] >= target {
			right--
		}
		if left < right {
			ans += right - left
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func PseudoPalindromicPaths(root *TreeNode) int {
	cnt := make([]int, 10)
	dfs(root, cnt)
	return ans
}

func dfs(root *TreeNode, cnt []int) {
	if root == nil {
		return
	}
	cnt[root.Val]++
	if root.Left == nil && root.Right == nil {
		if isFalsePalindromes(cnt) {
			ans++
		}
		return
	}
	dfs(root.Left, cnt)
	dfs(root.Right, cnt)
}

func isFalsePalindromes(cnt []int) bool {
	odd := 0
	for _, v := range cnt {
		if v%2 == 1 {
			odd++
		}
	}
	return odd <= 1
}

func pseudoPalindromicPaths(root *TreeNode) int {
	return dfs2(root, 0)
}

// 二叉树中的伪回文路径
func dfs2(root *TreeNode, mask int) int {
	if root == nil {
		return 0
	}
	mask ^= 1 << root.Val        // 修改 root.Val 出现次数的奇偶性
	if root.Left == root.Right { // root 是叶子节点
		if mask&(mask-1) == 0 {
			return 1
		}
		return 0
	}
	return dfs2(root.Left, mask) + dfs2(root.Right, mask)
}

// FrontMiddleBackQueue 设计前中后队列（数组/链表/双端队列）
type FrontMiddleBackQueue struct {
	queue []int
	size  int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		queue: make([]int, 1001),
		size:  0,
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	tmp := make([]int, 1001)
	tmp[0] = val
	for i := 1; i < this.size+1; i++ {
		tmp[i] = this.queue[i-1]
	}
	this.queue = tmp
	this.size++
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	tmp := make([]int, 1001)
	for i := 0; i < this.size/2; i++ {
		tmp[i] = this.queue[i]
	}
	tmp[this.size/2] = val
	for i := this.size/2 + 1; i < this.size+1; i++ {
		tmp[i] = this.queue[i-1]
	}
	this.queue = tmp
	this.size++
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	tmp := make([]int, 1001)
	for i := 0; i < this.size; i++ {
		tmp[i] = this.queue[i]
	}
	tmp[this.size] = val
	this.queue = tmp
	this.size++
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.size == 0 {
		return -1
	}
	ans := this.queue[0]
	this.queue = this.queue[1:]
	this.size--
	return ans
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.size == 0 {
		return -1
	}
	ans := this.queue[(this.size-1)/2]
	this.queue = append(this.queue[:(this.size-1)/2], this.queue[(this.size-1)/2+1:]...)
	this.size--
	return ans
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.size == 0 {
		return -1
	}
	ans := this.queue[this.size-1]
	this.queue = this.queue[:this.size-1]
	this.size--
	return ans
}

// SmallestInfiniteSet 无限集中的最小数字（哈希/堆）
type SmallestInfiniteSet struct {
	mp   map[int]bool
	less int
}

func Constructor2() SmallestInfiniteSet {
	tmp := map[int]bool{}
	for i := 1; i < 1001; i++ {
		tmp[i] = true
	}
	return SmallestInfiniteSet{
		mp:   tmp,
		less: 1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	this.mp[this.less] = false
	tmp := this.less
	for i := 1; i < 1001; i++ {
		if this.mp[i] == true {
			this.less = i
			break
		}
	}
	return tmp
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	this.mp[num] = true
	if num < this.less {
		this.less = num
	}
}
