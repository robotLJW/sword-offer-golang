package _1_offer_zhan_de_ya_ru_dan_chu_xu_lie_lcof

import "container/list"

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	stack := NewStack()
	j := 0
	for i := 0; i < len(pushed); i++ {
		stack.Push(pushed[i])
		for ; !stack.Empty() && stack.Peak().(int) == popped[j]; j++ {
			stack.Pop()
		}
	}
	return stack.Empty()
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list: list}
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

func (s *Stack) Peak() interface{} {
	e := s.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (s *Stack) Len() interface{} {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}
