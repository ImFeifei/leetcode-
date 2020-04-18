package main
import (
    "fmt"
    "sort"
)

// 2020.4.16 先按照左端点排序, 再看相邻的区间是否可以合并
// 时间O(nlogn) 空间(logn) 都用在排序上，实际上排序后只线性扫描，除解数组外也没有使用额外空间
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 { return intervals }
    sort.Slice(intervals, func(i, j int)bool{ return intervals[i][0] < intervals[j][0] })
    res := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        left, right := res[len(res)-1][0], res[len(res)-1][1]
        if intervals[i][0] >= left && intervals[i][0] <= right {
            res[len(res)-1][0] = min(left, intervals[i][0])
            res[len(res)-1][1] = max(right, intervals[i][1])
        } else {
            res = append(res, intervals[i])
        }
    }
    return res
}
func min(x,y int)int {
    if x < y { return x }
    return y
}
func max(x,y int)int {
    if x > y { return x }
    return y
}

func main() {
    intervals := [][]int{{1,3}, {15,18}, {8,10}, {2,6}}
    fmt.Printf("input: intervals = %v\n", intervals)
    fmt.Println("output: ", merge(intervals))
}
