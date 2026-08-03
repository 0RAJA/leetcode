// 使用 channel 来实现读写锁
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const RWMutexMaxReaders = 1 << 30 // 一个无法达到的最大读数量

type RWMutex struct {
	mu          chan struct{} // 互斥锁
	readerCount int32         // 读数量
	readerWait  int32         // 写等待的读的数量
	wChan       chan struct{} // 用于唤醒等待读的写
	rChan       chan struct{} // 用于唤醒等待写的读
}

func NewRWMutex() *RWMutex {
	return &RWMutex{mu: make(chan struct{}, 1), wChan: make(chan struct{}), rChan: make(chan struct{})}
}

func (rw *RWMutex) Lock() {
	rw.mu <- struct{}{} // 获取锁
	// 阻止之后的读操作，等待现有的读操作
	if r := atomic.AddInt32(&rw.readerCount, -RWMutexMaxReaders) + RWMutexMaxReaders; r > 0 {
		atomic.AddInt32(&rw.readerWait, r) // 增加写阻塞时读数量
		<-rw.wChan
	}
}

func (rw *RWMutex) Unlock() {
	// 唤醒等待的读
	if r := atomic.AddInt32(&rw.readerCount, RWMutexMaxReaders); r > 0 {
		for i := 0; i < int(r); i++ {
			rw.rChan <- struct{}{}
		}
	}
	// 解锁
	<-rw.mu
}

func (rw *RWMutex) RLock() {
	// 增加读数量，如果有写就等待写
	if r := atomic.AddInt32(&rw.readerCount, 1); r < 0 {
		<-rw.rChan
	}
}

func (rw *RWMutex) RUnlock() {
	// 减少读数量，有写等待就进一步判断如果自己是最后一个读就唤醒写
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
		if rwait := atomic.AddInt32(&rw.readerWait, -1); rwait == 0 {
			rw.wChan <- struct{}{}
		}
	}
}

func main() {
	rw := NewRWMutex()
	num := 0
	wg := new(sync.WaitGroup)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			switch i % 2 {
			case 0:
				rw.Lock()
				defer rw.Unlock()
				num++
				fmt.Println("Lock i:", i, "num:", num)
			case 1:
				time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
				rw.RLock()
				defer rw.RUnlock()
				fmt.Println("RLock i:", i, "num:", num)
			}
		}(i)
	}
	wg.Wait()
}
