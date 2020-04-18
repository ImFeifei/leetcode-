package main
import "fmt"

// 2020.4.17 方法3: 维护一个最远可达距离，遍历数组更新最远可达距离
// 若遍历的当前索引大于最远距离则说明已经超过最远可达距离，即当前索引已不可达
// 若最远可达距离大于等于终点索引，说明终点可达，立即返回
func canJump(nums []int) bool {
    if len(nums) <= 1 { return true }
    maxDest := 0
    for i := 0; i < len(nums); i++ {
        if i > maxDest { break }
        maxDest = max(maxDest, nums[i] + i)
        if maxDest >= len(nums)-1 { return true }
    }
    return false
}
func max(x,y int)int{
    if x > y { return x }
    return y
}

// 2020.4.17 方法2: 用了一个存储可达点索引的队列
// 1. 当队列为空时表示所有可达点都已遍历完却还没有到终点索引
// 2. 每次更新可达点队列后，检查终点索引是否已可达
// func canJump(nums []int) bool {
//     if len(nums) <= 1 { return true }
//     queue := []int{0}
//     cur, preEnd := 0, 0
//     for len(queue) != 0 {
//         cur, queue = queue[0], queue[1:]
//         if nums[cur] + cur > preEnd {
//             for j := preEnd + 1; j < len(nums) && j <= cur + nums[cur]; j++ {
//                 queue = append(queue, j)
//             }
//             preEnd = nums[cur] + cur
//         } 
//         if len(queue) == 0 { return false }
//         if queue[len(queue)-1] == len(nums)-1 { return true }
//     }
//     return false
// }

// 2020.4.17 方法1 递归
// func canJump(nums []int) bool {
//     if len(nums) <= 1 { return true }
    
//     return helper(nums, len(nums)-1)
// }
// // 最后一位是否可达：前一位的值大于等于间隔时，问题就等价于前一位是否可达
// // 递归超时
// func helper(nums []int, dest int) bool {
//     if len(nums) == 1 { return true }
//     res := false
//     for i := len(nums) - 2; i >= 0; i-- {
//         if nums[i] >= dest - i {
//             res = res || helper(nums[:i+1], i)
//             if res == true {
//                 return res
//             } 
//         }
//     }
//     return res
// }

func main() {
    nums := []int{3,2,1,0,4}. // 2,3,1,1,4   0,2   0,2,3
    fmt.Println("input: ", nums)
    fmt.Println("output: ", canJump(nums))

}