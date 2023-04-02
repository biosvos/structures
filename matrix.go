package structures

type Matrix interface {
	Set(row, col int, value int)
	Get(row, col int) int
}
