package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// LKQueue 无锁队列
type LKQueue struct {
	head unsafe.Pointer // 第一个
	tail unsafe.Pointer // 最后一个
}

// 节点
type node struct {
	value interface{}    // 当前 value
	next  unsafe.Pointer // next
}

// NewLKQueue returns an empty queue.
func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{head: n, tail: n}
}

// 原子读
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// cas
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}

// Enqueue 写入元素到队列
func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{value: v}
	for {
		tail := load(&q.tail)      // 当前最后一个
		next := load(&tail.next)   // 最后一个的下一个
		if tail == load(&q.tail) { // 判断是否被修改
			if next == nil { // 这是最后一个
				if cas(&tail.next, next, n) {
					cas(&q.tail, tail, n) // 尝试将当前队列末尾指向这个节点，失败说明有其他的操作了
					return
				}
			} else { // 不是最后一个 试着把tail摆动到下一个节点
				cas(&q.tail, tail, next)
			}
		}
	}
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *LKQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)      // 头
		tail := load(&q.tail)      // 尾
		next := load(&head.next)   // 头的下一个
		if head == load(&q.head) { // are head, tail, and next consistent?
			if head == tail { // is queue empty or tail falling behind?
				if next == nil { // is queue empty?
					return nil
				}
				// tail is falling behind.  try to advance it
				cas(&q.tail, tail, next)
			} else {
				// read value before CAS otherwise another dequeue might free the next node
				v := next.value
				if cas(&q.head, head, next) {
					return v // Dequeue is done.  return
				}
			}
		}
	}
}

type queueCheckResult struct {
	Producers        int
	Consumers        int
	ItemsPerProducer int
	Expected         uint64
	Produced         uint64
	Consumed         uint64
	Missing          uint64
	Duplicates       uint64
	Invalid          uint64
	Leftover         uint64
	TimedOut         bool
	Duration         time.Duration
}

func (r queueCheckResult) OK() bool {
	return !r.TimedOut &&
		r.Produced == r.Expected &&
		r.Consumed == r.Expected &&
		r.Missing == 0 &&
		r.Duplicates == 0 &&
		r.Invalid == 0 &&
		r.Leftover == 0
}

func runConcurrentQueueCheck(producers, consumers, itemsPerProducer int, timeout time.Duration) queueCheckResult {
	q := NewLKQueue()
	total := producers * itemsPerProducer
	seen := make([]uint32, total)

	var produced uint64
	var consumed uint64
	var duplicates uint64
	var invalid uint64
	var producersDone uint32

	start := time.Now()
	var producerWG sync.WaitGroup
	producerWG.Add(producers)
	for p := 0; p < producers; p++ {
		p := p
		go func() {
			defer producerWG.Done()
			base := p * itemsPerProducer
			for i := 0; i < itemsPerProducer; i++ {
				q.Enqueue(base + i)
				atomic.AddUint64(&produced, 1)
			}
		}()
	}

	var consumerWG sync.WaitGroup
	consumerWG.Add(consumers)
	for c := 0; c < consumers; c++ {
		go func() {
			defer consumerWG.Done()
			for {
				v := q.Dequeue()
				if v == nil {
					if atomic.LoadUint32(&producersDone) == 1 {
						return
					}
					runtime.Gosched()
					continue
				}

				id, ok := v.(int)
				if !ok || id < 0 || id >= total {
					atomic.AddUint64(&invalid, 1)
					atomic.AddUint64(&consumed, 1)
					continue
				}

				if atomic.AddUint32(&seen[id], 1) != 1 {
					atomic.AddUint64(&duplicates, 1)
				}
				atomic.AddUint64(&consumed, 1)
			}
		}()
	}

	producerWG.Wait()
	atomic.StoreUint32(&producersDone, 1)

	done := make(chan struct{})
	go func() {
		consumerWG.Wait()
		close(done)
	}()

	timedOut := false
	select {
	case <-done:
	case <-time.After(timeout):
		timedOut = true
	}

	var missing uint64
	var multiSeen uint64
	for i := range seen {
		count := atomic.LoadUint32(&seen[i])
		switch {
		case count == 0:
			missing++
		case count > 1:
			multiSeen += uint64(count - 1)
		}
	}

	leftover := uint64(0)
	if !timedOut {
		for q.Dequeue() != nil {
			leftover++
		}
	}

	result := queueCheckResult{
		Producers:        producers,
		Consumers:        consumers,
		ItemsPerProducer: itemsPerProducer,
		Expected:         uint64(total),
		Produced:         atomic.LoadUint64(&produced),
		Consumed:         atomic.LoadUint64(&consumed),
		Missing:          missing,
		Duplicates:       atomic.LoadUint64(&duplicates) + multiSeen,
		Invalid:          atomic.LoadUint64(&invalid),
		Leftover:         leftover,
		TimedOut:         timedOut,
		Duration:         time.Since(start),
	}
	return result
}

func printQueueCheckResult(r queueCheckResult) {
	rate := float64(0)
	if r.Duration > 0 {
		rate = float64(r.Consumed) / r.Duration.Seconds()
	}

	fmt.Printf("producers=%d consumers=%d items_per_producer=%d expected=%d\n",
		r.Producers, r.Consumers, r.ItemsPerProducer, r.Expected)
	fmt.Printf("produced=%d consumed=%d duration=%s throughput=%.0f items/s\n",
		r.Produced, r.Consumed, r.Duration.Round(time.Millisecond), rate)
	fmt.Printf("missing=%d duplicates=%d invalid=%d leftover=%d timed_out=%t\n",
		r.Missing, r.Duplicates, r.Invalid, r.Leftover, r.TimedOut)

	if r.OK() {
		fmt.Println("PASS: concurrent enqueue/dequeue integrity check passed")
		return
	}
	fmt.Println("FAIL: concurrent enqueue/dequeue integrity check failed")
}

func main() {
	producers := runtime.GOMAXPROCS(0)
	if producers < 4 {
		producers = 4
	}

	result := runConcurrentQueueCheck(producers, producers, 100000, 15*time.Second)
	printQueueCheckResult(result)
	if !result.OK() {
		os.Exit(1)
	}
}
