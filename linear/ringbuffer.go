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

// Cap returns the fixed capacity of the buffer.
func (r *RingBuffer[T]) Cap() int {
	return len(r.data)
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

// TryEnqueue adds an element to the end of the buffer and reports whether it succeeded.
func (r *RingBuffer[T]) TryEnqueue(item T) bool {
	if r.size == len(r.data) {
		return false
	}
	r.data[r.tail] = item
	r.tail = (r.tail + 1) % len(r.data)
	r.size++
	return true
}

// EnqueueOverwrite adds an element to the end of the buffer, overwriting the oldest element if full.
func (r *RingBuffer[T]) EnqueueOverwrite(item T) {
	r.data[r.tail] = item
	if r.size == len(r.data) {
		r.head = (r.head + 1) % len(r.data)
	} else {
		r.size++
	}
	r.tail = (r.tail + 1) % len(r.data)
}

// Peek returns the front element without removing it.
func (r *RingBuffer[T]) Peek() (T, error) {
	var zero T
	if r.size == 0 {
		return zero, ErrBufferEmpty
	}
	return r.data[r.head], nil
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

// TryDequeue removes and returns the front element and reports whether it succeeded.
func (r *RingBuffer[T]) TryDequeue() (T, bool) {
	var zero T
	if r.size == 0 {
		return zero, false
	}

	item := r.data[r.head]
	r.data[r.head] = zero // Clear reference for GC
	r.head = (r.head + 1) % len(r.data)
	r.size--
	return item, true
}

// Reset empties the buffer without clearing the underlying storage.
func (r *RingBuffer[T]) Reset() {
	r.head = 0
	r.tail = 0
	r.size = 0
}

// Clear empties the buffer and zeroes its underlying storage to release references for GC.
func (r *RingBuffer[T]) Clear() {
	var zero T
	for i := range r.data {
		r.data[i] = zero
	}
	r.Reset()
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
