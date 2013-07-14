package structures

import (
	"errors"
)

type Stack struct {
	Elements []int
}

type Stackable interface {
	Pop() (int, error)
	Push(int)
}

func (s *Stack) Pop() (int, error) {
	if len(s.Elements) > 0 {
		ret := s.Elements[len(s.Elements)-1]
		s.Elements = s.Elements[:len(s.Elements)-1]
		return ret, nil
	}

	return -1, errors.New("didnt work bro stacks empty")
}

func (s *Stack) Push(i int) {
	s.Elements = append(s.Elements, i)
}
