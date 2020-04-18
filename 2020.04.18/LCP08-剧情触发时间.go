package main
import "fmt"

// 2020.4.18 先暴力超时  然后想到优化思路：二分法 因为requirements是无序的，而状态的增长必定递增
// 对于有序的序列在时间方面上的优化可以想到【二分】
func getTriggerTime(increase [][]int, requirements [][]int) []int {
    res := []int{}
    status := [][]int{{0,0,0}}
    for i := 0; i < len(increase); i++ {
        tmp := status[len(status)-1]
        status = append(status, []int{tmp[0]+increase[i][0], tmp[1]+increase[i][1], tmp[2]+increase[i][2]})
    }

    for i := 0; i < len(requirements); i++ {
        left, right := 0, len(status) - 1
        a, b, c := requirements[i][0], requirements[i][1], requirements[i][2]
        for left < right {
            mid := left + (right - left) / 2
            ma, mb, mc := status[mid][0], status[mid][1], status[mid][2]
            if ma >= a && mb >= b && mc >= c {
                right = mid
            } else {
                left = mid + 1
            }
        }
        if status[left][0] >= a && status[left][1] >= b && status[left][2] >= c {
            res = append(res, left)
        } else {
            res = append(res, -1)
        }
    }
    return res
}

// 2020.4.18 一开始的暴力遍历解法 超时
// func getTriggerTime(increase [][]int, requirements [][]int) []int {
//     a, b, c := 0, 0, 0
//     day := 0
//     res := make([]int, len(requirements))
//     for i := 0; i < len(res); i++ {
//         res[i] = -1
//         if requirements[i][0] == 0 && requirements[i][1] == 0 && requirements[i][2] == 0 {
//             res[i] = day
//         }
//     }
    
//     for j := 0; j < len(increase); j++ {
//         a, b, c = a + increase[j][0], b + increase[j][1], c + increase[j][2]
//         day++
//         for i := 0; i < len(requirements); i++ {
//             a1, b1, c1 := requirements[i][0], requirements[i][1], requirements[i][2]
//             if a >= a1 && b >= b1 && c >= c1 && res[i] == -1{
//                 res[i] = day
//             }
//         }
//     }
//     return res
// }


func main() {
    increase := [][]int{{2,8,4},{2,5,0},{10,9,8}}
    requirements := [][]int{{2,11,3},{15,10,7},{9,17,12},{8,1,14}}
    fmt.Printf("input: \nincrease = %v\nrequirements = %v\n", increase, requirements)
    fmt.Println("output: ", getTriggerTime(increase, requirements))

}