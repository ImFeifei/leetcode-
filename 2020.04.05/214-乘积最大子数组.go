package main
import "fmt"

// 2020.4.5 状态转移方程一定要选好
// 这里只需存两个数，第i个元素之前（必须以元素i-1结尾,起始可以从任意开始）连乘的最大值和连乘的最小值
// 那么第i个元素若非负，则乘最大值得最大值；否则乘最小值得最大值，求变换过程中的最大值
func maxProduct(nums []int) int {
    if len(nums) == 1 { return nums[0] }
    res := nums[0]
    preMin, preMax := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            preMin, preMax = preMax * nums[i], preMin * nums[i]
            
        } else {
            preMin *= nums[i]
            preMax *= nums[i]
        }
        preMin = min(preMin, nums[i])
        preMax = max(preMax, nums[i])
        res = max(res, preMax)
    }
    return res
}

func max(x, y int)int {
    if x > y { return x }
    return y
}
func min(x, y int)int{
    if x < y { return x }
    return y
}


func main() {
    nums := []int{1,0,-1,2,3,-5,-2} 
    fmt.Println("input: ", nums)
    fmt.Println("output: ", maxProduct(nums))
}
