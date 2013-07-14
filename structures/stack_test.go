package structures

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {
	s := Stack{}
	s.Push(1)
	if val, err := s.Pop(); val != 1 || err != nil {
		t.Error("This test did not pass")
	}
}
