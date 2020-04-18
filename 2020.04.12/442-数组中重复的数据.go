package main
import "fmt"

// 2020.4.12 类似41、448题 原地哈希 将数字i放在下标i-1处
func findDuplicates(nums []int) []int {
    if len(nums) == 0 || len(nums) == 1 { return []int{} }
    res := []int{}
    for i := 0; i < len(nums); i++{
        for nums[i] != nums[nums[i] - 1] {
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        }
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] != i+1 {
            res = append(res, nums[i])
        }
    }
    return res
}

func main() {
    nums := []int{4,3,2,7,8,2,3,1}
    fmt.Printf("input: nums = %v\n", nums)
    fmt.Println("output: ", findDuplicates(nums))
}
