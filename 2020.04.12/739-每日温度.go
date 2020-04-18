package main
import "fmt"

// 2020.4.12 利用单调递减栈
func dailyTemperatures(T []int) []int {
    if len(T) == 1 { return []int{0} }
    res := make([]int, len(T))
    stack := []int{0}
    out := 0
    for i := 1; i < len(T); i++ {
        for len(stack) >= 1 && T[i] > T[stack[len(stack)-1]] {
            out, stack = stack[len(stack)-1], stack[:len(stack)-1]
            res[out] = i - out
        }
        stack = append(stack, i)
    }
    return res
}

// 2020.4.12 暴力循环遍历 时间O(n^2) 空间O(n)
// func dailyTemperatures(T []int) []int {
//     if len(T) == 1 { return []int{0} }
//     res := make([]int, len(T))
//     for i := 0; i < len(T) - 1; i++ {
//         for j := i + 1; j < len(T); j++ {
//             if T[j] > T[i] {
//                 res[i] = j - i
//                 break
//             }
//         } 
//     }
//     return res
// }

func main() {
    T := []int{73,74,75,71,69,72,76,73}
    fmt.Printf("input: T = %v\n", T)
    fmt.Println("output: ", dailyTemperatures(T))
}
