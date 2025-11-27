package structures

type Queue struct {
	data []string
}

func NewQueue() *Queue {
	return &Queue{data: []string{}}
}

func (q *Queue) Push(val string) {
	q.data = append(q.data, val)
}

func (q *Queue) Pop() string {
	if len(q.data) == 0 {
		return ""
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}
