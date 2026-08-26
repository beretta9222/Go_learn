package main

import (
	"fmt"
	"sync"
)

func main() {
	exemple1()
	exemple2()
	exemple3()
	exemple4()
}

func exemple4() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(tch chan<- int) {

		for i := 0; i < 50; i++ {
			tch <- i
		}
		close(tch)
		wg.Done()
	}(ch)
	go func(tch <-chan int) {

		for i := 0; i < 50; i++ {
			fmt.Println("result := ", <-tch)
		}
		wg.Done()
	}(ch)
	wg.Wait()
}

func exemple3() {
	ch := make(chan int)

	go func(_ch <-chan int) {
		var val int = 6
		ch <- val * val
	}(ch)

	ch <- 5
	fmt.Println("result := ", <-ch)
}

func exemple2() {
	ch := make(chan int)

	go func(_ch <-chan int) {
		val := <-_ch
		ch <- val * val
	}(ch)

	ch <- 5
	fmt.Println("result := ", <-ch)
}

func exemple1() {
	ch := make(chan int)
	go square(ch)
	ch <- 5
	fmt.Println("result := ", <-ch)
}

func square(ch chan int) {
	num := <-ch
	ch <- num * num
}
