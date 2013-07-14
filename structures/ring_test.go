package structures

import (
	"testing"
)

func Test_RingInsertionWorks(t *testing.T) {
	r := Ring{}
	for i := 1; i < 11; i++ {
		r.Insert(i)
	}

	r.Insert(41)
	r.Delete()
	r.Delete()
	r.Delete()
	r.Delete()
	r.Delete()
	r.Insert(42)
	r.Insert(43)
	r.Delete()
}
