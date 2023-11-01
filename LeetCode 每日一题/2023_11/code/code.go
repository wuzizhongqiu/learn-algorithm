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
