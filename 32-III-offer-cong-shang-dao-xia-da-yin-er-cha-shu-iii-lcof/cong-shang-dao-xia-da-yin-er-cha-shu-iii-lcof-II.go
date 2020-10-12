package _2_III_offer_cong_shang_dao_xia_da_yin_er_cha_shu_iii_lcof

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() Stack {
	l := list.New()
	return Stack{
		l,
	}
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() interface{} {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

func levelOrderTwo(root *TreeNode) [][]int {
	if root==nil{
		return [][]int{}
	}
	s1 := NewStack()
	s2 := NewStack()
	levelData := make([]int, 0)
	ans := make([][]int, 0)
	current := 0
	s1.Push(root)
	for !s1.Empty() || !s2.Empty() {
		if current == 0 {
			e := s1.Pop().(*TreeNode)
			if e != nil {
				levelData = append(levelData, e.Val)
				if e.Left != nil {
					s2.Push(e.Left)
				}
				if e.Right != nil {
					s2.Push(e.Right)
				}
			}
			if s1.Empty() {
				ans = append(ans, levelData)
				levelData = make([]int, 0)
				current = 1 - current
			}
		} else {
			e := s2.Pop().(*TreeNode)
			if e != nil {
				levelData = append(levelData, e.Val)
				if e.Right != nil {
					s1.Push(e.Right)
				}
				if e.Left != nil {
					s1.Push(e.Left)
				}
			}
			if s2.Empty() {
				ans = append(ans, levelData)
				levelData = make([]int, 0)
				current = 1 - current
			}
		}
	}
	return ans
}
