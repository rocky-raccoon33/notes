package main

import "fmt"

func pick_stone(n int) bool {
	if n <= 2 {
		return true
	}

	dp := make([]bool, n+1)
	dp[1], dp[2] = true, true
	i := 3
	for i <= n {
		dp[i] = !(dp[i-1] && dp[i-2])
		i++
	}
	return dp[n]

}

func main() {
	n := 150000
	output := "win"
	if !pick_stone(n) {
		output = "lose"
	}
	fmt.Printf("input total stones %d: your %s！", n, output)
}
