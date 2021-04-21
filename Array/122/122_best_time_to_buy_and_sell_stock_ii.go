/*  problem 122: Best Time to Buy and Sell Stock II
    - card ref: https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/564/
    - solution ref: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/solution/
*/
package main

import "fmt"

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{1, 2, 3, 4, 5}
	prices3 := []int{7, 6, 4, 3, 1}
	prices4 := []int{1, 7, 2, 3, 6, 7, 6, 7}

	//     fmt.Println(maxProfitBf(prices1, 0))
	//     fmt.Println(maxProfitBf(prices2, 0))
	//     fmt.Println(maxProfitBf(prices3, 0))
	fmt.Println(maxProfitBf(prices4, 0))

	//     fmt.Println(maxProfitPeak(prices1))
	//     fmt.Println(maxProfitPeak(prices2))
	//     fmt.Println(maxProfitPeak(prices3))
	fmt.Println(maxProfitPeak(prices4))

	fmt.Println(maxProfitSimpleOnePass(prices1))
	fmt.Println(maxProfitSimpleOnePass(prices2))
	fmt.Println(maxProfitSimpleOnePass(prices3))
	fmt.Println(maxProfitSimpleOnePass(prices4))
}

// 硬解
func maxProfitBf(prices []int, b int) (maxProfit int) {
	maxProfit = 0
	maxDay := len(prices)
	for ; b < maxDay; b++ {
		mb := 0 // max when buy at b day
		for s := b + 1; s < maxDay; s++ {
			// sell when positive profit
			if prices[s] > prices[b] {
				profit := (prices[s] - prices[b]) + maxProfitBf(prices, s+1) // recursive from s+1, cuz cannot sell and buy in same day
				// remain max profit (decide to use which sell day to make profit max)
				if profit > mb {
					mb = profit
				}
			}
		}
		if mb > maxProfit {
			maxProfit = mb
		}
	}
	return
}

// 核心概念, 找出連續的 valley + peak
func maxProfitPeak(prices []int) (m int) {
	var p, v int
	vi, pi := -1, -1
	for i, price := range prices {
		if i == len(prices)-1 {
			if vi > pi && price > v { // edge case
				m += (price - v)
			}
			break
		}
		// 先訂出 valley
		if vi <= pi && price < prices[i+1] {
			vi = i
			v = price
		}
		// 訂出 peak 之後, 算 profit
		if vi > pi && price > prices[i+1] {
			pi = i
			p = price
			m += (p - v)
		}
	}
	return
}

func maxProfitSimpleOnePass(prices []int) (m int) {
	for i, p := range prices {
		if i == 0 {
			continue
		}
		if p > prices[i-1] {
			m += (p - prices[i-1])
		}
	}
	return
}
