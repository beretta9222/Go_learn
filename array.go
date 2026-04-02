package main

import "fmt"

func main() {
	array()
}

func array() {
	var arr = [3]int{1, 2, 3}

	for indx, item := range arr {
		fmt.Printf("%d. %d\n", indx, item)
	}
	fmt.Printf("Array length - %d", len(arr))
}
