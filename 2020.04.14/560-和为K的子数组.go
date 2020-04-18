package main
import "fmt"

// 2020.4.14 方法2 前缀和思想 + hashmap   时间O(n) 空间O(n)
// 前缀和思想：若【前i个数的和】与【前j个数的和】的差为k,则说明i到j之间元素累加和为k
func subarraySum(nums []int, k int) int {
    if len(nums) == 0 { return 0 }
    if len(nums) == 1 && nums[0] != k { return 0 }
    hmap := make(map[int]int, 0)
    hmap[0] = 1
    sum, count := 0, 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if _, ok := hmap[sum - k]; ok {
            count += hmap[sum - k]
        } 
        hmap[sum]++
    }
    return count
}

// 2020.4.14 方法1 暴力遍历 时间O(n^2) 空间O(1)
// func subarraySum(nums []int, k int) int {
//     if len(nums) == 0 { return 0 }
//     if len(nums) == 1 && nums[0] != k { return 0 }
//     res := 0
//     for i := 0; i < len(nums); i++ {
//         sum := 0
//         for j := i; j < len(nums); j++ {
//             if sum + nums[j] == k {
//                 //fmt.Println("i:",i," j:",j)
//                 res++
//             }
//             sum += nums[j]     
//         }
//     }
//     return res
// }

func main() {
    nums := []int{1,1,1,2,-2,-2,-1,2}
    k := 2
    fmt.Printf("input: nums = %v, k = %d\n", nums, k)
    fmt.Println("output: ", subarraySum(nums, k))
}
