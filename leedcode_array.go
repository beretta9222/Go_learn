import {
	"fmt",
	"math"
}

func main() {
	arr := [3]int{1, 2, 3}
	target := 9
	res1 := twoSum(arr, target)
	res2 := twoSum2(arr, target)
	res3 := 
}
/*
1. Two Sum 
https://leetcode.com/problems/two-sum/description/?envType=problem-list-v2&envId=array

You are given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/
func twoSum(nums []int, target int) []int {
	n := len(nums);
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] + nums[j] == target
				return []int{i,j}; 
		}	
	}
	return []int{};
}

func twoSum2(nums []int, target int) []int {
	dic := make(map[int]int);
	for i,x := range nums {
		need := target - x;
		if j, ok := dic[need]; ok {
			return []int{j, i}
		}
		dic[x] = i
	}	
	return []int{};
}

/*
11. Container With Most Water 
https://leetcode.com/problems/container-with-most-water/description/?envType=problem-list-v2&envId=array

You are given an integer array height of length n. 
There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container. 
*/
func MaxArea(height []int) int {
	if len(height) < 2 {
		return 0
    }

	max := -1;
	r := len(height) - 1
	l := 0

	for {
		if l > r {
			break
        }
		tmp := Min(height[l], height[r]) * (r - l);
		if max < tmp {
			max = tmp
        }
		
		if height[r] >= height[l] {
			l++;
        } else {
			r--;
        }
		
	}
	return max
}