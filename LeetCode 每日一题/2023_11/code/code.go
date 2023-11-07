package code

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
