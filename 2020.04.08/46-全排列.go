package main
import "fmt"

// 2020.4.8 全排列回溯
func permute(nums []int) [][]int {
    res := [][]int{}
    backtrack([]int{}, nums, &res)
    return res
}

func backtrack(ans,nums []int, res *[][]int) {
    if len(nums) == 0 {
        *res = append(*res, ans)
        return
    }
    for i := 0; i < len(nums); i++ {
        nums[0], nums[i] = nums[i], nums[0]
        backtrack(append(ans, nums[0]), nums[1:], res)
        nums[0], nums[i] = nums[i], nums[0]
    }
}

func main(){
    nums := []int{1,2,3}
    fmt.Println("input: ", nums)
    fmt.Println("output: ", permute(nums))
}
