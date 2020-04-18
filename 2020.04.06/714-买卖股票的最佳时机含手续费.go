package main
import "fmt"

// 2020.4.6 二维dp dp[i][0]表示前i天且第i天不持有股票时的最大利润，dp[i][1]表示前i天且第i天持有股票时的最大利润
func maxProfit(prices []int, fee int) int {
    if len(prices) <= 1 { return 0 }
    dp := make([][]int, len(prices))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, 2)
    }
    dp[0][0], dp[0][1] = 0, -prices[0]
    for i := 1; i < len(prices); i++ {
        // 第i天不持有股票有两种情况：【第i-1天就没有股票】or【第i天卖掉了股票]
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
        // 第i天持有股票有两种情况：【第i-1天就持有股票】or【第i天买入了股票]
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }
    return dp[len(prices)-1][0]
}

// 2020.4.6 优化压缩dp，但总要做出常规dp才方便优化，否则思路不好想
// func maxProfit(prices []int, fee int) int {
//     if len(prices) <= 1 { return 0 }
//     pre_i_0, pre_i_1 := 0, -prices[0]
//     for i := 1; i < len(prices); i++ {
//         pre_i_0, pre_i_1 = max(pre_i_0, pre_i_1+prices[i]-fee), max(pre_i_1, pre_i_0-prices[i])
//     }
//     return pre_i_0
// }

func max(x,y int)int {
    if x > y { return x }
    return y
}

func main() {
    prices := []int{1, 3, 2, 8, 4, 9}
    fee := 2
    fmt.Printf("input: prices = %v, fee = %d\n", prices, fee)
    fmt.Println("output: ", maxProfit(prices, fee))
}
