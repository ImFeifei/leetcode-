package main
import (
    "fmt"
    "sort"
)

// 2020.4.7 难点在于去重 应该想到排序之后的左右双指针同时去重
func threeSum(nums []int) [][]int {
    if len(nums) < 3 { return nil }
    res := [][]int{}
    sort.Slice(nums, func (i, j int)bool { return nums[i] < nums[j]})
    i := 0
    for i < len(nums) - 2 && nums[i] <= 0{
        target := 0 - nums[i]
        left, right := i + 1, len(nums) - 1
        for left < right {
            if nums[left] + nums[right] < target {
                left++
                for nums[left] == nums[left-1] && left < right { left++ }
            } else if nums[left] + nums[right] > target {
                right--
                for nums[right] == nums[right+1] && left < right { right-- }
            } else {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                for nums[left] == nums[left-1] && left < right { left++ }
                for nums[right] == nums[right+1] && left < right { right-- }
            }
        } 
        i++
        for nums[i] == nums[i-1] && i < len(nums) - 2 && nums[i] <= 0 { i++ }
    }
    return res
}

func main() {
    nums := []int{0,0,0,0,0,0}
    fmt.Println("input: ", nums)
    fmt.Printf("output: %v\n", threeSum(nums))
}
