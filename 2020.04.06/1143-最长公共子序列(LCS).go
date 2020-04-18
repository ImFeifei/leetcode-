package main
import "fmt"

// 上一解法dp的base case直接从第一个字符开始算起 本解法在起始处加入空字符串 可省略i,j为0时的一些判断
func longestCommonSubsequence(text1 string, text2 string) int {
    if text1 == text2 { return len(text1) }

    dp := make([][]int, len(text1) + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(text2) + 1)
    }

    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[i]); j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], max(dp[i][j-1], dp[i-1][j-1]))
            }
        }
    }
    return dp[len(text1)][len(text2)]
}

// 2020.4.6 类似编辑距离，每日一题刚做过
// dp[i][j] 表示【必须以text1[i]结尾的text1的子序列】与【必须以text2[j]结尾的text2的子序列】的最长公共子序列
// func longestCommonSubsequence(text1 string, text2 string) int {
//     if text1 == text2 { return len(text1) }

//     dp := make([][]int, len(text1))
//     for i := 0; i < len(dp); i++ {
//         dp[i] = make([]int, len(text2))
//     }

//     for i := 0; i < len(dp); i++ {
//         for j := 0; j < len(dp[i]); j++ {
//             if i == 0 {
//                 for k := j; k >= 0; k-- {
//                     if text2[k] == text1[i]{
//                         dp[i][j] = 1
//                     }
//                 }
//                 continue
//             }
//             if j == 0 {
//                 for k := i; k >= 0; k-- {
//                     if text1[k] == text2[j]{
//                         dp[i][j] = 1
//                     }
//                 }
//                 continue
//             }
//             if text1[i] == text2[j] {
//                 dp[i][j] = dp[i-1][j-1] + 1
//             } else {
//                 dp[i][j] = max(dp[i-1][j-1], max(dp[i-1][j], dp[i][j-1]))
//             }
//         }
//     }
//     return dp[len(text1)-1][len(text2)-1]
// }
func max(x,y int)int {
    if x > y { return x }
    return y
}

func main() {
    text1 := "abcde"
    text2 := "ace"
    fmt.Printf("input: text1 = %s, text2 = %s\n", text1, text2)
    fmt.Println("output: ", longestCommonSubsequence(text1, text2))
}
