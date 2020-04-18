package main
import "fmt"

// 2020.4.12 从41题来 原地哈希：将数字i放在下标i-1处
func findDisappearedNumbers(nums []int) []int {
    if len(nums) == 0 { return []int{} }
    var res []int
    for i := 0; i < len(nums); i++ {
        for nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }

    for i := 0; i< len(nums); i++ {
        if nums[i] != i + 1 {
            res = append(res, i+1)
        }
    }
    return res
}

func main() {
    nums := []int{4,3,2,7,8,2,3,1}
    fmt.Printf("input: nums = %v\n", nums)
    fmt.Println("output: ", findDisappearedNumbers(nums))
}
