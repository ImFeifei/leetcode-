package main
import "fmt"

// 2020.4.14 方法2 不用额外空间，顺序遍历找到出现降序后的最小元素，再逆序遍历找到出现升序后的最大元素
func findUnsortedSubarray(nums []int) int {
    if len(nums) <= 1 { return 0 }
    if nums[len(nums)-1] < nums[0] { return len(nums) }
    minNum, maxNum := nums[len(nums)-1],nums[0]
    left, right := 0, len(nums)-1
    // 找: 出现降序后的最小元素
    flag := false
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[i-1] {
            flag = true
        } 
        if flag == true {
            minNum = min(minNum, nums[i])
        }
    }
    if flag == false { return 0 }
    // 找: 出现升序后的最大元素
    flag = false
    for i := len(nums)-2; i >= 0; i-- {
        if nums[i] > nums[i+1] {
            flag = true
        } 
        if flag == true {
            maxNum = max(maxNum, nums[i])
        }
    }
    // 顺序找minNum的正确位置
    for i := 0; i < len(nums); i++ {
        if nums[i] > minNum {
            left = i
            break
        }
    }
    // 逆序找maxNum的正确位置
    for i := len(nums)-1; i >= 0; i-- {
        if nums[i] < maxNum {
            right = i
            break
        }
    }
    return right - left + 1
}

// 2020.4.14 方法1 单调栈：顺序单调递增栈找left，逆序单调递减栈找right
// func findUnsortedSubarray(nums []int) int {
//     if len(nums) <= 1 { return 0 }
//     if nums[len(nums)-1] < nums[0] { return len(nums) }
//     left, right := len(nums)-1, 0
//     var cur int
//     stack := []int{0}
//     for i := 1; i < len(nums); i++{
//         if nums[i] >= nums[stack[len(stack)-1]] {
//             stack = append(stack, i)
//         } else if nums[i] < nums[stack[len(stack)-1]]  {
//             for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
//                 cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
//             }
//             left = min(left, cur)  // 找最小left, 之所以不找到第一个降序就退出是因为后面可能还会有更小的数要往前排序才行
//             stack = append(stack, i)
//         }
//     }
//     if len(stack) == len(nums) {   // 数组已经有序
//         return 0
//     }
    
//     stack = []int{len(nums)-1}
//     for i := len(nums)-2; i >= 0; i--{
//         if nums[i] <= nums[stack[len(stack)-1]] {
//             stack = append(stack, i)
//         } else if nums[i] > nums[stack[len(stack)-1]] {
//             for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
//                 cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
//             }
//             right = max(right, cur)  // 找最大right, 之所以不找到第一个升序就退出是因为前面可能还会有更大的数
//             stack = append(stack, i)
//         }
//     }
    
//     return right - left + 1
// }
func min(x, y int)int {
    if x < y { return x }
    return y
}
func max(x, y int)int {
    if x > y { return x }
    return y
}

func main() {
    nums := []int{2, 6, 4, 8, 10, 9, 15}
    fmt.Printf("input: nums = %v\n", nums)
    fmt.Println("output: ", findUnsortedSubarray(nums))
}
