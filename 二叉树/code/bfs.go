package code

// 二叉树的层序遍历：

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 二叉树的层序遍历
func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	curQueue := []*TreeNode{root}
	for len(curQueue) > 0 {
		nextQueue := []*TreeNode{}
		tmp := []int{}
		for _, node := range curQueue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		ans = append(ans, tmp)
		curQueue = nextQueue
	}
	return ans
}

// 二叉树的层序遍历 II
func levelOrderBottom(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	curQueue := []*TreeNode{root}
	for len(curQueue) > 0 {
		nextQueue := []*TreeNode{}
		tmp := []int{}
		for _, node := range curQueue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		ans = append(ans, tmp)
		curQueue = nextQueue
	}
	reverseI(ans)
	return ans
}

func reverseI(a [][]int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
