package collections

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Enqueue(t T) {
	q.data = append(q.data, t)
}

func (q *Queue[T]) Dequeue() T {
	t := q.data[0]
	q.data = q.data[1:]
	return t
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}
