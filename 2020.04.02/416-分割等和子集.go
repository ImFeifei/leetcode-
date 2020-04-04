package main
import "fmt"

// dp，类比0-1背包问题 dp[i][j]：前i个元素中是否存在子集和为j的子集(true or false)
// 故最终返回值为dp[len(nums)-1][target]
func canPartition(nums []int) bool {
    if len(nums) == 1 { return false }
    totalSum := sumOfSlice(nums)
    if totalSum % 2 == 1 { return false }
    // 找是否存在和为target的子集
    target := totalSum / 2

    dp := make([][]bool, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = make([]bool, target+1)
    }
    // 初始化dp[0][...],第一个数只能凑出nums[0]这个子集和
    if nums[0] == target {
        return true
    } else if nums[0] < target {
        dp[0][nums[0]] = true
    }
    // 从第二行开始填充dp
    for i := 1; i < len(nums); i++ {
        for j := 0; j < target+1; j++ {
            if nums[i] == j {
                    dp[i][j] = true
            } else if nums[i] < j {  // 凑出j可能是加上当前元素nums[i]，也可能是前i-1个本就已经能凑出j
                dp[i][j] = dp[i-1][j - nums[i]] || dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
        if dp[i][target] == true { return true } 
    }
    return dp[len(nums)-1][target]
}


// // 2020.4.2 回溯法遍历所有子集看是否有和为sum/2 ,当前子集和>target或res已被标记为true时直接返回 会超时
// func canPartition(nums []int) bool {
//     if len(nums) == 1 { return false }
//     totalSum := sumOfSlice(nums)
//     if totalSum % 2 == 1 { return false }
//     // 找是否存在和为target的子集
//     target := totalSum / 2
//     res := false  
//     backtrack([]int{}, nums, target, &res, len(nums))
//     return res
// }
// func backtrack(cur, nums []int, target int, res *bool, length int) {
//     if *res == true { return }
//     if sumOfSlice(cur) == target {
//         *res = true
//     } else if sumOfSlice(cur) > target {
//         return
//     }
    
//     for i, num := range nums {
//         backtrack(cur, nums[i+1:], target, res, length)
//         cur = append(cur, num)
//         backtrack(cur, nums[i+1:], target, res, length)
//         cur = cur[:len(cur)-1]
//     }
// }
func sumOfSlice(nums []int) int{
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func main() {
    nums := []int{23,13,11,7,6,5,5}
    fmt.Println("input: ", nums)
    res := canPartition(nums)
    fmt.Println("output: ", res)
}
