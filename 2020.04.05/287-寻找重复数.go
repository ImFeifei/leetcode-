package main
import "fmt"

// 不能更改原数组，则不能对原数组排序二分
// 那么就对值域二分
func findDuplicate(nums []int) int {
    if len(nums) == 2 { return nums[0] }
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right - left) / 2
        count := 0
        for _, num := range nums {
            if num <= mid { count++ }
        }
        if count > mid { //重复的数在左半边[1,mid]
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
func main() {
    nums := []int{1,3,4,2,3}
    fmt.Println("input: ", nums)
    fmt.Println("output: ", findDuplicate(nums))
}
