package structures

import (
	"errors"
)

type Queue struct {
	Elements []int
}

type Queueable interface {
	Enqueue(int)
	Dequeue() (int, error)
}

func (q *Queue) Enqueue(i int) {
	q.Elements = append(q.Elements, i)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.Elements) == 0 {
		return -1, errors.New("The queue is empty my man")
	}

	res := q.Elements[0]
	q.Elements = q.Elements[1:]
	return res, nil
}
