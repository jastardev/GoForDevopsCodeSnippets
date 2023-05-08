package main

import (
	"fmt"
	"sync"
)

type sum struct {
	mu  sync.Mutex
	sum int
}

func (s *sum) get() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.sum
}
func (s *sum) add(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sum += n
}

func main() {
	mySum := &sum{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			mySum.add(x)
		}(1)
	}

	wg.Wait()
	fmt.Println("final sum: ", mySum.get())
}
