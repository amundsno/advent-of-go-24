package matrix

type Matrix[T any] struct {
	matrix [][]T
	n, m   int
}

func New[T any](m [][]T) Matrix[T] {
	return Matrix[T]{m, len(m), len(m[0])}
}

func (m *Matrix[T]) IsInbounds(i, j int) bool {
	return i >= 0 && i < m.n && j >= 0 && j < m.m
}

func (m *Matrix[T]) At(i, j int) *T {
	return &m.matrix[i][j]
}

func (m *Matrix[T]) Get(i, j int) T {
	return m.matrix[i][j]
}

func (m *Matrix[T]) Rows() int {
	return m.n
}

func (m *Matrix[T]) Cols() int {
	return m.m
}
