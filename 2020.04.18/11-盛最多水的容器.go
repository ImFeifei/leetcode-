package main
import "fmt"

// 本方法参考42题 双指针 时间O(n) 空间O(1)
// 思路：水面积 = 【两端的较短板高度】*【两端之间距离】，左右指针表示【可能出现更大面积的左右边界】，每次将较短的边界向内收缩
// 因为收缩，两端之间距离肯定变小；若收缩较长边界，新边界的较短板一定小于等于当前最短板，则面积不可能变大
// 而收缩较短板，虽然两端之间距离变小，但新的较短板可能会增大，因此乘积面积也可能增大
func maxArea(height []int) int {
    if len(height) <= 1 { return 0 }
    maxArea := 0
    left, right := 0, len(height)-1
    for left < right {
        if height[left] <= height[right] {
            maxArea = max(maxArea, height[left] * (right - left))
            left++
        } else {
            maxArea = max(maxArea, height[right] * (right - left))
            right--
        }
    }
    return maxArea
}

// 2020.4.18 类似84题：柱状图最大矩形面积 以及 42题：接雨水
// 本方法先参考84题 时间平均O(nlogn)  空间O(1)
// 思路：计算每个元素作为较短边界时（即高）可得的最大面积，分两种情况：该元素为左边界 & 该元素为右边界
// 因此找元素两边大于它的值更新面积，最后取两边最大值即为该元素作为高时可得的最大面积
// 同时更新一个全局最大面积值
// func maxArea(height []int) int {
//     if len(height) <= 1 { return 0 }
//     maxArea := 0
//     for i := 0; i < len(height); i++ {
//         cur := height[i]
//         // 找左边大于cur的最大值
//         leftMax := 0
//         for left := i-1; left >= 0; left-- {
//             if height[left] >= cur {
//                 // fmt.Println("leftside: i=",i,"area: ",(i - left) * cur)
//                 leftMax = max(leftMax, (i - left) * cur)
//             }  
//         }
//         // 找右边大于cur的最大值
//         rightMax := 0 
//         for right := i+1; right < len(height); right++ {
//             if height[right] >= cur {
//                 //fmt.Println("rightside: i=",i,"area: ",(right - i) * cur)
//                 rightMax = max(rightMax, (right - i) * cur)
//             }
//         }
//         maxArea = max(maxArea, max(leftMax, rightMax))
//     }
//     return maxArea
// }

func max(x, y int) int {
    if x > y { return x }
    return y
}


func main() {
    height := []int{1,8,6,2,5,4,8,3,7}
    fmt.Printf("input: height = %v\n", height)
    fmt.Println("output: ", maxArea(height))
}
