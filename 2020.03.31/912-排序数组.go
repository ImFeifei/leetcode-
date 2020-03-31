package main
import "fmt"

// 2020.3.31 归并排序
func sortArray(nums []int) []int {
    if len(nums) == 0 || len(nums) == 1 { return nums }
    l, r := 0, len(nums)-1
    mid := l + (r - l) / 2
    return merge(sortArray(nums[:mid]), sortArray(nums[mid:]))
}
// 合并两个有序数组
func merge(nums1, nums2 []int) []int {
    if len(nums1) == 0 { return nums2 }
    if len(nums2) == 0 { return nums1 }
    res := make([]int, 0)
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] <= nums2[j] {
            res = append(res, nums1[i])
            i++
        } else {
            res = append(res, nums1[i])
            j++
        }
    }
    for ; i < len(nums1); i++ {
        res = append(res, nums1[i])
    }
    for ; j < len(nums2); j++ {
        res = append(res, nums2[j])
    }
    return res
}

// 2020.3.31 堆排序，升序排列要首先生成大顶堆(此时只能保证)，再通过末尾元素与根元素互换+调整变成升序序列
// func sortArray(nums []int) []int {
//     // 构建大顶堆
//     buildHeap(nums)
//     // 进行排序
//     for i := len(nums)-1; i > 0; i-- {
//         nums[0], nums[i] = nums[i], nums[0]
//         adjustHeap(nums, 0, i - 1)

//     }

//     return nums
// }
// func buildHeap(nums []int){
//     // 从第一个非叶子结点开始调整堆
//     for i := len(nums)/2; i >= 0; i-- {
//         adjustHeap(nums, i, len(nums)-1)
//     }
// }
// // 大顶堆
// func adjustHeap(nums []int, start, end int) {
//     cur := nums[start]
//     child := start*2 + 1
//     for child <= end {
//         // 选出更大的孩子
//         if child < end && nums[child] < nums[child+1] { child++ }
//         // 当前cur比nums[child]大，则不与当前child交换，跳出将start赋值为cur
//         if cur >= nums[child] { break }
//         // 否则当前cur比nums[child]小，将孩子上移至父节点处,实际上这里仅是复制操作而不是交换
//         nums[start] = nums[child]
//         // 移至下一层进行调整
//         start = child
//         child = start*2+1
//     }
//     nums[start] = cur
// }

// 2020.3.31 快速排序
// func sortArray(nums []int) []int {
//     if len(nums) == 0 || len(nums) == 1 { return nums }
//     index := partition(nums,0, len(nums)-1)
//     sortArray(nums[0:index])
//     sortArray(nums[index+1:])
//     return nums
// }
// func partition(nums []int, left, right int) int {
//     for left < right {
//         for left < right && nums[left] <= nums[right]{ right-- }
//         if left < right && nums[left] > nums[right]{
//             nums[left], nums[right] = nums[right], nums[left]
//             left++
//         }
//         for left < right && nums[left] <= nums[right]{ left++ }
//         if left < right && nums[left] > nums[right]{
//             nums[left], nums[right] = nums[right], nums[left]
//             right--
//         }
//     }
//     return left
// }
func main() {
    nums := []int{5,2,3,1}
    fmt.Println("input: ", nums)
    fmt.Println("output: ", sortArray(nums))
}