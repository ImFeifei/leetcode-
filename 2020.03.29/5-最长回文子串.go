package main
import "fmt"

//均类似于647-回文子串

// solution 2: 中心扩散法
func longestPalindrome(s string) string {
    if len(s) == 0 || len(s) == 1 { return s }

    max_left, max_right := 0, 0
    for i := 0; i< len(s); i++ {
        left, right := i, i
        for left >= 0 && right < len(s) {
            if s[left] == s[right] {
                if right - left > max_right - max_left {
                    max_right, max_left = right, left
                }
                left--
                right++
            } else { break }
        }

        left, right = i, i+1
        for left >= 0 && right < len(s) {
            if s[left] == s[right] {
                if right - left > max_right - max_left {
                    max_right, max_left = right, left
                }
                left--
                right++
            } else { break }
        }       
    }
    return s[max_left:max_right+1]
}

// solution 1: 传统二维dp找回文子串
// func longestPalindrome(s string) string {
//     if len(s) == 0 || len(s) == 1 { return s }

//     dp := make([][]bool, len(s))
//     for i := 0; i < len(s); i++ {
//         dp[i] = make([]bool, len(s))
//     }

//     var res string
//     for j := 0; j < len(s); j++ {
//         for i := j; i >= 0; i-- {
//             if i == j {
//                 dp[i][j] = true
//             } else if s[i] == s[j] && (dp[i+1][j-1] || j-i < 2) {
//                 dp[i][j] = true
//             }
//             if dp[i][j] == true && len(res) < j-i+1 {
//                 res = s[i:j+1]
//             }
//         }
//     }
//     return res
// }

func main() {
    s := "babad"
    fmt.Println("input: ", s)
    fmt.Println("output: ", longestPalindrome(s))
}