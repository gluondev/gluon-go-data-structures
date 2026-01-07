package linear

import (
	"errors"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	t.Run("Initialization", func(t *testing.T) {
		capacity := 5
		rb := NewRingBuffer[int](capacity)

		if rb.Len() != 0 {
			t.Errorf("Expected initial length 0, got %d", rb.Len())
		}
		if !rb.IsEmpty() {
			t.Error("New buffer should be empty")
		}
		if rb.IsFull() {
			t.Error("New buffer should not be full")
		}
	})

	t.Run("EnqueueAndDequeue", func(t *testing.T) {
		rb := NewRingBuffer[int](3)

		// Basic Enqueue
		rb.Enqueue(1)
		rb.Enqueue(2)

		if rb.Len() != 2 {
			t.Errorf("Expected length 2, got %d", rb.Len())
		}

		// Basic Dequeue
		val, err := rb.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error on dequeue: %v", err)
		}
		if val != 1 {
			t.Errorf("Expected 1, got %d", val)
		}
		if rb.Len() != 1 {
			t.Errorf("Expected length 1, got %d", rb.Len())
		}
	})

	t.Run("FullAndEmptyErrors", func(t *testing.T) {
		rb := NewRingBuffer[int](2)

		// Test Empty Dequeue
		_, err := rb.Dequeue()
		if !errors.Is(err, ErrBufferEmpty) {
			t.Errorf("Expected ErrBufferEmpty, got %v", err)
		}

		// Fill the buffer
		rb.Enqueue(10)
		rb.Enqueue(20)

		if !rb.IsFull() {
			t.Error("Buffer should be full now")
		}

		// Test Full Enqueue
		err = rb.Enqueue(30)
		if !errors.Is(err, ErrBufferFull) {
			t.Errorf("Expected ErrBufferFull, got %v", err)
		}
	})

	t.Run("WrapAroundLogic", func(t *testing.T) {
		// This is the most critical test for a circular buffer
		rb := NewRingBuffer[int](3)

		// Fill and empty partially to move head/tail
		rb.Enqueue(1) // [1, _, _]
		rb.Enqueue(2) // [1, 2, _]
		rb.Dequeue()  // [_, 2, _] head is at index 1

		rb.Enqueue(3) // [_, 2, 3]
		rb.Enqueue(4) // [4, 2, 3] tail wraps to index 0

		if !rb.IsFull() {
			t.Error("Buffer should be full after wrap-around enqueue")
		}

		// Dequeue all and check order
		results := []int{}
		for !rb.IsEmpty() {
			v, _ := rb.Dequeue()
			results = append(results, v)
		}

		expected := []int{2, 3, 4}
		for i, v := range results {
			if v != expected[i] {
				t.Errorf("At index %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("GenericsSupport", func(t *testing.T) {
		type custom struct {
			ID   int
			Name string
		}
		rb := NewRingBuffer[custom](1)
		
		val := custom{ID: 1, Name: "Gluon"}
		rb.Enqueue(val)
		
		out, _ := rb.Dequeue()
		if out.Name != "Gluon" {
			t.Errorf("Generics failed, expected 'Gluon', got %s", out.Name)
		}
	})
}