package stack

import (
	"data_struct/list/dlist"
	"data_struct/utils"
)

var (
	defaultContainer = dlist.New()
)

type Stack struct {
	container utils.Container
}

func New() *Stack {
	return &Stack{container: defaultContainer}
}

func (s *Stack) Push(value interface{}) {
	s.container.PushBack(value)
}

func (s *Stack) Pop() interface{} {
	return s.container.PopBack()
}

func (s *Stack) Top() interface{} {
	return s.container.Back()
}

func (s *Stack) Len() uint {
	return s.container.Len()
}

func (s *Stack) Empty() bool {
	return s.container.Empty()
}

func (s *Stack) Back() interface{} {
	return s.container.Back()
}