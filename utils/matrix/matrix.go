package matrix

import (
	"advent-of-code/utils"
)

type Matrix[T any] struct {
	matrix [][]T
	n, m   int
}

func New[T any](m [][]T) Matrix[T] {
	return Matrix[T]{m, len(m), len(m[0])}
}

func (m *Matrix[T]) IsInbounds(i, j int) bool {
	return !utils.IsOutOfBounds2D(i, j, m.n-1, m.m-1)
}

func (m *Matrix[T]) At(i, j int) *T {
	return &m.matrix[i][j]
}
