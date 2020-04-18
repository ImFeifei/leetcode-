package main
import "fmt"

// 2020.4.14 方法2 前缀和思想 + hashmap
// 由于这次是要满足和为k的倍数，因此map中的key存为 sum%k
// 这样一来，若元素i到j之间的和为n*k，那么【前j元素和】=【前i元素和】+【n*k】
// 此时，【前j元素和】% k = (【前i元素和】+【n*k】) % k =【前i元素和】% k
func checkSubarraySum(nums []int, k int) bool {
    if len(nums) <= 1  { return false }
    hm := make(map[int]int, 0)
    hm[0] = -1
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if k != 0 {
            sum = sum % k
        }
        if val, ok := hm[sum]; ok {
            if i - val > 1 {
                return true
            }
        } else{
            hm[sum] = i
        }
    }
    return false
}

// 2020.4.14 方法1 暴力循环 时间O(n^2) 空间(1)
// func checkSubarraySum(nums []int, k int) bool {
//     if len(nums) <= 1 { return false }
//     for i := 0; i < len(nums); i++ {
//         sum := nums[i]
//         for j := i+1; j < len(nums); j++ {
//             sum += nums[j]
//             if k != 0 && sum % k == 0 {
//                 return true
//             } else if k == 0 && sum == 0 {
//                 return true
//             }
//         }
//     }
//     return false
// }

func main() {
    nums := []int{23,2,4,6,7}
    k := 6
    fmt.Printf("input: nums = %v, k = %d\n", nums, k)
    fmt.Println("output: ", checkSubarraySum(nums, k))
}
