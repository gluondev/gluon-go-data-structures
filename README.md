# Gluon Go Data Structures ⚛️

High-performance, generic, and systems-oriented data structures for Go. 

This repository is part of the [gluondev](https://github.com/gluondev) ecosystem, focusing on the "glue" that powers modern infrastructure: efficient, safe, and cache-friendly data management.

## 🛠 Project Philosophy
While the Go standard library is excellent, it is intentionally minimalist. This library aims to provide:
- **Type Safety**: Full use of Go 1.18+ Generics.
- **Performance**: Minimized allocations and optimized for CPU cache locality.
- **Clarity**: Clean, idiomatic Go code that is easy to audit and extend.

---

## 🗺 Data Structures Roadmap

### 1. Linear Data Structures
*Focus: Memory efficiency and contiguous storage.*
- [ ] **Dynamic Array (Vector)**: Custom implementation with manual growth control.
- [ ] **Ring Buffer (Circular Buffer)**: Fixed-capacity, zero-allocation buffer for streaming.
- [ ] **BitSet / BitMap**: Space-efficient boolean array using bitwise operations.
- [ ] **Gap Buffer**: Optimized for text editors and frequent local edits.
- [ ] **Doubly Linked List**: Standard pointer-based list for LRU implementations.



### 2. Search & Key-Value Structures
*Focus: High-speed lookups and prefix matching.*
- [ ] **Trie (Prefix Tree)**: Standard string-based search tree for routing.
- [ ] **Radix Tree**: Space-optimized trie (compressed edges).
- [ ] **Skip List**: Probabilistic ordered structure (concurrent-friendly).
- [ ] **Swiss Map**: High-performance hash map implementation using SIMD-like techniques.

### 3. Concurrency-Safe Structures
*Focus: Data flow between goroutines with minimal contention.*
- [ ] **Thread-Safe Map**: Sharded implementation to reduce mutex lock contention.
- [ ] **Lock-Free Queue**: Using `sync/atomic` for ultra-low latency.
- [ ] **Priority Queue (Binary Heap)**: Essential for task scheduling.
- [ ] **Blocking Queue**: Producer-consumer queue with wait/notify mechanics.



### 4. Advanced Trees & Indices
*Focus: Complex data organization and range queries.*
- [ ] **B-Tree / B+ Tree**: High-branching factor trees for storage engines.
- [ ] **Red-Black Tree**: Self-balancing binary search tree.
- [ ] **Segment Tree**: Efficient for $O(\log n)$ range queries.
- [ ] **Interval Tree**: Storing and querying overlapping ranges.



### 5. Probabilistic Data Structures
*Focus: Handling massive datasets with constant memory footprints.*
- [ ] **Bloom Filter**: Fast membership testing with zero false negatives.
- [ ] **Cuckoo Filter**: Supports deletions and better cache locality than Bloom.
- [ ] **HyperLogLog**: Unique element count (cardinality) estimation.
- [ ] **Count-Min Sketch**: Frequency estimation in data streams.

### 6. Specialized & System Structures
*Focus: Integrity, Memory, and Cache management.*
- [ ] **Merkle Tree**: Hash-based verification for distributed systems.
- [ ] **Arena Allocator**: Manual memory block management for high-speed allocation.
- [ ] **LRU / LFU Cache**: Intelligent data eviction policies.
- [ ] **QuadTree**: Spatial partitioning for 2D coordinate indexing.

---

## 🚀 Getting Started

### Installation
```bash
go get [github.com/gluondev/gluon-go-data-structures](https://github.com/gluondev/gluon-go-data-structures)
```

### Build a Demo Binary
This repo is primarily a library, so packages like `./linear` won't produce a standalone binary by themselves. A small demo `main` package is available under `./cmd`.

```bash
go build -o ringbuffer-demo ./cmd/ringbuffer-demo
./ringbuffer-demo -capacity 3
```
