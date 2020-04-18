package main
import "fmt"

// 2020.4.17 方法2 双端单调递减队列 时间O(n) 空间O(k)
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 1 || k == 1 { return nums }
    queue, res := []int{0}, []int{}
    for i := 1; i < len(nums); i++ {
        // 保证队列递减，这里平均下来是每个元素O(1)
        for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
            queue = queue[:len(queue)-1]
        }
        queue = append(queue, i)
        // 这时需要输出窗口中的最大值，即队首元素
        if i >= k - 1 {
            res = append(res, nums[queue[0]])
        }
        // 这时说明以队首元素为最大值的窗口已遍历完，窗口需要继续右移一位，队首出队
        if i == queue[0] + k - 1{
            queue = queue[1:]
        }
    }
    return res
}

// 2020.4.17 方法1 记录上个窗口最大值索引，窗口右移时看是否需要更新 时间O(n*k) 空间O(1)
// func maxSlidingWindow(nums []int, k int) []int {
//     if len(nums) == 1 || k == 1 { return nums }
//     left := 0
//     right := left + k - 1
//     maxIndex := max(nums,left,right)
//     res := []int{nums[maxIndex]}
//     for right < len(nums) {
//         left++
//         right++
//         if right < len(nums) && nums[right] >= nums[maxIndex] {
//             maxIndex = right
//         } else if right < len(nums) && maxIndex < left {
//             maxIndex = max(nums, left, right)
//             fmt.Println("maxIndex: ", maxIndex)
//         } else if right >= len(nums){
//             break
//         }
//         res = append(res, nums[maxIndex])
//     }
//     return res
// }
// func max(nums []int, start, end int) int {
//     resIndex := start
//     for i := start; i <= end; i++ {
//         if nums[i] >= nums[resIndex] {
//             resIndex = i
//         }
//     }
//     return resIndex
// }

func main() {
    nums := []int{1,3,-1,-3,5,3,6,7}
    k := 3
    fmt.Printf("input: nums = %v, k = %d\n", nums, k)
    fmt.Println("output: ", maxSlidingWindow(nums, k))

}