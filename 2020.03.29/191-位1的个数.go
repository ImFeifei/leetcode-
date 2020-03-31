package main
import "fmt"

// solution 2: 用n&1，n>>1数n的每个位是否为1
func hammingWeight(num uint32) int {
    if num == 0 { return 0 }
    res := 0
    for num != 0 {
        // tips: n&1还可以用作判断n的奇偶，若为1则奇，0为偶
        if num & 1 == 1 {
            res++
        }
        num >>= 1
    }
    return res
}

// solution 1: 用 n&n-1 打掉n最低位的1
// func hammingWeight(num uint32) int {
//     if num == 0 { return 0 }
//     res := 0
//     for num != 0 {
//         res++
//         num = num&(num-1)
//     }
//     return res
// }

func main() {
    var num uint32 = 11 //这里输入的是十进制
    fmt.Println("input: ", num)
    fmt.Println("output: ", hammingWeight(num))
}