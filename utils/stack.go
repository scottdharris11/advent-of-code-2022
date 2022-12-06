package utils

type Stack struct {
	count  int
	crates []interface{}
}

func (s *Stack) Push(c interface{}) {
	s.crates = append(s.crates, c)
	s.count++
}

func (s *Stack) Pop() interface{} {
	if s.count == 0 {
		return nil
	}
	i := s.count - 1
	c := s.crates[i]
	s.crates = s.crates[:i]
	s.count--
	return c
}

func (s *Stack) Peek() interface{} {
	if s.count == 0 {
		return nil
	}
	i := s.count - 1
	return s.crates[i]
}
