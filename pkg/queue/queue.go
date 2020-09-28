package queue

// Queue represents a queue struct
type Queue struct {
	items []interface{}
}

// Enqueue appends item
func (q *Queue) Enqueue(i interface{}) {
	q.items = append(q.items, i)
}

// Dequeue removes item in front
func (q *Queue) Dequeue() interface{} {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
