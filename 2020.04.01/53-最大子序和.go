package main
import "fmt"

// 2020.4.2 分治算法 时间O(nlogn)
// 二分后，子序被完整且不相交地包括在【左半部分及其子序】,【右半部分及其子序】,【横跨中心两元素的子序】这三种情况中
// 求这三者最大值即可
func maxSubArray(nums []int) int {
    if len(nums) == 1 { return nums[0] }

    mid :=  len(nums) / 2
    // 左半部分及其子序所能取到的最大子序
    max_in_left := maxSubArray(nums[:mid])

    // 右半部分及其子序所能取到的最大子序
    max_in_right := maxSubArray(nums[mid:])

    // 横跨中心两元素的所能取到的最大子序，必然包含nums[mid-1]和nums[mid]
    // 左半子序中找结尾元素为nums[mid-1]的子序的最大值
    tmp_max_l, cur := nums[mid-1], nums[mid-1]
    for i := mid-2; i >= 0; i-- {
        cur += nums[i]
        tmp_max_l = maxFromTwoNum(tmp_max_l, cur)
    }
    // 右半子序中找起始元素为nums[mid]的子序的最大值
    tmp_max_r, cur := nums[mid], nums[mid]
    for i := mid+1; i < len(nums); i++ {
        cur += nums[i]
        tmp_max_r = maxFromTwoNum(tmp_max_r, cur)
    }
    max_cross_mid := tmp_max_l + tmp_max_r

    return maxFromTwoNum(max_in_left, maxFromTwoNum(max_in_right, max_cross_mid))
}
// 2020.4.1 solution 2: 可以看到dp[i]仅依赖于前一个dp[i-1]，因此可压缩dp至常数空间
// 时间O(n)，空间O(1)
// func maxSubArray(nums []int) int {
//     if len(nums) == 0 { return 0 }
//     if len(nums) == 1 { return nums[0] }
//     max := nums[0]
//     pre := nums[0] 
//     for i := 1; i < len(nums); i++ {
//         cur := maxFromTwoNum(nums[i], pre + nums[i])
//         max = maxFromTwoNum(max, cur)
//         pre = cur
//     }
//     return max
// }

// 2020.4.1 solution1: 一维dp
// 时间O(n)，空间O(n)
// func maxSubArray(nums []int) int {
//     if len(nums) == 0 { return 0 } // 这里工程上应wrapError 输入违法
//     if len(nums) == 1 { return nums[0] }

//     // dp[i]：以第i个元素结尾的连续子数组的最大和，故返回max(dp)
//     // 也可有另一种定义方式：
//     // dp[i]为前i个元素连续子数组的最大和，这样就返回dp[len(nums)-1],状态转移方程再加一个dp[i]取三者最大
//     dp := make([]int, len(nums))
//     dp[0] = nums[0]
//     for i := 1; i < len(nums); i++{
//         dp[i] = maxFromTwoNum(nums[i], dp[i-1]+nums[i])
//     }
//     return maxFromSlice(dp)
// }
func maxFromTwoNum(x, y int)int {
    if x > y { return x }
    return y
}
// func maxFromSlice(nums []int) int {
//     if len(nums) == 0 { return -1 }
//     max := nums[0]
//     for _, num := range nums {
//         if num > max {
//             max = num
//         }
//     }
//     return max
// }

func main(){
    nums := []int{-2,1,-3,4,-1,2,1,-5,4}
    fmt.Println("input: ", nums)
    fmt.Println("output: ", maxSubArray(nums))
}

