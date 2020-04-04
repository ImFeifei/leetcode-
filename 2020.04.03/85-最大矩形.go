package main
import "fmt"

func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 { return 0 }
    m, n := len(matrix), len(matrix[0])
    // 将matrix赋值给nums对应int
    nums := make([][]int, m)
    for i:=0; i < m; i++ {
        nums[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                nums[i][j] = 1
            }
        }
    }
    max_area := 0
    for i := 0; i < m; i++ {
        // 将每行的元素压缩为每列的连续1的个数 转换成84题柱状图
        for j := 0; j < n; j++ {
            if i != 0 && nums[i-1][j] != 0 && nums[i][j] == 1 {
                nums[i][j] += nums[i-1][j]
            }
        }
        max_area = max(max_area, largestRectangleArea(nums[i]))
    }
    return max_area
}
// 84题函数
func largestRectangleArea(heights []int) int {
    if len(heights) == 0 { return 0 }
    if len(heights) == 1 { return heights[0] }

    stack := []int{-1}
    heights = append(heights, -1)
    cur, max_area := 0, 0
    for i := 0; i < len(heights); i++ {
        if len(stack) == 1 || heights[i] >= heights[stack[len(stack)-1]] {
            stack = append(stack, i)
        } else {
            for len(stack) != 1 && heights[stack[len(stack)-1]] > heights[i] {
                cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
                max_area = max(max_area, heights[cur] * (i - stack[len(stack)-1] - 1))
            }
            stack = append(stack, i)
        }
    }
    return max_area
}
func max(x,y int) int {
    if x > y { return x }
    return y
}
func main() {
    matrix := make([][]byte, 4)

    matrix[0] = []byte{'1','0','1','0','0'}
    matrix[1] = []byte{'1','0','1','1','1'}
    matrix[2] = []byte{'1','1','1','1','1'}
    matrix[3] = []byte{'1','0','0','1','0'}
    // matrix := make([][]byte, 2)
    // matrix[0] = []byte{'1','1','1','0','0'}
    // matrix[1] = []byte{'0','0','0','0','1'}
    
    fmt.Println("input: ", matrix)
    fmt.Println("output: ", maximalRectangle(matrix))
}

