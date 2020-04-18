package main
import "fmt"

// 动态规划，dp[i]表示必须以元素i结尾的LIS长度，故最终返回dp中的最大值
// 对于每个元素i，倒回去找严格小于它的元素j的最长上升子序列长度，加一即可得以j,i顺序结尾的LIS的长度，找出可使长度最大的j即可
func lengthOfLIS(nums []int) int {
    if len(nums) < 2 { return len(nums) }
    dp := make([]int, len(nums))
    dp[0] = 1
    for i := 1; i < len(dp); i++ {
        dp[i] = 1
        for j := i-1; j >= 0; j-- {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }
    return maxFromSlice(dp)
}
func max(x,y int)int {
    if x > y { return x }
    return y
}
func maxFromSlice(nums []int)int {
    res := nums[0]
    for _, num := range nums {
        res = max(res, num)
    }
    return res
}


// 2020.4.6  贪心+二分 最后greedy也不是LIS，只是长度相同
// greedy中的元素表示 长度为该元素索引+1的、可选LIS中尾部元素最小的(体现贪心)那个LIS的 尾部元素
// func lengthOfLIS(nums []int) int {
//     if len(nums) < 2 { return len(nums) }
    
//     greedy := []int{nums[0]}

//     for _, num := range nums {
//         if num > greedy[len(greedy)-1] {
//             greedy = append(greedy, num)
//         } else if num < greedy[len(greedy)-1] {
//             left, right := 0, len(greedy) - 1 
//             for left < right {
//                 mid := left + (right - left) / 2
//                 if num <= greedy[mid] {
//                     right = mid
//                 } else {
//                     left = mid + 1
//                 }
//             }
//             greedy[left] = num
//         }
//     }
//     return len(greedy)
// }

func main() {
    nums := []int{3,5,6,2,5,4,19,5,6,7,12}
    fmt.Println("input: ", nums)
    fmt.Println("output: ", lengthOfLIS(nums))
}
