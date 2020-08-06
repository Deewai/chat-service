package structures

import (
	"container/list"
)

// Queue - Custom List type
type Queue struct {
	*list.List
}

// Enqueue - Used to add new item to the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	q.PushBack(value)
}

// Len - Get the no of items currently in the list
func (q *Queue) Len() int {
	return q.List.Len()
}

// Dequeue - Remove the first item from the queue
func (q *Queue) Dequeue() interface{} {
	if q.Len() == 0 {
		return nil
	}
	front := q.List.Front()
	val := front.Value
	q.Remove(front)
	return val
}

// Front - Peak the first item of the queue without removing it
func (q *Queue) Front() interface{} {
	return q.List.Front().Value
}

// NewQueue - Create a new queue
func NewQueue() *Queue {
	return &Queue{
		List: list.New().Init(),
	}
}
