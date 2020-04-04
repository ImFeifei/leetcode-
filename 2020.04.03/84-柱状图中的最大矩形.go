package main
import "fmt"

// summary: 从【以i为高的最大矩形面积】入手


// 巧用栈2： 维护一个非严格的单调递增栈存索引。
// 核心思想：从【以i为高的最大矩形面积】入手更新最大值，即每次弹出栈顶时，意味着以该栈顶为高的矩形无法再向右侧延伸，故求左边宽度
// 在上一做法基础上在原数组末尾再加一个高度-1，使得遍历完成后栈必定为空（因为栈内下标代表的元素 >= 0，-1入栈则前面都要出栈）
// 不用最后再循环一遍出栈；计算也不用变动，因为以-1为高的面积是负值不会被选择 时间O(n) 空间O(n)
func largestRectangleArea(heights []int) int {
    if len(heights) == 0 { return 0 }
    if len(heights) == 1 { return heights[0] }
    // 优化点
    heights = append(heights, -1)
    //stack := []int{}
    stack := []int{-1}
    max_area := 0
    var cur int
    for i, height := range heights {
        if i == 0 { 
            stack = append(stack, i)
            continue
        }
        if len(stack) == 1 || height >= heights[stack[len(stack)-1]] {
            stack = append(stack, i)
        } else {
            for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] > height  {
                cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
                // 这里存在问题，若cur是栈内最后一个元素，则stack = [],再取stack[len(stack)-1]就报错了
                // 因此，解决方式是初始时在栈内填一个兜底元素-1，这样也不影响计算
                max_area = maxFromTwoSum(max_area, heights[cur] * (i - stack[len(stack)-1] - 1) )
            }
            stack = append(stack, i)
        }
    }
    return max_area
}

// 巧用栈1 维护一个非严格的单调递增栈存索引。时间O(n) 空间O(n)
// 核心思想：每次弹出栈顶时，意味着以该栈顶为高的矩形无法再向右侧延伸，故求左边宽度
// 遍历数组元素：当遍历到的元素值大于栈顶元素，入栈；
// 否则非但不入栈，反而弹出当前栈顶元素，弹出同时计算以该弹出元素作为高可得到的最大矩形面积，并更新最大面积
// 直到弹出后的栈顶元素小于当前遍历到的数组元素，再将当前遍历元素入栈，重复line2
// func largestRectangleArea(heights []int) int {
//     if len(heights) == 0 { return 0 }
//     if len(heights) == 1 { return heights[0] }
    
//     //stack := []int{}
//     stack := []int{-1}
//     max_area := 0
//     var cur int
//     for i, height := range heights {
//         if i == 0 { 
//             stack = append(stack, i)
//             continue
//         }
//         if len(stack) == 1 || height >= heights[stack[len(stack)-1]] {
//             stack = append(stack, i)
//         } else {
//             for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] > height  {
//                 cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
//                 // 这里存在问题，若cur是栈内最后一个元素，则stack = [],再取stack[len(stack)-1]就报错了
//                 // 因此，解决方式是初始时在栈内填一个兜底元素-1，这样也不影响计算
//                 max_area = maxFromTwoSum(max_area, heights[cur] * (i - stack[len(stack)-1] - 1) )
//             }
//             stack = append(stack, i)
//         }
//     }
//     // 数组元素遍历完后栈长不为1（存在兜底-1），说明此时栈中元素作为高时都可以延伸到数组最右端
//     for len(stack) != 1 {
//         cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
//         max_area = maxFromTwoSum(max_area, heights[cur] * (len(heights) - stack[len(stack)-1] - 1))
//     }
//     return max_area
// }

// 分治 最大面积的矩形分为三种：
//【包括最矮的矩形，长度为数组长度】【不包括最矮矩形且在最矮矩形左边】【不包括最矮矩形且在最矮矩形右边】
// 求这三者最大值即可 拆分为二分递归，时间平均O(nlogn)，最坏数组有序时O(n^2) 空间O(nlogn)
// 让我想起另一道题 最大子序和的分治做法
// func largestRectangleArea(heights []int) int {
//     if len(heights) == 0 { return 0 }
//     if len(heights) == 1 { return heights[0] }
//     // 找最矮矩形的index和min_height
//     min_index, min_height:= 0, heights[0]
//     for i, height := range heights {
//         if height < min_height {
//             min_height = height
//             min_index = i
//         }
//     }
//     cross_area := min_height * len(heights)
//     max_area_left := largestRectangleArea(heights[:min_index])
//     max_area_right := largestRectangleArea(heights[min_index+1:])
//     return maxFromTwoSum(cross_area, maxFromTwoSum(max_area_left, max_area_right))
// }

// 暴力3 优化 找slice内最小值
// 时间O(n^2) 空间O(1)
// func largestRectangleArea(heights []int) int {
//     if len(heights) == 0 { return 0 }
//     if len(heights) == 1 { return heights[0] }
//     max_area := heights[0]
//     //min_height := heights[0]
//     for i := 1; i < len(heights); i++ {
//         max_area = maxFromTwoSum(max_area, heights[i])
//         min_height := heights[i]
//         for j := i-1; j >= 0; j-- {
//             // 遍历左起始位置
//             min_height = minFromTwoSum(heights[j],min_height)
//             max_area = maxFromTwoSum((i-j+1)*min_height, max_area)
//         }
//     }
//     return max_area
// }

// 暴力2 遍历所有矩形 第94个case有两万条数据会超时 
// 时间O(n^3) 空间O(1)
// func largestRectangleArea(heights []int) int {
//     if len(heights) == 0 { return 0 }
//     if len(heights) == 1 { return heights[0] }
//     max_area := heights[0]
//     for i := 1; i < len(heights); i++ {
//         max_area = maxFromTwoSum(max_area, heights[i])
//         for j := i-1; j >= 0; j-- {
//             // 遍历左起始位置
//             max_area = maxFromTwoSum((i-j+1)*minFromSlice(heights[j:i+1]), max_area)
//         }
//     }
//     return max_area
// }

// 2020.4.3 暴力1 遍历所有矩形 第94个case有两万条数据会超时
// 时间O(n^3) 空间O(n)
// 根本没有状态转移方程 用一个常量计算dp最大值即可 被dp洗脑了
// func largestRectangleArea(heights []int) int {
//     if len(heights) == 0 { return 0 }
//     if len(heights) == 1 { return heights[0] }
//     // dp[i] 表示以第i个柱子为右下角的可得到的最大面积， 故最终返回max(dp)
//     dp := make([]int, len(heights))
//     dp[0] = heights[0]
//     // 初始化dp[i] = heights[i]*1，至少都包括自己的面积；再向前找左下角
//     for i := 1; i < len(heights); i++ {
//         dp[i] = heights[i]
//         for j := i-1; j >= 0; j-- {
//             // 遍历左下角
//             dp[i] = maxFromTwoSum((i-j+1)*minFromSlice(heights[j:i+1]), dp[i])
//         }
//     }
//     return maxFromSlice(dp)
// }
func maxFromTwoSum(x, y int)int {
    if x > y { return x }
    return y 
}
func minFromTwoSum(x,y int) int {
    if x < y { return x }
    return y
}
func maxFromSlice(nums []int)int {
    res := 0
    for _, num := range nums{
        if num > res {
            res = num
        }
    }
    return res
}
func minFromSlice(nums []int)int {
    res := nums[0]
    for _, num := range nums{
        if num < res {
            res = num
        }
    }
    return res
}

func main() {
    heights := []int{3,1,3,2,2}
    fmt.Println("input: ", heights)
    fmt.Println("output: ", largestRectangleArea(heights))
}

