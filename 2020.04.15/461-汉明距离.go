package main
import "fmt"

// 两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
// 2020.4.15 先亦或，然后数结果的1的个数：x&(x-1)就是打掉x从右往左第一个1
func hammingDistance(x int, y int) int {
    tmp := x ^ y
    count := 0
    for tmp != 0 {
        count++
        tmp = tmp & (tmp - 1)
    }
    return count
}

func main() {
	x, y := 4, 1
	fmt.Printf("input: x = %d, y = %d\n", x, y)
	fmt.Println("output: ", hammingDistance(x, y))
}
