package main
import "fmt"
// solution 2   时间O(mn)  空间(mn)
func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 { return 0 }
    // 根据matrix构造一个 【每行数字代表：截止到该处时该列连续1的个数】的int二维矩阵
    nums := make([][]int, len(matrix))
    for i := 0; i < len(matrix); i++ {
        nums[i] = make([]int, len(matrix[0]))
        for j := 0; j < len(matrix[0]); j++ {
            if i == 0 {
                if matrix[i][j] == '1' {
                    nums[i][j] = 1
                }
            } else {
                if matrix[i][j] == '1' {
                    nums[i][j] += (1+nums[i-1][j])
                }
            }
        }
    }

    maxArea := 0
    for i := 0; i < len(nums); i++ {
        maxArea = max(maxArea, maximalSquareInALine(nums[i]))
    }
    return maxArea
}
// 方法来自84题 巧用栈 柱状图中的最大正方形
// 每次计算弹出栈的元素为高是否可构成正方形
func maximalSquareInALine(heights []int) int {
    if len(heights) == 0 { return 0 }
    stack := []int{-1}
    heights = append(heights, -1)
    cur, maxArea := 0, 0
    for i, height := range heights {
        if i == 0 {
            stack = append(stack, i)
        } else {
            if len(stack) > 1 && height >= heights[stack[len(stack)-1]] {
                stack = append(stack, i)
            } else {
                for len(stack) > 1 && height < heights[stack[len(stack)-1]] {
                    cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
                    if i - stack[len(stack)-1] - 1 >= heights[cur] {
                        maxArea = max(maxArea, heights[cur]*heights[cur])
                    }
                }
                stack = append(stack, i)
            }
        }
    }
    return maxArea
}

// 2020.4.4 solution 1: 按照自己的思路压缩列的连续1，再向左延伸找是否存在min_height可使当前元素为右下角构成square
// 时间O(m^2 * n)  空间O(m * n)
// func maximalSquare(matrix [][]byte) int {
//     if len(matrix) == 0 { return 0 }
//     // 根据matrix构造一个 【每行数字代表：截止到该处时该列连续1的个数】的int二维矩阵
//     nums := make([][]int, len(matrix))
//     for i := 0; i < len(matrix); i++ {
//         nums[i] = make([]int, len(matrix[0]))
//         for j := 0; j < len(matrix[0]); j++ {
//             if i == 0 {
//                 if matrix[i][j] == '1' {
//                     nums[i][j] = 1
//                 }
//             } else {
//                 if matrix[i][j] == '1' {
//                     nums[i][j] += (1+nums[i-1][j])
//                 }
//             }
//         }
//     }
//     maxArea := 0
//     flag := false
//     min_height := 0
//     for i := 0; i < len(nums); i++ {
//         for j := 0; j < len(nums[0]); j++ {
//             if nums[i][j] > 1 {  
//                 min_height = nums[i][j]
//                 k := 1
//                 for k < min_height { // 向左延伸寻找是否存在min_hight可构成正方形
//                     if j - k < 0 { 
//                         flag = false
//                         break 
//                     }
//                     if nums[i][j-k] == 0 {
//                         flag = false
//                         break 
//                     } 
//                     if nums[i][j-k] < min_height {
//                         min_height = nums[i][j-k] 
//                     } 
//                     flag = true
//                     k++
//                 }
//                 if flag == true {
//                     maxArea = max(maxArea, min_height*min_height)
//                 }
//             } else if nums[i][j] == 1 {
//                 maxArea = max(maxArea, 1)
//             }
//         }
//     }
//     return maxArea
// }
func max(x,y int)int {
    if x < y { return y }
    return x
}


func main() {
    matrix := make([][]byte, 5)

    matrix[0] = []byte{'0','0','0','1'}
    matrix[1] = []byte{'1','1','0','1'}
    matrix[2] = []byte{'1','1','1','1'}
    matrix[3] = []byte{'0','1','1','1'}
    matrix[4] = []byte{'0','1','1','1'}

    // matrix[0] = []byte{'1','0','1','0','0'}
    // matrix[1] = []byte{'1','0','1','1','1'}
    // matrix[2] = []byte{'1','1','1','1','1'}
    // matrix[3] = []byte{'1','0','0','1','0'}

    fmt.Println("input: ", matrix)
    fmt.Println("output: ", maximalSquare(matrix))
}

