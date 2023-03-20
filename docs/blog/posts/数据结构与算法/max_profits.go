package main

import (
	"fmt"
	"math"
)

func main() {
	// [3,3,5,0,0,3,1,4]
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}

func maxProfit(prices []int) int {
	max := 0
	for i := range prices {
		tmp := iter(prices[:i]) + iter(prices[i:])
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func iter(prices []int) int {

	min := math.MaxInt
	profit := 0
	for _, v := range prices {
		if v < min {
			min = v
		}

		if profit < v-min {
			profit = v - min
		}
	}
	return profit
}
