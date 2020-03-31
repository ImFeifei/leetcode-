package main
import "fmt"

//solution 2: 在原grid数组上计算，空间O(1)
func minPathSum(grid [][]int) int {
    if len(grid) == 0 { return 0 }
    rows, cols := len(grid), len(grid[0])

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if i == 0 && j == 0 {
                continue
            } else if i == 0 && j != 0{
                grid[i][j] += grid[i][j-1]
            } else if j ==0 && i != 0 {
                grid[i][j] += grid[i-1][j]
            } else {
                grid[i][j] += min(grid[i-1][j], grid[i][j-1])
            }
        }
    }
    return grid[rows-1][cols-1]
}

// solution1: 新建二维dp数组
// func minPathSum(grid [][]int) int {
//     if len(grid) == 0 { return 0 }

//     //初始化dp方格
//     rows := len(grid)
//     cols := len(grid[0])
//     dp := make([][]int, rows)
//     for i := 0; i < rows; i++ {
//         dp[i] = make([]int, cols)
//     }
//     // 自左上向右下遍历并计算dp[i][j]：
//     for i := 0; i < rows; i++ {
//         for j := 0; j < cols; j++ {
//             if i == 0 && j == 0{
//                 dp[i][j] = grid[i][j]
//             } else if i == 0 && j != 0{
//                 dp[i][j] = grid[i][j] + dp[i][j-1]
//             } else if j == 0 && i != 0{
//                 dp[i][j] = grid[i][j] + dp[i-1][j]
//             } else {
//                 dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
//             }
            
//         }
//     }
//     return dp[rows-1][cols-1]
// }

func min(x, y int)int {
    if x < y {return x}
    return y
}

func main() {
    grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
    fmt.Println("input: ", grid)
    fmt.Println("output: ", minPathSum(grid))
}