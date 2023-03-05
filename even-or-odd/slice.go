package main

import "fmt"

type slice []int

func newSlice(n int) slice {
	nums := slice{}

	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}

	return nums
}

func (s slice) print() {
	for i, num := range s {
		fmt.Printf("%d: %d\n", i, num)
	}
}

func evenOrOdd(s slice) {
	for _, num := range s {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}
