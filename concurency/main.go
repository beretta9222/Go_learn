package main

import (
	"fmt"
	"sync"
)

func main() {
	go WaitGroupExemple()
	go MutexExemple()
}

func WaitGroupExemple() {
	var twg sync.WaitGroup
	twg.Add(6)

	for i := 0; i < 6; i++ {
		go func(ind int) {
			defer twg.Done()
			fmt.Println(ind)
		}(i)
	}
	twg.Wait()
}

var counter int = 0
var wg sync.WaitGroup

func MutexExemple() {

	var mutex sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(ind int, mutex *sync.Mutex) {

			mutex.Lock()
			var tmp_counter int = counter
			counter = 0
			for j := 1; j < ind; j++ {
				counter++
				fmt.Println("Goroutine", ind, "-", counter)
			}

			counter += tmp_counter
			fmt.Println(ind)
			mutex.Unlock()

			wg.Done()
		}(i, &mutex)
	}
	wg.Wait()
	fmt.Println(counter)
}
