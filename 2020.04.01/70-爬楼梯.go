package main
import "fmt"

// dp[i]只依赖dp[i-1] & dp[i-2]，可压缩
// 实际上就是斐波那契数列
func climbStairs(n int) int {
    if n == 1 { return 1 }
    if n == 2 { return 2 }
    a, b := 1, 2
    for i := 3; i < n+1; i++ {
        c := a + b
        a = b
        b = c
    }
    return b
}

// solution 1: 传统dp
// func climbStairs(n int) int {
//     if n == 1 { return 1 }
//     if n == 2 { return 2 }
    
//     // dp[i]为爬到阶梯i共有的爬法种数，最终返回dp[n]
//     // 状态转移：爬到台阶i的方式有【从i-1爬一阶】和【从i-2一下爬两阶】
//     dp := make([]int, n+1)
//     dp[1] = 1
//     dp[2] = 2
//     for i := 3; i < n+1; i++ {
//         dp[i] = dp[i-1] + dp[i-2]
//     }
//     return dp[n]
// }

func main(){
	n := 5
	fmt.Println("input: ", n)
	fmt.Println("output: ", climbStairs(n))
}
