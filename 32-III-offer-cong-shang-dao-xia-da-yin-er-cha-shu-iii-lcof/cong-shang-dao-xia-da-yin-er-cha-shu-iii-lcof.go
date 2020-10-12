package _2_III_offer_cong_shang_dao_xia_da_yin_er_cha_shu_iii_lcof

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	list *list.List
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := NewQueue()
	queue.Push(root)
	currentCount, nextCount := 1, 0
	levelCount := 1
	for !queue.Empty() {
		levelData := make([]int, 0)
		for i := 0; i < currentCount; i++ {
			e := queue.Pop().(*TreeNode)
			levelData = append(levelData, e.Val)
			if e.Left != nil {
				queue.Push(e.Left)
				nextCount++
			}
			if e.Right != nil {
				queue.Push(e.Right)
				nextCount++
			}
		}
		currentCount = nextCount
		nextCount = 0
		if levelCount%2 == 0 {
			ans = append(ans, reverse(levelData))
		} else {
			ans = append(ans, levelData)
		}

		levelCount++
	}
	return ans
}

func reverse(data []int) []int {
	for i, j := 0, len(data)-1; i < j; i, j = i-1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func NewQueue() Queue {
	l := list.New()
	return Queue{
		l,
	}
}

func (q *Queue) Push(value interface{}) {
	q.list.PushBack(value)
}

func (q *Queue) Pop() interface{} {
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}
