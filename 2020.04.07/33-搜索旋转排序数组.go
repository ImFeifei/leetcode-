package main
import "fmt"

func search(nums []int, target int) int {
    if len(nums) == 0 { return -1 }
    left, right := 0, len(nums) - 1
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] == target { return mid }
        if nums[mid] < nums[left] { // 右边一定递增
            if target > nums[mid] && target < nums[right] {
                left = mid + 1
            } else if target == nums[mid] {
                return mid
            } else if target == nums[right] {
                return right
            } else {
                right = mid - 1
            }
        } else { //左边一定递增
            if target > nums[left] && target < nums[mid] {
                right = mid - 1
            } else if target == nums[left] {
                return left
            } else if target == nums[mid] {
                return mid
            } else {
                left = mid + 1
            }
        }
    }
    if nums[left] == target { return left }
    return -1
}

func main() {
    nums := []int{4,5,6,7,8,1,2}
    target := 1
    fmt.Printf("input: nums = %v, target = %d\n", nums, target)
    fmt.Println("output: ", search(nums, target))
}
