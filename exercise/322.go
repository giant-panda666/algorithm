package main

import "fmt"

//You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
//
//Example 1:
//coins = [1, 2, 5], amount = 11
//return 3 (11 = 5 + 5 + 1)
//
//Example 2:
//coins = [2], amount = 3
//return -1.
//
//Note:
//You may assume that you have an infinite number of each kind of coin.

// f(amount, n) 总金额为amount时，从n种面额的硬币种选择的最少数量，使得这些硬币的总额不超过amount
func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				if dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	var coins = []int{1, 2}
	var amount = 3
	fmt.Println(coinChange(coins, amount))
}
