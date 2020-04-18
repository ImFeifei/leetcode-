package main
import "fmt"

//均类似于647-回文子串

// solution 2: 中心扩散法
// func longestPalindrome(s string) string {
//     if len(s) == 0 || len(s) == 1 { return s }

//     max_left, max_right := 0, 0
//     for i := 0; i< len(s); i++ {
//         left, right := i, i
//         for left >= 0 && right < len(s) {
//             if s[left] == s[right] {
//                 if right - left > max_right - max_left {
//                     max_right, max_left = right, left
//                 }
//                 left--
//                 right++
//             } else { break }
//         }

//         left, right = i, i+1
//         for left >= 0 && right < len(s) {
//             if s[left] == s[right] {
//                 if right - left > max_right - max_left {
//                     max_right, max_left = right, left
//                 }
//                 left--
//                 right++
//             } else { break }
//         }       
//     }
//     return s[max_left:max_right+1]
// }

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

// 2020.4.5 回顾2  时间O(n^2) 空间O(1)
func longestPalindrome(s string) string {
    if len(s) < 2 { return s }
    var res string
    for i := 0; i < len(s); i++ {
        left, right := i, i
        for left >= 0 && right < len(s) {
            if s[left] == s[right] {
                if right - left + 1 > len(res) {
                    res = s[left:right+1]
                }
                left--
                right++
            } else {
                break
            }
        }

        left, right = i, i+1
        for left >= 0 && right < len(s) {
            if s[left] == s[right] {
                if right - left + 1 > len(res) {
                    res = s[left:right+1]
                }
                left--
                right++
            } else {
                break
            }
        }
    }
    return res
}

// 2020.4.5 回顾1 时间O(n^2) 空间O(n^2)
// func longestPalindrome(s string) string {
//     if len(s) < 2 { return s }
//     dp := make([][]bool, len(s))
//     for i := 0; i < len(s); i++ {
//         dp[i] = make([]bool, len(s))
//     }
//     max_len, left, right := 0, 0, 0
//     dp[0][0] = true
//     for i := 0; i < len(s); i++ {
//         for j := i; j >= 0; j-- {
//             if j == i {
//                 dp[j][i] = true
//                 //fmt.Printf("dp[%d][%d]:%v\n", j, i, dp[j][i])
//             } else if s[i] == s[j] && (dp[j+1][i-1] == true || i - j < 2) {
//                 dp[j][i] = true
//                 //fmt.Printf("dp[%d][%d]:%v\n", j, i, dp[j][i])
//             } 
//             if dp[j][i] == true && i - j + 1 > max_len {
//                 left, right = j, i
//                 max_len = i - j + 1
//                 //fmt.Printf("left=%d  right=%d  max_len:%v\n", left, right, max_len)
//             }
//         }
//     }
//     return s[left:right+1]
// }


func main() {
    s := "babad"
    fmt.Println("input: ", s)
    fmt.Println("output: ", longestPalindrome(s))
}