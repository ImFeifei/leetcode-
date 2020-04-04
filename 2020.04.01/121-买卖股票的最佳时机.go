package main
import "fmt"

// solution 2: 只维护两个值：前序天中的【最小价格】 & 在前序天中买卖可得的【最大利润】
func maxProfit(prices []int) int {
    if len(prices) == 0 || len(prices) == 1 { return 0 }
    minPrice, maxProfit := prices[0], 0
    for _, price := range prices {
        maxProfit = max(maxProfit, price - minPrice)
        minPrice = min(minPrice, price)
    }
    return maxProfit
}
func max(x,y int)int {
    if x > y { return x }
    return y
}
func min(x,y int)int {
    if x < y { return x }
    return y
}

// solution 1: 一维dp
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
