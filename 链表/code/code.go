package code

type ListNode struct {
	 Val int
	 Next *ListNode
 }

// 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	phead := &ListNode{Next: head} // 创建哨兵位
	cur := phead
	for cur.Next != nil {
		tmp := cur
		cur = cur.Next // 遍历
		if cur.Val == val { // 移除节点的操作
			tmp.Next = cur.Next
			cur = tmp
		}
	}
	return phead.Next
}