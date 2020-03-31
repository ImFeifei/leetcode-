package main
import "fmt"

// 递归 结束条件：当剩下两个元素时，若m为奇数则去掉index=1的元素，返回最终存活元素index=0；若m为偶则相反
// 重要的点：n个元素操作一次后去掉第 m%n 个，此时剩余n-1个元素；
// 相当于当前这n-1个元素中，此时从第 m%n 个重新开始编号
// 那么最终剩余的值就是从第 m%n 个往后数【n-1个元素最终存活的数所在位置(index)】个
// 最后因为是圈，取mod n
func lastRemaining(n int, m int) int {
    if n == 2 {
        if m & 1 == 0 {
            return 0
        } else {
            return 1
        }
    }
    return (lastRemaining(n-1, m) + m) % n

}

func main() {
	n := 5
	m := 3
	fmt.Printf("input: n = %d, m = %d\n", n, m)
	fmt.Println("output: ", lastRemaining(n,m))
}