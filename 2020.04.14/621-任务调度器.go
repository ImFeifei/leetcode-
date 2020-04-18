package main
import (
    "fmt"
    "sort"
)

// 最短时间取决于【出现次数最多的任务】
// 对于出现次数最多且次数为mostCount的任务来说，至少需要(mostCount - 1) * (n + 1) + 1的时间
// 再把剩下任务填进每个(n + 1)时间片中，同时考虑出现最多次数的任务不止一个，可能有mostCountTasks个
// 因此最终至少需要(mostCount - 1) * (n + 1) + mostCountTasks的时间
// 最后再考虑当【任务种类远远大于冷却时间n】或【mostCountTasks接近n、占用了很多个(n+1)时间片】时
// 每个(n+1)时间片填满了后面都还有很多不同种类，这时最短时间就等于任务总数,因此最后比较若时间小于任务总数则返回任务总数
func leastInterval(tasks []byte, n int) int {
    if n == 0 { return len(tasks) }
    if len(tasks) == 1 { return 1 }
    hm := make([]int, 26)
    for _, task := range tasks {
        hm[task - 'A']++
    }
    sort.Slice(hm, func(i, j int)bool{ return hm[i] > hm[j] })
    mostCount := hm[0]
    mostCountTasks := 0
    for _, cnt := range hm {
        if cnt == mostCount {
            mostCountTasks++
        }
    }
    res := (mostCount - 1) * (n + 1) + mostCountTasks
    if res < len(tasks) { res = len(tasks) }
    return res
}

func main() {
    tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
    n := 2
    fmt.Printf("input: tasks = %v, n = %d\n", tasks, n)
    fmt.Println("output: ", leastInterval(tasks, n))
}
