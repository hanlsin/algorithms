package data_structures

/*
https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
 */

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Top() interface{}
	Get(int) interface{}
	Len() int
}

type YPStack struct {
	size int
	list []Node
}

func NewYPStack() *YPStack {
	list := make([]Node, 2)

	return &YPStack{
		size: 0,
		list: list,
	}
}

func (s *YPStack) Push(obj interface{}) {
	l := len(s.list)
	if s.size == l {
		nl := make([]Node, 2*l)
		copy(nl, s.list)
		s.list = nl
	}

	s.list[s.size].object = obj
	s.size++
}

func (s *YPStack) Pop() interface{} {
	v := s.Top()
	s.size--
	return v
}

func (s *YPStack) Top() interface{} {
	return s.list[s.size-1].Value()
}

func (s *YPStack) Get(idx int) interface{} {
	if idx >= s.size {
		return nil
	}
	return s.list[idx].Value()
}

func (s *YPStack) Len() int {
	return s.size
}
