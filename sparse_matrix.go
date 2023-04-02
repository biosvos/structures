package structures

var _ Matrix = &SparseMatrix{}

type SparseMatrix struct {
	elements     map[int]map[int]int
	defaultValue int
}

func (s *SparseMatrix) Set(row, col int, value int) {
	_, ok := s.elements[row]
	if !ok {
		s.elements[row] = map[int]int{}
	}
	s.elements[row][col] = value
}

func (s *SparseMatrix) Get(row, col int) int {
	_, ok := s.elements[row]
	if !ok {
		return s.defaultValue
	}
	ret, ok := s.elements[row][col]
	if !ok {
		return s.defaultValue
	}
	return ret
}

func NewSparseMatrix(defaultValue int) *SparseMatrix {
	return &SparseMatrix{
		elements:     map[int]map[int]int{},
		defaultValue: defaultValue,
	}
}
