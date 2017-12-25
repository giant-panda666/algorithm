package main

import "fmt"

//Say you have an array for which the ith element is the price of a given stock on day i.
//
//Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:
//
//You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
//After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
//Example:
//
//prices = [1, 2, 3, 0, 2]
//maxProfit = 3
//transactions = [buy, sell, cooldown, buy, sell]

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var profit = make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		profit[i] = profit[i-1]
		fmt.Println(i, profit[i])
		for k := 0; k <= i-1; k++ {
			tmp := prices[i] - prices[k]
			if k-2 > 0 {
				tmp += profit[k-2]
			}
			if tmp > profit[i] {
				profit[i] = tmp
			}
		}
	}

	return profit[len(prices)-1]
}

func main() {
	var prices = []int{1, 2}
	fmt.Println(maxProfit(prices))
}
