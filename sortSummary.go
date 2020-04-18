package main
import "fmt"

// 1.快速排序
// func quickSort(nums []int){
// 	if len(nums) <= 1 { return }
// 	mid := partition(nums, 0, len(nums)-1)
// 	quickSort(nums[:mid])
// 	quickSort(nums[mid+1:])
// }
// func partition(nums []int, left, right int) int {
// 	for left < right {
// 		for left < right && nums[right] > nums[left] {
// 			right--
// 		}
// 		if nums[left] > nums[right] {
// 			nums[left], nums[right] = nums[right], nums[left]
// 			left++
// 		}
// 		for left < right && nums[right] > nums[left] {
// 			left++
// 		}
// 		if nums[left] > nums[right] {
// 			nums[left], nums[right] = nums[right], nums[left]
// 			right--
// 		}	
// 	}
// 	return left
// }

//2. 堆排序
func heapSort(nums []int) {
	buildHeap(nums)
	// 每次将堆顶元素和最后一个元素互换，再缩小最后一个长度调整堆
	// 因为堆顶元素可以保证为最大元素，每次将剩余元素中的最大元素换到最尾部，再调整之前再次满足大顶堆
	for i := len(nums)-1; i >= 1; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjustHeap(nums, 0, i-1)
	}
}
// 建堆
func buildHeap(nums []int) {
	for i := len(nums)/2-1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums)-1)
	}
	fmt.Println("初始建堆为: ", nums)
}
// 调整堆 升序排列则先生成大顶堆
func adjustHeap(nums []int, start, end int) {
	child := start*2 + 1
	cur := nums[start]
	for child <= end {
		if child < end && nums[child] < nums[child+1] { child += 1 }
		if cur >= nums[child] { break }
		nums[start] = nums[child]
		start = child
		child = start*2+1
	}
	nums[start] = cur
}


func main() {
	nums := []int{4,7,9,8,5,2,6,3,1}
	fmt.Println("input: ", nums)
	//quickSort(nums)
	heapSort(nums)
	fmt.Println("output: ", nums)
}