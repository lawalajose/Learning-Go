package queue

type Queue struct {
	item     []int
	capacity int
}

func New(capacity int) Queue {
	return Queue{make([]int, 0, capacity), capacity}
}

func (q *Queue) Append(item int) bool {
	if len(q.item) == q.capacity {
		return false
	}
	q.item = append(q.item, item)
	return true
}

func (q *Queue) Next() (int, bool) {
	if len(q.item) > 1 {
		next := q.item[0]
		q.item = q.item[1:]
		return next, true
	} else {
		return 0, false
	}
}
