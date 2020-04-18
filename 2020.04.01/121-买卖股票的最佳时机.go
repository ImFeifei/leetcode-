package main
import "fmt"


// 2020.4.6 方法3 dp[i]表示前i天卖出可得最大利润 最终返回dp[len(prices)-1]
// func maxProfit(prices []int) int {
//     if len(prices) <= 1 { return 0 }
//     dp := make([]int, len(prices))
//     minPrice := prices[0]
//     for i := 1; i < len(prices); i++ {
//         minPrice = min(minPrice, prices[i])
//         dp[i] = max(dp[i-1], prices[i] - minPrice)
//     }
//     return dp[len(prices)-1]
// }
// func max(x,y int)int {
//     if x > y { return x }
//     return y
// }
// func min(x,y int)int {
//     if x < y { return x }
//     return y
// }

// 2020.4.6 方法2 dp[i]表示第i天卖出可得最大利润 最终返回max(dp)
// func maxProfit(prices []int) int {
//     if len(prices) <= 1 { return 0 }
//     dp := make([]int, len(prices))
//     for i := 1; i < len(prices); i++{
//         dp[i] = max(dp[i], dp[i-1] + prices[i] - prices[i-1])
//     }
//     return maxInSlice(dp)
// }
// func max(x,y int)int {
//     if x < y { return y }
//     return x
// }
// func maxInSlice(nums []int) int{
//     res := nums[0]
//     for _, num := range nums {
//         if num > res { res = num }
//     }
//     return res
// }

// 2020.4.6 方法一(推荐,比较容易想) 只维护两个值：前序天中的【最小价格】 & 在前序天中买卖可得的【最大利润】
func maxProfit(prices []int) int {
    if len(prices) <= 1 { return 0 }
    minPrice := prices[0]
    maxProfit := 0
    for _, price := range prices {
        minPrice = min(minPrice, price)
        maxProfit = max(maxProfit, price - minPrice)
    }
    return maxProfit
}
func max(x,y int)int {
    if x < y { return y }
    return x
}
func min(x,y int)int {
    if x < y { return x }
    return y
}


// 2020.4.1 solution 2: 只维护两个值：前序天中的【最小价格】 & 在前序天中买卖可得的【最大利润】
// func maxProfit(prices []int) int {
//     if len(prices) == 0 || len(prices) == 1 { return 0 }
//     minPrice, maxProfit := prices[0], 0
//     for _, price := range prices {
//         maxProfit = max(maxProfit, price - minPrice)
//         minPrice = min(minPrice, price)
//     }
//     return maxProfit
// }
// func max(x,y int)int {
//     if x > y { return x }
//     return y
// }
// func min(x,y int)int {
//     if x < y { return x }
//     return y
// }

// 2020.4.1 solution 1: 一维dp
// func maxProfit(prices []int) int {
//     if len(prices) == 0 || len(prices) == 1 { return 0 }
//     // dp[i]为第i天卖出能得到的最大利润， 故最后返回dp[i]中最大值
//     dp := make([]int, len(prices))
//     min := prices[0]
//     for i, price := range prices {
//         if price > min {
//             dp[i] = price - min
//         } else {
//             min = price
//         }
//     }
//     return max(dp)
// }
// func max(nums []int) int {
//     maxNum := nums[0]
//     for _, num := range nums {
//         if num > maxNum {
//             maxNum = num
//         }
//     }
//     return maxNum
// }

func main() {
    prices := []int{7,1,5,3,6,4}
    fmt.Println("input: ", prices)
    fmt.Println("output: ", maxProfit(prices))
}
