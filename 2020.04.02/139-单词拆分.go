package main
import "fmt"

// 一维dp dp[i]表示前i个元素是否能拆分成词典中的单词
//dp[i]：表示前i个元素是否能拆分成词典中的单词，故最终返回dp[len(s)]。
// 对每一个起始元素i遍历截止元素j（j范围[i+1, len(s)]）
// 若前i个元素已经可以成功拆分且i到j也是一个可选单词，则前j个元素也可以成功拆分，即d[j]置为true。

func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i := 0; i < len(s); i++ {
        for j := i+1; j < len(s)+1; j++ {
            for _, word := range wordDict {
                if s[i:j] == word && dp[i] == true {
                    dp[j] = true
                }
            }
        }
    }
    return dp[len(s)]
}

// 2020.4.2 自己一遍过的思路
// dp[i][j] 表示s[i:j+1]是否可拆分为字典中的单词，则最后返回dp[0][len(s)-1]
// 状态转移：对每个起始元素i遍历word，若s[i:i+len(word)]==word，则dp[i][i+len(word)-1]置为true
// 同时，倒回寻找以i-1为结尾的true组合dp[j][i-1]，将其拼接成的组合dp[j][i+len(word)-1]也置为true
// func wordBreak(s string, wordDict []string) bool {
//     dp := make([][]bool, len(s))
//     for i := 0; i< len(s); i++ {
//         dp[i] = make([]bool, len(s))
//     }
//     for i := 0; i < len(s); i++ {
//         for _, word := range wordDict {
//             if len(word)+i-1 < len(s) {
//                 if s[i:i+len(word)] == word {
//                     dp[i][i+len(word)-1] = true
//                     for j := i-1; j >= 0; j-- {
//                         if dp[j][i-1] == true {
//                             dp[j][i+len(word)-1] = true
//                         }
//                     }
//                 }    
//             }
//         }
//     }
//     return dp[0][len(s)-1]
// }

func main(){
    s := "leetcode"
    wordDict := []string{"leet", "code"}
    fmt.Printf("input: s = %s, wordDict = %v\n", s, wordDict)
    fmt.Println("output: ", wordBreak(s, wordDict))
}
