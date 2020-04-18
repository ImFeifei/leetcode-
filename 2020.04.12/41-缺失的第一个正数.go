package main
import (
	"fmt"
	"sort"
)

// 方法3 先排序，再从1开始找，找到就加一找下一个正整数 
// 时间O(nlogn) 空间O(1)。时间主要在排序，后面遍历只有O(n)
func firstMissingPositive(nums []int) int {
    if len(nums) == 0 { return 1 }
    sort.Slice(nums, func(i, j int)bool{ return nums[i]<nums[j] })
    res := 1
    for i := 0; i < len(nums); i++ {
        if nums[i] == res {
            res++
        }
    }
    return res
}

// 方法2 原地哈希 时间O(n) 空间O(1)
// 将正整数字i放在下标为i-1的位置, 最后遍历若下标i位置的数不是i+1,说明缺少的就是i+1
// func firstMissingPositive(nums []int) int {
//     if len(nums) == 0 { return 1 }
//     for i := 0; i < len(nums); i++ {
//         // nums[i] != nums[nums[i]-1]条件是为了考虑重复值的情况
//         // 或者更直接地说,只是要确定nums[i]这个值有没有在它该在的位置上,此时其实不需要nums[i] != i + 1
//         for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i + 1 && nums[i] != nums[nums[i]-1] {
//             nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
//         }
//     }

//     for i := 0; i < len(nums); i++ {
//         if nums[i] != i + 1{
//             return i+1
//         }
//     }
//     return len(nums)+1
// }

// 方法1 哈希 时间O(n) 空间O(n)
// func firstMissingPositive(nums []int) int {
//     if len(nums) == 0 { return 1 }
//     hm := make(map[int]int, len(nums))
//     for _, num := range nums {
//         hm[num]++
//     }

//     // 因为结果只可能出现在[1,len(nums)+1]之间，所以只要遍历这区间内第一个没出现的数
//     for i := 1; i < len(nums) + 1; i++ {
//         if _, ok := hm[i]; !ok {
//             return i
//         }
//     }
//     return len(nums)+1
// }

func main() {
    nums := []int{3,4,-1,1}
    fmt.Printf("input: nums = %v\n", nums)
    fmt.Println("output: ", firstMissingPositive(nums))
}
