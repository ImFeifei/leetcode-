package main
import "fmt"

// 要求原地算法：故在原数组上进行记录是否应改变，应改变加10，否则不变
// 则 11 代表这一位应由1变成0，10 代表这一位应由0变成1，其余不变
// 注意，此时计算周围1的个数的时候最后需取个位数，即%10
// 最后遍历board，将十位数与个位数作亦或就可以得到最终值
func gameOfLife(board [][]int)  {
    if len(board) == 0 { return }
    m, n := len(board), len(board[0])
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            count := 0
            if board[i][j] == 0 { //当前是死细胞
                if i-1 >=0 && j-1>=0{
                    count += (board[i-1][j-1]+board[i-1][j]+board[i][j-1])
                } else {
                    if i - 1 >= 0 { count += board[i-1][j] }
                    if j - 1 >= 0 { count += board[i][j-1] }
                }
                if i+1 < m && j+1 < n {
                    count += (board[i][j+1] + board[i+1][j] + board[i+1][j+1])
                } else {
                    if i + 1 < m { count += board[i+1][j] }
                    if j + 1 < n { count += board[i][j+1] }
                }
                if i+1 < m && j-1 >=0 {
                    count += board[i+1][j-1]
                }
                if i-1 >=0 && j + 1 < n {
                    count += board[i-1][j+1]
                }
                if count%10 == 3 { board[i][j] += 10 }
            } else { // 当前为活细胞
                if i-1 >=0 && j-1>=0{
                    count += (board[i-1][j-1]+board[i-1][j]+board[i][j-1])
                } else {
                    if i - 1 >= 0 { count += board[i-1][j] }
                    if j - 1 >= 0 { count += board[i][j-1] }
                }
                if i+1 < m && j+1 < n {
                    count += (board[i][j+1] + board[i+1][j] + board[i+1][j+1])
                } else {
                    if i + 1 < m { count += board[i+1][j] }
                    if j + 1 < n { count += board[i][j+1] }
                }
                if i+1 < m && j-1 >=0 {
                    count += board[i+1][j-1]
                }
                if i-1 >=0 && j + 1 < n {
                    count += board[i-1][j+1]
                }
                if count%10 > 3 || count%10 < 2 { 
                    board[i][j] += 10
                }
            }

        }
    }
    for i:=0; i < m; i++ {
        for j:=0; j < n; j++ {
            board[i][j] = (board[i][j]/10)^(board[i][j]%10) // 十位数与个位数亦或
        }
    }
}


// 2020.4.2 额外开辟了dp二维数组记录原board数组上相应位置是否发生变化，1为变，0不变，最后再遍历一次作异或即得答案
// 其实可以 直接在原board数组上记录，因为board数组目前只用了最低位0/1
// func gameOfLife(board [][]int)  {
//     if len(board) == 0 { return }
//     m, n := len(board), len(board[0])
//     dp := make([][]int, m)
//     for i := 0; i < m; i++ {
//         dp[i] = make([]int, n)
//     }
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             count := 0
//             if board[i][j] == 0 { //当前是死细胞
//                 if i-1 >=0 && j-1>=0{
//                     count += (board[i-1][j-1]+board[i-1][j]+board[i][j-1])
//                 } else {
//                     if i - 1 >= 0 { count += board[i-1][j] }
//                     if j - 1 >= 0 { count += board[i][j-1] }
//                 }
//                 if i+1 < m && j+1 < n {
//                     count += (board[i][j+1] + board[i+1][j] + board[i+1][j+1])
//                 } else {
//                     if i + 1 < m { count += board[i+1][j] }
//                     if j + 1 < n { count += board[i][j+1] }
//                 }
//                 if i+1 < m && j-1 >=0 {
//                     count += board[i+1][j-1]
//                 }
//                 if i-1 >=0 && j + 1 < n {
//                     count += board[i-1][j+1]
//                 }
//                 if count == 3 { dp[i][j] = 1 }
//             } else { // 当前为活细胞
//                 if i-1 >=0 && j-1>=0{
//                     count += (board[i-1][j-1]+board[i-1][j]+board[i][j-1])
//                 } else {
//                     if i - 1 >= 0 { count += board[i-1][j] }
//                     if j - 1 >= 0 { count += board[i][j-1] }
//                 }
//                 if i+1 < m && j+1 < n {
//                     count += (board[i][j+1] + board[i+1][j] + board[i+1][j+1])
//                 } else {
//                     if i + 1 < m { count += board[i+1][j] }
//                     if j + 1 < n { count += board[i][j+1] }
//                 }
//                 if i+1 < m && j-1 >=0 {
//                     count += board[i+1][j-1]
//                 }
//                 if i-1 >=0 && j + 1 < n {
//                     count += board[i-1][j+1]
//                 }
//                 if count != 3 && count != 2 { 
//                     dp[i][j] = 1
//                 }
//             }

//         }
//     }
//     for i:=0; i < m; i++ {
//         for j:=0; j < n; j++ {
//             board[i][j] ^= dp[i][j]
//         }
//     }

// }

func main() {
    board := [][]int{}
    board = append(board, []int{0,1,0})
    board = append(board, []int{0,0,1})
    board = append(board, []int{1,1,1})
    board = append(board, []int{0,0,0})
    fmt.Println("input: ", board)
    gameOfLife(board)
    fmt.Println("output: ", board)
}
