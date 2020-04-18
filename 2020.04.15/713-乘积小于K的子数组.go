package main
import "fmt"

// 2020.4.15 双指针（滑动窗口）
// 如果一个子串的乘积小于k，那么他的每个子集都小于k，循环中的每一步是在计算以nums[right]结尾的乘积小于k的连续子数组个数
func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1 { return 0 }  // 因为要严格小于k，小于1也不可能
    count := 0
    left, product := 0, 1
    for right := 0; right < len(nums); right++ {
        product *= nums[right]
        for left <= right && product >= k {
            product /= nums[left]
            left++
        }
        count += right - left + 1   // 以nums[right]结尾的乘积小于k的连续子数组个数
    }
    return count
}
// 2020.4.15 暴力遍历+剪枝 时间O(n^2) 空间O(1) 超时
// func numSubarrayProductLessThanK(nums []int, k int) int {
//     if k <= 0 { return 0 }
//     res := 0
//     for i := 0; i < len(nums); i++ {
//         product := 1
//         for j := i; j < len(nums); j++ {
//             product *= nums[j]
//             if product < k {
//                 res++
//             } else {  // 剪枝
//                 break
//             }
//         }
//     }
//     return res
// }

func main() {
    nums := []int{10,5,2,6}
    k := 100
    fmt.Printf("input: nums = %v, k = %d\n", nums, k)
    fmt.Println("output: ", numSubarrayProductLessThanK(nums, k))
}
