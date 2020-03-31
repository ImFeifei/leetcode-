package main
import "fmt"

func maxDistance(grid [][]int) int {
    queue := make([][]int, 0)
    //将陆地点坐标放入队列
    for i, r := range grid {
        for j, c := range r {
            if c == 1 {
                queue = append(queue, []int{i,j})
            }
        }
    }
    //若队列中无点 or 全在队列中
    if len(queue) == 0 || len(queue) == len(grid)*len(grid) {
        return -1
    }
    depth := 0
    land := []int{}
    for (len(queue) != 0) {
        // 当前层depth的陆地点个数levelNum，最后要输出的就是depth-1：经过几次遍历可使最后一个海洋也能被遍历到
        // 其实曼哈顿距离就是网格上从一个格子走几步能到达另一格子，每步只能横竖向走
        depth++
        levelNum := len(queue)  
        //对当前层陆地点进行BFS
        for i := 0; i < levelNum; i++ {
            land, queue = queue[0], queue[1:]
            //对陆地点进行上下左右遍历，若为海洋0则将其置为2并入队作为下一层，表示已遍历过该海洋
            if land[0]-1 >= 0 && grid[land[0]-1][land[1]] == 0 {
                grid[land[0]-1][land[1]] = 2
                queue = append(queue, []int{land[0]-1, land[1]})
            }
            if land[0]+1 < len(grid) && grid[land[0]+1][land[1]] == 0 {
                grid[land[0]+1][land[1]] = 2
                queue = append(queue, []int{land[0]+1, land[1]})
            }
            if land[1]-1 >= 0 && grid[land[0]][land[1]-1] == 0 {
                grid[land[0]][land[1]-1] = 2
                queue = append(queue, []int{land[0], land[1]-1})
            }
            if land[1]+1 < len(grid) && grid[land[0]][land[1]+1] == 0 {
                grid[land[0]][land[1]+1] = 2
                queue = append(queue, []int{land[0], land[1]+1})
            }
        }
    }
    return depth - 1
}

func main() {
    grid := [][]int{{1,0,1},{0,0,0},{1,0,1}}
    // grid = append(grid, []int{1,0,1})
    // grid = append(grid, []int{0,0,0})
    // grid = append(grid, []int{1,0,1})
    fmt.Println("input: ", grid)
    fmt.Println("output: ", maxDistance(grid))
}
