package structures

var _ Matrix = &ArrayMatrix{}

type ArrayMatrix struct {
	elements [][]int
}

func NewArrayMatrix(height int, width int, defaultValue int) *ArrayMatrix {
	ret := ArrayMatrix{}
	for i := 0; i < height; i++ {
		var sub []int
		for j := 0; j < width; j++ {
			sub = append(sub, defaultValue)
		}
		ret.elements = append(ret.elements, sub)
	}
	return &ret
}

func (m *ArrayMatrix) Set(row int, col int, value int) {
	m.elements[row][col] = value
}

func (m *ArrayMatrix) Get(row int, col int) int {
	return m.elements[row][col]
}
