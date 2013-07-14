package structures

type Ring struct {
	Buffer  [10]int
	HeadIdx int
	TailIdx int
}

func (r *Ring) Insert(i int) {
	if r.TailIdx == r.HeadIdx {
		r.HeadIdx++
	}

	r.Buffer[r.TailIdx] = i
	r.TailIdx++

	if r.TailIdx == 10 {
		r.TailIdx = 0
	}

	if r.HeadIdx == 10 {
		r.HeadIdx = 0
	}
}

func (r *Ring) Delete() {
	r.Buffer[r.HeadIdx] = 0
	r.HeadIdx++

	if r.TailIdx == r.HeadIdx {
		r.TailIdx++
	}

	if r.TailIdx == 10 {
		r.TailIdx = 0
	}

	if r.HeadIdx == 10 {
		r.HeadIdx = 0
	}
}
