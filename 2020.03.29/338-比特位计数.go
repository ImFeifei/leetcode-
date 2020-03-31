package main
import "fmt"

// solution 2: 不分奇偶，用n&n-1可打掉n的从右向左数第一个1
func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 1; i <= num; i++ {
        // 1的个数 = 打掉最低位1后的数中1的个数 + 打掉的一位
        res[i] = res[i&(i-1)] + 1
    }  
    return res  
}

// solution 1: 分奇偶情况dp
// func countBits(num int) []int {
//     dp := make([]int, num+1)
//     for i:=1; i <= num; i++ {
//         if i % 2 != 0{
//             dp[i] = dp[i-1] + 1
//         } else {
//             dp[i] = dp[i/2]
//         }
//     }
//     return dp
// }

func main() {
	num := 5
	fmt.Println("input: ", num)
	fmt.Println("output: ", countBits(num))
}