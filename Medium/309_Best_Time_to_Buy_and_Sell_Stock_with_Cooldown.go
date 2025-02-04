/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

You are given an array prices where prices[i] is the price
of a given stock on the ith day.

Find the maximum profit you can achieve.
You may complete as many transactions as you like
(i.e., buy one and sell one share of the stock multiple times)
with the following restrictions:

After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
*/

package Medium

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	hold := -prices[0] // Buy stock
	sell := 0          // Sell stock
	cooldown := 0      // Wait`

	for i := 1; i < n; i++ {
		// Save previous condition of cooldown
		prevCooldown := cooldown

		// Update cond. of cooldown
		cooldown = max(cooldown, sell)

		// Update cond. of sell
		sell = hold + prices[i]

		// Update cond. of hold
		hold = max(hold, prevCooldown-prices[i])
	}

	// Maximum: in the sell or in the hold condition
	return max(sell, cooldown)
}
