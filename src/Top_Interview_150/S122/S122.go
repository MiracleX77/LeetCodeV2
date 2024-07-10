package main

import (
	"fmt"
)

func main() {
	nums1 := []int{
		2, 1, 2, 0, 1,
	}
	fmt.Println(maxProfit(nums1))
}

func maxProfit(prices []int) int {
	total_profit := 0
	max_profit := 0
	buy_price := prices[0]

	for _, price := range prices {
		if price < buy_price {
			if max_profit != 0 {
				total_profit += max_profit
				max_profit = 0
			}
			buy_price = price
		} else {
			profit := price - buy_price
			if profit > max_profit {
				max_profit = profit
			} else if profit < max_profit {
				total_profit += max_profit
				buy_price = price
				max_profit = 0
			}
		}

	}
	total_profit += max_profit

	return total_profit
}
