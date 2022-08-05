package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Event struct {
	E int
}

func (Event) Process(e Event) {
	fmt.Println(e.E)
}

func producer(id int, wg *sync.WaitGroup, EB *MutexEventBuffer) {
	defer wg.Done()

	e := Event{E: id}

	EB.Produzir(e)
}

func consumer(id int, wg *sync.WaitGroup, EB *MutexEventBuffer) {
	defer wg.Done()

	e := EB.Consumir()

	e.Process(e)
}

type MutexEventBuffer struct {
	mu       sync.Mutex
	capacity int
	buffer   []Event
}

func NewMutexEventBuffer(capacity int) *MutexEventBuffer {
	aux := &MutexEventBuffer{
		capacity: capacity,
	}
	return aux
}

func (s *MutexEventBuffer) Produzir(e Event) {
	s.mu.Lock()
	for len(s.buffer) == s.capacity {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	s.buffer = append(s.buffer, e)
	s.mu.Unlock()
}

func (s *MutexEventBuffer) Consumir() Event {
	s.mu.Lock()
	for len(s.buffer) == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	ret := s.buffer[0]
	s.buffer = s.buffer[1:]
	s.mu.Unlock()

	return ret
}

func main() {
	wg := sync.WaitGroup{}

	n := 100

	EB := NewMutexEventBuffer(1)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go consumer(i, &wg, EB)
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go producer(i, &wg, EB)
	}

	wg.Wait()

}
