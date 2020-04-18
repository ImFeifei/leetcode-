package main
import "fmt"

// 2020.4.15  队列 + BFS从元素0开始遍历所有元素  时间O(m*n)  空间O(m*n) m为行n为列
// 空间主要是队列和visited占的
func updateMatrix(matrix [][]int) [][]int {
    if len(matrix) == 0 { return nil }
    res := make([][]int, len(matrix))
    visited := make([][]bool, len(matrix))
    queue := [][]int{}
    // 设res为结果矩阵，将matrix中为0位置对应为0，同时0元素坐标入队，visited中标记为true；res其余赋值-1
    for i := 0; i < len(matrix); i++ {
        res[i] = make([]int, len(matrix[0]))
        visited[i] = make([]bool, len(matrix[0]))
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                res[i][j] = 0
                visited[i][j] = true
                queue = append(queue, []int{i, j})
            } else {
                res[i][j] = -1
            }
        }
    }
    // BFS模板：队列中元素出队，遍历res中其上下左右元素
    // 上下左右元素若为-1，则赋值为当前出队的元素值+1；若大于0，取最小值
    // 上下左右元素入队，标记visited为true
    pos := []int{}
    directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}  // 方向数组
    for len(queue) != 0 {
        pos, queue = queue[0], queue[1:]
        for _, d := range directions {
            i, j := pos[0] + d[0], pos[1] + d[1]
            if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) && res[i][j] != 0 && visited[i][j] == false{
                if res[i][j] == -1 {
                    res[i][j] = res[pos[0]][pos[1]] + 1
                } else if res[i][j] > 0 {
                    res[i][j] = min(res[pos[0]][pos[1]] + 1, res[i][j])
                }
                queue = append(queue, []int{i, j})
                visited[i][j] = true
            }
        }
    }
    return res
}
func min(x,y int)int {
    if x < y { return x }
    return y
}

func main() {
    matrix := [][]int{{0,0,0},{0,1,0},{1,1,1}}
    fmt.Println("input:", matrix)
    res := updateMatrix(matrix)
    fmt.Println("output:", res)
}
