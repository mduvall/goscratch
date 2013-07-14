package structures

import (
	"testing"
)

func Test_QueueHasEnqueueAndQueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)

	if val, _ := q.Dequeue(); val != 1 {
		t.Error("Dequeue did not return desired result")
	}
}
