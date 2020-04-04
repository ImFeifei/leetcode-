package main
import "fmt"

// 2020.4.2 从状态转移方程可以看出dp[i][0 or 1]依赖的只有[i-1][0 and 1]和[i-2][0]的数据
// 故用常量来记录，压缩dp, 可读性不高，仅用于第一种常规dp写出来后优化用
func maxProfit(prices []int) int {
    if len(prices) == 0 || len(prices) == 1 { return 0 }
    profit_pre_0 := 0
    //profit_pre_1 := -prices[0] 最终也可压缩掉
    profit_i_0 := max(profit_pre_0, prices[1] - prices[0])
    profit_i_1 := max(-prices[0], -prices[1])

    for i := 2; i < len(prices); i++ {
        tmp_pre_0 := profit_pre_0
        profit_pre_0 = max(profit_i_0, profit_i_1 + prices[i])
        profit_i_1 = max(profit_i_1, tmp_pre_0 - prices[i])
        profit_pre_0, profit_i_0 = profit_i_0, profit_pre_0
        // profit_pre_1 = max(profit_i_1, tmp_pre_0 - prices[i])
        // profit_pre_0, profit_i_0 = profit_i_0, profit_pre_0
        // profit_pre_1, profit_i_1 = profit_i_1, profit_pre_1
    }
    return max(profit_i_0, profit_i_1)
}

// 2020.4.1 构建二维dp,看起来比较清晰
// func maxProfit(prices []int) int {
//     if len(prices) == 0 || len(prices) == 1 { return 0 }
//     // dp[i][0]表示第i天最终不持股时获得的最大利润
//     // dp[i][1]表示第i天最终持股时获得的最大利润
//     dp := make([][]int, len(prices))
//     for i := 0; i < len(prices); i++ {
//         dp[i] = make([]int, 2)
//     }
//     // 前两天不用考虑冷冻期，拎出来先算掉
//     dp[0][0] = 0  //第0天最终不持股，即没买，利润为0
//     dp[0][1] = -prices[0]   //第0天最终持股，即买入，利润为-prices[0]
//     dp[1][0] = max(dp[0][1]+prices[1], dp[0][0])
//     dp[1][1] = max(dp[0][0]-prices[1], dp[0][1])

//     for i := 2; i < len(prices); i++ {
//         // dp[i][0]可由两种情况得到：【i-1天持有股票，今天卖出】和【i-1天就没有股票】
//         dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
//         // dp[i][1]可由两种情况得到：【i-1天本来就持有股票】和【i-1天没有股票，是今天进行了买入操作】
//         // 后者由于是今天(i天)进行了买入操作，那么i-1天没有股票的原因就一定【不能是】i-1天进行了卖出操作（因为冷冻期）
//         // 即i-1天没有股票的原因一定是因为i-2天就没有股票，因此直接用dp[i-2][0]表示
//         // 可以避免dp[i-1][0]中可能选择的上述违禁操作
//         dp[i][1] = max(dp[i-2][0]-prices[i], dp[i-1][1])
        
//     }
//     return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
// }
func max(x, y int)int {
    if x > y { return x }
    return y
}

func main(){
    prices := []int{1,2,3,0,2}
    fmt.Println("input: ", prices)
    fmt.Println("output: ", maxProfit(prices))
}