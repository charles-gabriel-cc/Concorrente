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

}

func producer(id int, wg *sync.WaitGroup, EB *MutexEventBuffer) {
	defer wg.Done()

	e := Event{E: id}

	EB.Produzir(e, id)
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

func (s *MutexEventBuffer) Produzir(e Event, id int) {
	s.mu.Lock()
	for len(s.buffer) > 0 || consumidores == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	for i := 0; i < consumidores; i++ {
		s.buffer = append(s.buffer, e)
	}
	fmt.Println("P" + strconv.Itoa(id) + " (produtor) invoca produzir")
	s.mu.Unlock()
}

func (s *MutexEventBuffer) Consumir(id int) Event {
	s.mu.Lock()
	fmt.Println("C" + strconv.Itoa(id) + " (consumidor) invoca consumir()")
	consumidores++
	for len(s.buffer) == 0 {
		s.mu.Unlock()
		runtime.Gosched()
		s.mu.Lock()
	}
	consumidores--
	ret := s.buffer[0]
	s.buffer = s.buffer[1:]
	fmt.Println("C" + strconv.Itoa(id) + " (consumidor) consome " + strconv.Itoa(ret.E))
	s.mu.Unlock()

	return ret
}

func main() {
	wg := sync.WaitGroup{}

	EB := NewMutexEventBuffer(100)

	wg.Add(1)
	go producer(3, &wg, EB)
	wg.Add(1)
	go producer(2, &wg, EB)
	wg.Add(1)
	go consumer(1, &wg, EB)
	wg.Add(1)
	go consumer(2, &wg, EB)

	wg.Wait()
}
