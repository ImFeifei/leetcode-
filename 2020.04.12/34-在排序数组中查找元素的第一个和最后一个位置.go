package main
import "fmt"

// 2020.4.12 双指针左右开弓查找
func searchRange(nums []int, target int) []int {
    res := []int{-1, -1}
    if len(nums) == 0 { return res }
    left, right := 0, len(nums) - 1
    for left <= right {
        if nums[left] != target && nums[right] != target {
            left++
            right--
        } else if nums[left] == target && nums[right] != target {
            right--
        } else if nums[left] != target && nums[right] == target {
            left++
        } else {
            res[0], res[1] = left, right
            break
        }
    }
    return res
}

// 2020.4.12 二分查找 分别找左中点和右中点
// func searchRange(nums []int, target int) []int {
//     res := []int{-1, -1}
//     if len(nums) == 0 { return res }
//     left, right := 0, len(nums) - 1
//     // 找start 左中点
//     for left < right {
//         mid := left + (right - left) / 2
//         if nums[mid] < target {
//             left = mid + 1
//         } else if nums[mid] > target {
//             right = mid - 1
//         } else {
//             right = mid
//         }
//     }
//     if nums[left] == target {
//         res[0] = left   //其实这里left = right
//     }
    
//     // 找end，右中点
//     left, right = 0, len(nums) - 1
//     for left < right {
//         mid := left + (right - left + 1) / 2
//         if nums[mid] < target {
//             left = mid + 1
//         } else if nums[mid] > target {
//             right = mid - 1
//         } else {
//             left = mid
//         }
//     }
//     if nums[right] == target {
//         res[1] = right   //其实这里left = right
//     }
    
//     return res
// }

func main() {
    nums := []int{5,7,7,8,8,10}
    target := 8
    fmt.Printf("input: nums = %v, target = %d\n", nums, target)
    fmt.Println("output: ", searchRange(nums, target))
}
