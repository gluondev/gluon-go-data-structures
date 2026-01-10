package linear

import (
	"fmt"
	"testing"
)

func TestRingBuffer_Basic(t *testing.T) {
	buf := NewRingBuffer[int](3)

	if got := buf.Cap(); got != 3 {
		t.Fatalf("Cap() = %d, want 3", got)
	}
	if !buf.IsEmpty() {
		t.Fatalf("expected empty buffer")
	}

	if err := buf.Enqueue(1); err != nil {
		t.Fatalf("Enqueue(1): %v", err)
	}
	if err := buf.Enqueue(2); err != nil {
		t.Fatalf("Enqueue(2): %v", err)
	}
	if err := buf.Enqueue(3); err != nil {
		t.Fatalf("Enqueue(3): %v", err)
	}
	if !buf.IsFull() {
		t.Fatalf("expected full buffer")
	}
	if err := buf.Enqueue(4); err != ErrBufferFull {
		t.Fatalf("Enqueue(4) error = %v, want %v", err, ErrBufferFull)
	}

	if got, err := buf.Peek(); err != nil || got != 1 {
		t.Fatalf("Peek() = (%v, %v), want (1, nil)", got, err)
	}

	v, err := buf.Dequeue()
	if err != nil || v != 1 {
		t.Fatalf("Dequeue() = (%v, %v), want (1, nil)", v, err)
	}
	v, err = buf.Dequeue()
	if err != nil || v != 2 {
		t.Fatalf("Dequeue() = (%v, %v), want (2, nil)", v, err)
	}

	if err := buf.Enqueue(4); err != nil {
		t.Fatalf("Enqueue(4): %v", err)
	}
	if err := buf.Enqueue(5); err != nil {
		t.Fatalf("Enqueue(5): %v", err)
	}

	v, err = buf.Dequeue()
	if err != nil || v != 3 {
		t.Fatalf("Dequeue() = (%v, %v), want (3, nil)", v, err)
	}
	v, err = buf.Dequeue()
	if err != nil || v != 4 {
		t.Fatalf("Dequeue() = (%v, %v), want (4, nil)", v, err)
	}
	v, err = buf.Dequeue()
	if err != nil || v != 5 {
		t.Fatalf("Dequeue() = (%v, %v), want (5, nil)", v, err)
	}

	if _, err := buf.Dequeue(); err != ErrBufferEmpty {
		t.Fatalf("Dequeue() error = %v, want %v", err, ErrBufferEmpty)
	}
}

func TestRingBuffer_Try(t *testing.T) {
	buf := NewRingBuffer[int](2)

	if got := buf.TryEnqueue(1); !got {
		t.Fatalf("TryEnqueue(1) = false, want true")
	}
	if got := buf.TryEnqueue(2); !got {
		t.Fatalf("TryEnqueue(2) = false, want true")
	}
	if got := buf.TryEnqueue(3); got {
		t.Fatalf("TryEnqueue(3) = true, want false")
	}

	v, ok := buf.TryDequeue()
	if !ok || v != 1 {
		t.Fatalf("TryDequeue() = (%v, %v), want (1, true)", v, ok)
	}
	v, ok = buf.TryDequeue()
	if !ok || v != 2 {
		t.Fatalf("TryDequeue() = (%v, %v), want (2, true)", v, ok)
	}
	_, ok = buf.TryDequeue()
	if ok {
		t.Fatalf("TryDequeue() ok = true, want false")
	}
}

func TestRingBuffer_EnqueueOverwrite(t *testing.T) {
	buf := NewRingBuffer[int](3)
	_ = buf.Enqueue(1)
	_ = buf.Enqueue(2)
	_ = buf.Enqueue(3)

	buf.EnqueueOverwrite(4)

	v, _ := buf.Dequeue()
	if v != 2 {
		t.Fatalf("Dequeue() = %d, want 2", v)
	}
	v, _ = buf.Dequeue()
	if v != 3 {
		t.Fatalf("Dequeue() = %d, want 3", v)
	}
	v, _ = buf.Dequeue()
	if v != 4 {
		t.Fatalf("Dequeue() = %d, want 4", v)
	}
}

func TestRingBuffer_ClearAndReset(t *testing.T) {
	buf := NewRingBuffer[int](2)
	_ = buf.Enqueue(1)
	_ = buf.Enqueue(2)

	buf.Reset()
	if !buf.IsEmpty() || buf.Len() != 0 {
		t.Fatalf("after Reset expected empty buffer, got Len=%d", buf.Len())
	}

	_ = buf.Enqueue(3)
	buf.Clear()
	if !buf.IsEmpty() || buf.Len() != 0 {
		t.Fatalf("after Clear expected empty buffer, got Len=%d", buf.Len())
	}
}

func TestNewRingBuffer_InvalidCapacityPanics(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Fatalf("expected panic for capacity 0")
			}
		}()
		_ = NewRingBuffer[int](0)
	})
	t.Run("negative", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Fatalf("expected panic for negative capacity")
			}
		}()
		_ = NewRingBuffer[int](-1)
	})
}

func ExampleRingBuffer() {
	buf := NewRingBuffer[string](2)
	buf.EnqueueOverwrite("a")
	buf.EnqueueOverwrite("b")
	buf.EnqueueOverwrite("c") // overwrites "a"

	for !buf.IsEmpty() {
		v, _ := buf.Dequeue()
		fmt.Println(v)
	}

	// Output:
	// b
	// c
}

