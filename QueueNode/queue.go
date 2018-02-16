package queueNode

import "fmt"

type queueNode struct {
	TElement interface{}
	Next     *queueNode
}

// Queue build queue
type Queue struct {
	headNode *queueNode
	tailNode *queueNode
	size     int
}

// Enqueue adds item to queue
func (q *Queue) Enqueue(item interface{}) {
	newNode := &queueNode{TElement: item, Next: nil}
	if q.Size() == 0 {
		q.headNode = newNode
	} else {
		q.tailNode.Next = newNode
	}
	q.tailNode = newNode
	q.size++
}

// Dequeue removes and returns item from queue
func (q *Queue) Dequeue() interface{} {
	if q.Size() > 0 {
		oldNode := q.headNode
		if q.headNode.Next != nil {
			q.headNode = q.headNode.Next
		}
		q.size--
		return oldNode.TElement
	}
	return "Empty Queue Error"
}

// Size of the queue
func (q *Queue) Size() int {
	return q.size
}

// NewQueue creates a new queue
func NewQueue(items ...interface{}) *Queue {
	queue := new(Queue)
	for _, item := range items {
		queue.Enqueue(item)
	}
	return queue
}

func (q *Queue) String() string {
	var items string
	current := q.headNode
	for i := 0; i < q.Size(); i++ {
		items += fmt.Sprintf("%v", current.TElement)
		if i < q.Size()-1 {
			items += ", "
		}
		current = current.Next
	}
	return "[" + items + "]"
}
