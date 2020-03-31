package main
import "fmt"

func numTrees(n int) int {
    //dp[i]表示给定1 ... i节点可组成的bst个数
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            // j为根,i-j个大于j的点组成的不同右子树个数 * j-1个小于j的点组成的不同左子树个数
            dp[i] += dp[i-j] * dp[j-1]
        }
    }
    return dp[n]
}

func main() {
    n := 3
    fmt.Println("input: ", n)
    fmt.Println("output: ", numTrees(n))
}
