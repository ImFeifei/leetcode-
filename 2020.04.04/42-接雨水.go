package main
import "fmt"

// solution 3: 双指针 左右开弓同时向中间计算并累加 时间O(n) 空间O(1)
func trap(height []int) int {
    if len(height) < 3 { return 0 }
    start := 0 
    for i, h := range height {
        if h != 0 {
            start = i
            break
        }
    }

    res := 0
    left, right := start+1, len(height)-2
    max_left, max_right := height[start], height[len(height)-1]
    for left <= right {
        // 从左向右计算下标left可存的水
        if max_left < max_right {    // 保证短板在左
            if max_left > height[left] {
                res += (max_left - height[left])
            }
            max_left = max(max_left, height[left])
            left++
        } else { // 从右向左计算下标right可存的水  保证短板在右
            if max_right > height[right] {
                res += (max_right - height[right])
            }
            max_right = max(max_right, height[right])
            right--
        }
    }
    return res
}

// solution 2: 对解法1的时间优化  时间O(n) 空间O(n)
// 维护左右两端最大值数组，每次仅与前一个比较即可，不用循环
// func trap(height []int) int {
//     if len(height) < 3 { return 0 }
//     start := 0 
//     for i, h := range height {
//         if h != 0 {
//             start = i
//             break
//         }
//     }
//     res := 0
//     max_left := make([]int, len(height))
//     max_right := make([]int, len(height))
//     for i := start + 1; i < len(height); i++ {
//         max_left[i] = max(height[i-1], max_left[i-1])
//     }
//     for i := len(height) - 2; i >= start; i-- {
//         max_right[i] = max(height[i+1], max_right[i+1])
//     }
//     for i := start+1; i < len(height); i++ {
//         h := min(max_left[i], max_right[i])
//         if h > height[i] {
//             res += (h - height[i])
//         }
//     }
//     return res
// }

// solution 1: 按列求 时间O(n^2) 空间O(1)
// func trap(height []int) int {
//     if len(height) < 3 { return 0 }
//     // 找第一个非空柱子下标start
//     start := 0
//     for i, h := range height {
//         if h != 0 { 
//             start = i
//             break
//         }
//     }
//     // 对每一列求它上方能存的水累加（左右找高于它的最高值，较小的那个与它的高度差就是它上方能存的水）
//     res := 0
//     l, r := start, start
//     lFlag, rFlag := false, false
//     for i := start+1; i < len(height); i++ {
//         lFlag, rFlag = false, false
//         cur := height[i]
//         for j := i - 1; j >= start; j-- {
//             if height[j] > cur {
//                 l = j
//                 cur = height[j]
//                 lFlag = true
//             }
//         }
//         cur = height[i]
//         for k := i + 1; lFlag == true && k < len(height); k++ {
//             if height[k] > cur {
//                 r = k
//                 cur = height[k]
//                 rFlag = true
//             }
//         }
//         if lFlag == true && rFlag == true {
//             res += (min(height[l], height[r]) - height[i])
//         }
//     }
//     return res
// }

func min(x,y int)int {
    if x < y { return x }
    return y
}

func max(x,y int)int {
    if x > y { return x }
    return y
}


func main(){
    height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
    fmt.Println("input: ", height)
    fmt.Println("output: ", trap(height))
}

