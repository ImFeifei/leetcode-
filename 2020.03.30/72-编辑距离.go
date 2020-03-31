package main
import "fmt"

func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    // dp[i][j]表示word1的前i个字母变成word2的前j个字母所需的最小步数
    // 注：dp[0][0]表示由空串到空串，dp[0][2]表示由空串到word2[:2],dp[2][0]表示由word1[:2]到空串 所需的最小步数
    // 故最终返回的就是dp[m][n]
    dp := make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        dp[i] = make([]int, n+1)
    }

    for i := 0; i < m+1; i++ {
        for j := 0; j < n+1; j++ {
            if i == 0 {
                dp[i][j] = j
                continue
            }
            if j == 0 {
                dp[i][j] = i
                continue
            }
            //注意dp和word对应的下标关系
            if word1[i-1] == word2[j-1] {
                // dp[i-1][j]是dp[i][j]删除一个字母可得
                // dp[i][j-1]是增加一个字母可得dp[i][j]
                // dp[i-1][j-1]是替换最后一个字母可得dp[i][j]
                // 两个单词的末尾字母相同，则不需要替换操作 于是dp[i-1][j-1]-1
                dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]-1))
            } else {
                // 两个单词的末尾字母不相同，则需要加上替换操作
                dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
            }
        }
    }
    return dp[m][n]
}
func min(x, y int)int{
    if x < y { return x }
    return y
}

func main() {
    word1 := "horse"
    word2 := "ros"
    fmt.Printf("input: word1 = %s, word2 = %s\n", word1, word2)
    fmt.Println("output: ", minDistance(word1, word2))
}
