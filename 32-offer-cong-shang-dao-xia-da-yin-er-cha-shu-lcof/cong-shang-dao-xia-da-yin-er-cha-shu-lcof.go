package _2_offer_cong_shang_dao_xia_da_yin_er_cha_shu_lcof

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	list *list.List
}

func NewQueue() Queue {
	l := list.New()
	return Queue{
		l,
	}
}

func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}

func (q *Queue) Push(value interface{}) {
	q.list.PushBack(value)
}

func (q *Queue) Pop() interface{} {
	v := q.list.Front()
	if v != nil {
		q.list.Remove(v)
		return v.Value
	}
	return nil
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	queue := NewQueue()
	queue.Push(root)
	for !queue.Empty() {
		temp := queue.Pop().(*TreeNode)
		ans = append(ans, temp.Val)
		if temp.Left != nil {
			queue.Push(temp.Left)
		}
		if temp.Right != nil {
			queue.Push(temp.Right)
		}
	}
	return ans
}
