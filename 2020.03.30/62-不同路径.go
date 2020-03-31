package main
import "fmt"

// 一维DP
func uniquePaths(m int, n int) int {
    if n == 1 || m == 1 {
        return 1
    }
    // 初始化dp为第一行
    dp := make([]int, m)
    for i := 0; i < m; i++ {
        dp[i] = 1
    }

    for j := 1; j < n; j++ {
        // 对计算第i行dp：每个格子值等于左格子值加上格子值，上格子值即为待更新的当前值
        for i := 1; i < m; i++ {
            dp[i] += dp[i-1]
        }
    }
    
    return dp[m-1]
}
// 二维DP
// func uniquePaths(m int, n int) int {
//     if n == 1 || m == 1{
//         return 1
//     }
//     dp := make([][]int, n)
//     for i := 0; i < n; i++ {
//         dp[i] = make([]int, m)
//     }

//     for i := 0; i < n; i++ {
//         for j := 0; j < m; j++ {
//             if i == 0 || j == 0 {
//                 dp[i][j] = 1
//             } else {
//                 dp[i][j] = dp[i-1][j] + dp[i][j-1]
//             }
//         }
//     }
//     return dp[n-1][m-1]
// }
func main() {
    m, n := 7, 3
    fmt.Printf("input: m = %d, n = %d\n", m, n)
    fmt.Println("output: ", uniquePaths(m,n))
}
