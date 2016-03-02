package queue

import (
	"fmt"
	"sync"
)

type Queue struct {
	mutex sync.Mutex

	nodes []interface{}
	size  int
	head  int
	tail  int
	count int
}

func NewQueue(s int) *Queue {
	q := &Queue{
		nodes: make([]interface{}, s),
		size:  s,
	}
	return q
}

func (q Queue) String() string {
	return fmt.Sprintf("Queue %v", q.nodes)
}

// Push adds a node to the queue.
func (q *Queue) Push(n interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.head == q.tail && q.count > 0 {
		nodes := make([]interface{}, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.nodes[q.head] = nil
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func (q Queue) Len() int {
	return q.count
}
