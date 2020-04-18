package main
import "fmt"

// 2020.4.6 思路：
// 只要价格比前一天大，那么【前一天买入后一天卖出，同时利润做累加】等价于【在递增价格中最低买入，最高卖出】
// 若价格小于前一天，那么不做操作就不会亏损
func maxProfit(prices []int) int {
    if len(prices) <= 1 { return 0 }
    maxProfit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            maxProfit += prices[i] - prices[i-1]
        }
    }
    return maxProfit
}
func main() {
	prices := []int{7,1,5,3,6,4}
	fmt.Println("input: ", prices)
	fmt.Println("output: ", maxProfit(prices))
}
