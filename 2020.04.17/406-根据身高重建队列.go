package main
import (
    "fmt"
    "sort"
)

// 2020.4.17 题解的稍有优化的 排序后 + 移位插入
// 排序：按h降序排序，h相同时k小的排前面
// 这样从头遍历每个元素，将元素安排在其k所代表的索引上即可
func reconstructQueue(people [][]int) [][]int {
    if len(people) <= 1 { return people }
    sort.Slice(people, func(i, j int)bool{
        if people[i][0] == people[j][0] {
            return people[i][1] < people[j][1]
        }
        return people[i][0] > people[j][0]
    })
    for i := 0; i < len(people); i++ {
        index := people[i][1]
        if i != index {
            tmp := people[i]
            for j := i; j > index; j-- {
                people[j] = people[j-1]
            }
            people[index] = tmp
        }
    }
    return people
}

// 2020.4.17 排序后暴力移位插入
// 排序：先按k升序排序，k相同时h小的排前面
// func reconstructQueue(people [][]int) [][]int {
//     if len(people) <= 1 { return people }
//     res := make([][]int, 0)
//     sort.Slice(people, func(i, j int) bool {
//         if people[i][1] == people[j][1] {
//             return people[i][0] < people[j][0]
//         }
//         return people[i][1] < people[j][1]
//     })

//     res = append(res, people[0])
//     for i := 1; i < len(people); i++ {
//         if people[i][0] >= res[len(res)-1][0] {
//             res = append(res, people[i])
//             continue
//         } else {
//             res = append(res, people[i])  // 先插入，再移位
//             count := people[i][1]
//             j := 0
//             for j < len(res)-1 {
//                 if res[j][0] >= people[i][0] {
//                     count--
//                 }
//                 if count == 0 { // 寻找插入下标：j+1
//                     for j < len(res)-1 && people[i][1] == res[j+1][1] && people[i][0] > res[j+1][0]{
//                         j++
//                     }
//                     for k := len(res)-1; k >= j+2; k-- { // 移位
//                         res[k] = res[k-1]
//                     }
//                     res[j+1] = people[i]
//                     break
//                 }
//                 j++
//             }
//         }       
//     }
//     return res
// }

func main() {
    people := [][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}
    fmt.Printf("input: people = %v\n", people)
    fmt.Println("output: ", reconstructQueue(people))

}