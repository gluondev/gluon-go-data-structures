package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gluondev/gluon-go-data-structures/linear"
)

func main() {
	capacity := flag.Int("capacity", 3, "ring buffer capacity (must be > 0)")
	flag.Parse()

	buf := linear.NewRingBuffer[string](*capacity)

	mustEnqueue := func(s string) {
		if err := buf.Enqueue(s); err != nil {
			fmt.Fprintf(os.Stderr, "enqueue %q: %v\n", s, err)
		}
	}
	mustDequeue := func() {
		v, err := buf.Dequeue()
		if err != nil {
			fmt.Fprintf(os.Stderr, "dequeue: %v\n", err)
			return
		}
		fmt.Printf("dequeue -> %q (len=%d)\n", v, buf.Len())
	}

	fmt.Printf("ringbuffer-demo (capacity=%d)\n", *capacity)
	fmt.Println("enqueue: a, b, c; dequeue twice; enqueue d; dequeue until empty")

	mustEnqueue("a")
	mustEnqueue("b")
	mustEnqueue("c")
	mustDequeue()
	mustDequeue()
	mustEnqueue("d")
	for !buf.IsEmpty() {
		mustDequeue()
	}
}

