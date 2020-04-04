package main
import "fmt"

// 2020.4.1 用BFS广度优先搜索 BFS最典型应用为树的层次遍历，使用队列来完成
// 这里的树类似于求全排列时的解集树，边为每次减去的完全平方数，最终第一个节点值为0的深度就是减去的完全平方数个数
import "math"
func numSquares(n int) int {
    queue := []int{n}
    depth := 0
    var cur int
    visited := make(map[int]int)
    for len(queue) != 0 {
        levelNum := len(queue)
        // 同一层的一起遍历
        for i := 0; i < levelNum; i++ {
            cur, queue = queue[0], queue[1:]
            if cur == 0 {
                return depth
            }
            //这里j为sqrt(cur)更好，import math; 否则用cur/2+1还要加一个判断next是否>=0
            for j := int(math.Sqrt(float64(cur)))/*cur/2+1*/; j >= 1; j-- {  
                next := cur - j * j
                if _, ok := visited[next]; ok { continue } //该完全平方数因子在之前已出现过，故不再重复计数
                visited[next]++
                queue = append(queue, next)                
            }
        }
        depth++
    }
    return depth
}
// 2020.3.31 动态规划
// 每个正整数都可由【若干个完全平方数相加】得到， 因为有1这个完全平方数
// 找小于i的完全平方数j*j
// func numSquares(n int) int {
//     dp := make([]int, n+1)
//     dp[1] = 1
//     for i := 2; i < n+1; i++ {
//         // 最多由i个1组成
//         dp[i] = i  
//         for j := 1; j <= i/2; j++ {
//             if i - j*j >= 0 {
//                 // 1+dp[i-j*j]中1代表j*j这个完全平方数
//                 dp[i] = min(dp[i], 1+dp[i-j*j])
//             } else {
//                 break
//             }
//         } 
//     }
//     return dp[n]
// }
// func min(x,y int)int {
//     if x < y { return x }
//     return y
// }

func main() {
    n := 12
    fmt.Println("input: n = ", n)
    fmt.Println("output: ", numSquares(n))
}
