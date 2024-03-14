package main

import "fmt"

func main() {
	fmt.Println(max([]int{3, 5, 1, 9, 2}))
}

func max(value []int) int {
	max := 0
	for i := 0; i < len(value); i++ {
		if value[i] > max {
			max = value[i]
		}
	}
	return max
}
