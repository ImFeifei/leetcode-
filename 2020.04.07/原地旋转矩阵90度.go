package main
import "fmt"

// 2020.4.7 题解思路：先对角线翻转，再以中心列左右翻转 即为旋转90度
func rotate(matrix [][]int)  {
    if len(matrix) == 0 { return }
    n := len(matrix)
    // 对角线翻转： 右上与左下翻转
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    // 以中心列为轴左右翻转
    for i := 0; i < n; i++ {
        for j := 0; j < n / 2; j++ {
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }
    }
}

// 2020.4.7 自己做的方法 原地旋转n*n数组 找规律 每行从对角线元素开始一圈一圈换
// func rotate(matrix [][]int)  {
//     if len(matrix) == 0 { return }
//     n := len(matrix)
//     for i := 0; i < n; i++ {
//         for j := i; j < n - i - 1; j++ {
//             tmp := matrix[i][j]
//             u, v := i, j
//             for k := 0; k < 3; k++ {  // 四个数轮换
//                 matrix[u][v] = matrix[n-v-1][u]
//                 u, v = n-v-1, u
//             }
//             matrix[u][v] = tmp 
//         }   
//     }
// }

func main() {
    matrix := [][]int{}
    matrix = append(matrix, []int{1,2,3,4})
    matrix = append(matrix, []int{5,6,7,8})
    matrix = append(matrix, []int{9,10,11,12})
    matrix = append(matrix, []int{13,14,15,16})
    fmt.Println("input: ", matrix)
    rotate(matrix)
    fmt.Println("output: ", matrix)
}
