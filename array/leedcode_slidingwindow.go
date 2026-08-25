
package main

import (
    "fmt"
)

func main() {
	arr := []int { 1 , 2, 3, 1}
	k := 3
	res := ContainsDuplicate(arr, k)
	fmt.Printf("%t", res)
}

/*
219. Contains Duplicate II
Given an integer array nums and an integer k, 
return true if there are two distinct indices i and j 
in the array such that nums[i] == nums[j] 
and abs(i - j) <= k.
*/
func ContainsDuplicate(arr []int, k int) bool {
	queue := make(map[int]int)

	for i, val := range arr {
		if idx, ok := queue[val]; ok && i-idx <= k {
			return true
		}
		queue[val] = i
	}
	return false
}
