package main
import "fmt"

// 2020.4.4 dp
func findTargetSumWays(nums []int, S int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if S > sum || S < -sum { return 0 }
    // dp[i][j]表示前i个元素构成j的种数，最终返回dp[len(nums)][S]
    dp := make([][]int, len(nums))
    // S会存在负值，j也会,范围在 [-sum(nums),sum(nums)]
    for i := 0; i < len(nums); i++ {
        dp[i] = make([]int, 2*sum+1)
    }
    // 这里忽略了当第一个数为0时对于nums[0]=0加减都可以得到
    if nums[0] != 0 {
        dp[0][sum+nums[0]], dp[0][sum-nums[0]] = 1, 1
    } else {
        dp[0][sum+nums[0]], dp[0][sum-nums[0]] = 2, 2
    }
    
    for i := 1; i < len(nums); i++ {
        for j := -sum; j < sum+1; j++ {
            // 当求最小负值和时，已经无法从上一步相加得来
            if j-nums[i]+sum >= 0 { 
                dp[i][j+sum] += dp[i-1][j-nums[i]+sum]
            }
            if j+nums[i]+sum <= 2*sum {// 当求最大正值和时，已经无法从上一步相减得来
                dp[i][j+sum] += dp[i-1][j+nums[i]+sum] 
            }    
        }    
    }
    return dp[len(nums)-1][S+sum]
}

// func findTargetSumWays(nums []int, S int) int {
//     result := 0
//     backtrack(0, nums, S, &result)
//     return result
// }
// // 2020.4.3 回溯遍历所有可能 时间O(2^n) 惊了，回溯竟然过了
// func backtrack(ans int, nums []int, S int, result *int) {
//     if len(nums) == 0 {
//         if ans == S { *result++ }
//         return
//     }
//     cur := nums[0]
//     ans += cur
//     backtrack(ans, nums[1:], S, result)
//     ans -= cur
//     ans -= cur
//     backtrack(ans, nums[1:], S, result)
//     ans += cur
// }

func main() {
    nums := []int{0,0,0,0,0,0,0,0,1}
    S := 1
    fmt.Printf("input: nums = %v,  S = %d\n", nums, S)
    fmt.Println("output: ", findTargetSumWays(nums, S))
}
