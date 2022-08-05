package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var consumidores = 0

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

	e := EB.Consumir(id)

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
	for len(s.buffer) == s.capacity && consumidores == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	for i := 0; i < consumidores; i++ {
		s.buffer = append(s.buffer, e)
	}
	s.mu.Unlock()
}

func (s *MutexEventBuffer) Consumir(id int) Event {
	s.mu.Lock()
	consumidores++
	fmt.Println("C" + strconv.Itoa(id) + " (consumidor) invoca consumir() e fica bloqueado")
	for len(s.buffer) == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	consumidores--
	ret := s.buffer[0]
	fmt.Println("C" + strconv.Itoa(id) + " (consumidor) consome " + strconv.Itoa(ret.E))
	s.buffer = s.buffer[1:]
	s.mu.Unlock()

	return ret
}

func main() {
	wg := sync.WaitGroup{}

	EB := NewMutexEventBuffer(100)

	wg.Add(1)
	go consumer(1, &wg, EB)
	wg.Add(1)
	go consumer(2, &wg, EB)
	wg.Add(1)
	go producer(2, &wg, EB)

	wg.Wait()
}
