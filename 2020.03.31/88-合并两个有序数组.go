package main
import "fmt"

// 2020.03.31 从后向前
func merge(nums1 []int, m int, nums2 []int, n int)  {
    if len(nums2) == 0 { return }
    tail := m + n -1
    i, j := m-1, n-1
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[tail] = nums1[i]
            i--
        } else {
            nums1[tail] = nums2[j]
            j--
        }
        tail--
    }
    for j >= 0 {
        nums1[tail] = nums2[j]
        tail--
        j--
    }
}

// 2020.03.31 插入排序，暴力移位
// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     if len(nums2) == 0 { return }
//     i, j := 0, 0
//     for i < m+j && j < n {
//         if nums1[i] > nums2[j] {
//             for k := m + j; k > i; k-- {
//                 nums1[k] = nums1[k-1]
//             }
//             nums1[i] = nums2[j]
//             j++
//         } 
//         i++
//     }
//     for j < n {
//         nums1[i] = nums2[j]
//         j++
//         i++
//     }
// }
func main() {
    nums1 := []int{1,2,3,0,0,0}
    nums2 := []int{2,5,6}
    m, n := 3, 3
    fmt.Printf("input: nums1 = %v, m = %d, nums2 = %v, n = %d\n", nums1, m, nums2, n)
    merge(nums1, m, nums2, n)
    fmt.Println("output: ", nums1)
}