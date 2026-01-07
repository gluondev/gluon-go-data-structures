package linear

import (
	"testing"
)

// BenchmarkEnqueue measures how fast we can fill an empty buffer.
// This is a test of raw write speed and modulo arithmetic overhead.
func BenchmarkEnqueue(b *testing.B) {
	rb := NewRingBuffer[int](b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rb.Enqueue(i)
	}
}

// BenchmarkDequeue measures how fast we can drain a full buffer.
// This also checks the cost of clearing the reference for GC.
func BenchmarkDequeue(b *testing.B) {
	rb := NewRingBuffer[int](b.N)
	for i := 0; i < b.N; i++ {
		_ = rb.Enqueue(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = rb.Dequeue()
	}
}

// BenchmarkCycle measures a common real-world scenario: 
// The buffer is under constant load with simultaneous additions and removals.
func BenchmarkCycle(b *testing.B) {
	const capacity = 1024
	rb := NewRingBuffer[int](capacity)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rb.Enqueue(i)
		_, _ = rb.Dequeue()
	}
}

// BenchmarkVaryingSizes helps identify if performance degrades as the buffer gets larger.
// This tests CPU cache efficiency (L1/L2/L3 cache misses).
func BenchmarkVaryingSizes(b *testing.B) {
	sizes := []int{64, 1024, 8192, 65536}
	for _, size := range sizes {
		b.Run(string(rune(size)), func(b *testing.B) {
			rb := NewRingBuffer[int](size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = rb.Enqueue(i)
				if i%size == 0 && i > 0 {
					for j := 0; j < size; j++ {
						_, _ = rb.Dequeue()
					}
				}
			}
		})
	}
}