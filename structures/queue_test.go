package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	assert.NotNil(t, q)
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	item := map[string]string{"key": "value"}
	q.Enqueue(item)
	t.Run("Length of queue has increased", func(t *testing.T) {
		assert.Equal(t, 1, q.Len())
	})
	t.Run("Item in front of queue equal local item", func(t *testing.T) {
		assert.Truef(t, assert.ObjectsAreEqual(item, q.Front()), "Expected: %v, got: %v", item, q.Front())
	})
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	item := map[string]string{"key": "value"}
	q.Enqueue(item)
	got := q.Dequeue()
	t.Run("Length of queue has decreased", func(t *testing.T) {
		assert.Equal(t, 0, q.Len())
	})
	t.Run("Item dequeued equal local item", func(t *testing.T) {
		assert.Truef(t, assert.ObjectsAreEqual(item, got), "Expected: %v, got: %v", item, got)
	})
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := NewQueue()
	got := q.Dequeue()
	t.Run("Length of queue is zero", func(t *testing.T) {
		assert.Equal(t, 0, q.Len())
	})
	t.Run("Item dequeued equal nil", func(t *testing.T) {
		assert.Nil(t, got)
	})
}
