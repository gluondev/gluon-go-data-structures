package linear

import "errors"

// ErrBufferFull is returned when enqueueing to a full buffer.
var ErrBufferFull = errors.New("ring buffer is full")

// ErrBufferEmpty is returned when dequeueing from an empty buffer.
var ErrBufferEmpty = errors.New("ring buffer is empty")

// RingBuffer is a fixed-size circular queue using Go generics.
type RingBuffer[T any] struct {
	data []T
	head int
	tail int
	size int
}

// NewRingBuffer initializes a buffer with a fixed capacity.
// It panics if capacity is not positive.
func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity <= 0 {
		panic("ring buffer capacity must be positive")
	}
	return &RingBuffer[T]{
		data: make([]T, capacity),
	}
}

// Enqueue adds an element to the end of the buffer.
func (r *RingBuffer[T]) Enqueue(item T) error {
	if r.size == len(r.data) {
		return ErrBufferFull
	}

	r.data[r.tail] = item
	r.tail = (r.tail + 1) % len(r.data)
	r.size++
	return nil
}

// Dequeue removes and returns the front element.
func (r *RingBuffer[T]) Dequeue() (T, error) {
	var zero T
	if r.size == 0 {
		return zero, ErrBufferEmpty
	}

	item := r.data[r.head]
	r.data[r.head] = zero // Clear reference for GC
	r.head = (r.head + 1) % len(r.data)
	r.size--
	return item, nil
}

// IsFull returns true if the buffer is at capacity.
func (r *RingBuffer[T]) IsFull() bool {
	return r.size == len(r.data)
}

// IsEmpty returns true if the buffer has no elements.
func (r *RingBuffer[T]) IsEmpty() bool {
	return r.size == 0
}

// Len returns the current number of elements.
func (r *RingBuffer[T]) Len() int {
	return r.size
}
