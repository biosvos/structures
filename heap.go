package structures

type Heap struct {
	elements []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func parentIndex(index int) int {
	return (index+1)/2 - 1
}

func childIndices(index int) (int, int) {
	pivot := (index + 1) * 2
	return pivot - 1, pivot
}

func lessThan(a, b int) bool {
	return a < b
}

func (h *Heap) Push(value int) {
	curIdx := len(h.elements)
	h.elements = append(h.elements, value)
	parentIdx := parentIndex(curIdx)
	for h.existIndex(parentIdx) && lessThan(h.elements[curIdx], h.elements[parentIdx]) {
		h.elements[curIdx], h.elements[parentIdx] = h.elements[parentIdx], h.elements[curIdx]
		curIdx = parentIdx
		parentIdx = parentIndex(curIdx)
	}
}

func (h *Heap) existIndex(index int) bool {
	return 0 <= index && index < len(h.elements)
}

func (h *Heap) Pop() int {
	lastIndex := len(h.elements) - 1
	h.elements[0], h.elements[lastIndex] = h.elements[lastIndex], h.elements[0]
	ret := h.elements[lastIndex]
	h.elements = h.elements[:lastIndex]

	cur := 0
	next := 0
	for {
		cur = next
		left, right := childIndices(cur)
		if !h.existIndex(left) {
			break
		}
		if !h.existIndex(right) {
			next = left
			if lessThan(h.elements[next], h.elements[cur]) {
				h.elements[cur], h.elements[next] = h.elements[next], h.elements[cur]
			}
			break
		}

		if lessThan(h.elements[left], h.elements[right]) {
			next = left
		} else {
			next = right
		}

		if lessThan(h.elements[next], h.elements[cur]) {
			h.elements[cur], h.elements[next] = h.elements[next], h.elements[cur]
		} else {
			break
		}
	}

	return ret
}

func (h *Heap) IsEmpty() bool {
	return len(h.elements) == 0
}
