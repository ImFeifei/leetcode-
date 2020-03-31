package main
import "fmt"

//  solution 2: 传统判断回文的二维dp,实际只用了右上三角
// dp[i][j]表示s[i]...s[j]是否为回文数
func countSubstrings(s string) int {
    if len(s) == 0 || len(s) == 1 { return len(s) }
    dp := make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
    }

    res := 0
    //要先从右侧遍历：确保算[i][j]时，[i+1][j-1]已被计算过
    for j := 0; j < len(s); j++ {
        for i := j; i >= 0; i-- {
            if i == j {
                dp[i][j] = true
                res++
            } else if s[i] == s[j] && (dp[i+1][j-1] == true || j - i < 2) {  // j-i<2代表双点中心,中间没有一个dp值
                dp[i][j] = true
                res++
            }
        }
    }
    return res
}

// solution 1: 中心扩展法：分为单点为中心的奇数串 和 双点中心的偶数串
// func countSubstrings(s string) int {
//     if len(s) == 0 || len(s) == 1 { return len(s) }

//     res := 0
//     for i := 0; i < len(s); i++ {
//         //先判断单点s[i]为中心的奇数串
//         left, right := i, i
//         for left >= 0 && right < len(s) {
//             if s[left] == s[right] {
//                 res++
//                 left--
//                 right++
//             } else {
//                 break
//             }
//         }
//         // 再判断s[i]与s[i+1]为中心的偶数串
//         left, right = i, i+1
//         for left >= 0 && right < len(s) {
//             if s[left] == s[right] {
//                 res++
//                 left--
//                 right++
//             } else {
//                 break
//             }
//         }
//     }
//     return res
// }
func main() {
    s := "abba"
    fmt.Println("input: ", s)
    fmt.Println("output: ", countSubstrings(s))
}