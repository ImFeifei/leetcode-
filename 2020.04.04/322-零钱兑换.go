package main
import "fmt"

func coinChange(coins []int, amount int) int {
    if amount <= 0 { return 0 }
    dp := make([]int, amount+1)
    dp[0] = 0
    for i := 1; i <= amount; i++ {
        dp[i] = amount+1
        for _, coin := range coins {
            if i >= coin {
                dp[i] = min(dp[i-coin]+1, dp[i])
            }   
        }
    }
    if dp[amount] < amount + 1 { return dp[amount] }
    return -1
    
}

func min(x, y int)int {
    if x < y { return x }
    return y
}

func main(){
    coins := []int{2}
    amount := 1
    fmt.Println("input: coins = ",coins, "  amount = ", amount)
    fmt.Println("output: ", coinChange(coins,amount))
}