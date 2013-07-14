package structures

type Heap struct {
	Elements []int
}

func BuildHeap(vals []int) *Heap {
	h := Heap{Elements: []int{-10}}
	for _, num := range vals {
		h.Insert(num)
	}

	return &h
}

func (h *Heap) Insert(i int) {
	h.Elements = append(h.Elements, i)

	insertIdx := len(h.Elements) - 1
	for h.Elements[h.Parent(insertIdx)] == -1 || h.Elements[h.Parent(insertIdx)] > h.Elements[insertIdx] {
		h.SwapParent(insertIdx)
		insertIdx = h.Parent(insertIdx)
	}
}

func (h *Heap) DeleteMin() {
	h.Elements[1] = h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[:len(h.Elements)-1]
	h.BubbleDown(1)
}

func (h *Heap) BubbleDown(idx int) {
	var childIndex int

	if idx*2 > len(h.Elements)-1 {
		return
	}

	// First pick the smaller child
	if (idx*2)+1 > len(h.Elements)-1 || h.Elements[idx*2] < h.Elements[(idx*2)+1] {
		childIndex = (idx * 2)
	} else {
		childIndex = (idx * 2) + 1
	}

	// Swap the smaller child with the daddy if applicable
	if h.Elements[idx] > h.Elements[childIndex] {
		h.Swap(childIndex, idx)
		h.BubbleDown(childIndex)
	}
}

func (h *Heap) Parent(idx int) int {
	if idx == 1 {
		return 1
	}

	return int(idx / 2)
}

func (h *Heap) SwapParent(idx int) {
	h.Swap(idx, h.Parent(idx))
}

func (h *Heap) Swap(idx int, idxa int) {
	tmp := h.Elements[idx]
	h.Elements[idx] = h.Elements[idxa]
	h.Elements[idxa] = tmp
}

func (h *Heap) Min() int {
	return h.Elements[1]
}
