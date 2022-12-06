package utils

type Stack struct {
	count int
	items []interface{}
}

func (s *Stack) Push(c interface{}) {
	s.items = append(s.items, c)
	s.count++
}

func (s *Stack) Pop() interface{} {
	if s.count == 0 {
		return nil
	}
	i := s.count - 1
	c := s.items[i]
	s.items = s.items[:i]
	s.count--
	return c
}

func (s *Stack) Peek() interface{} {
	if s.count == 0 {
		return nil
	}
	i := s.count - 1
	return s.items[i]
}
