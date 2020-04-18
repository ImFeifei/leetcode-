package main
import (
    "fmt"
    "strconv"
)

// 2020.4.17 辅助栈
func decodeString(s string) string {
    if len(s) <= 3 { return s }
    res, cnt := "", ""
    stack, cur := [][]string{}, []string{}
    for i := 0; i < len(s); i++{
        if s[i] >= '0' && s[i] <= '9' {  // 当s[i]为数字时，将数字字符转化为字符串存储的数字cnt，用于后续倍数计算；
            cnt += string(s[i])
        } else if s[i] == '[' {  // 当s[i]为 [ 时，将当前 cnt 和 res 入栈，并分别置重置它们
            stack = append(stack, []string{cnt, res})
            cnt, res = "", ""
        } else if s[i] == ']' {  // 当s[i]为 ] 时，stack出栈，拼接字符串 res = pre_res + count * res，其中:
            cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
            count, _ := strconv.Atoi(cur[0])  // count：cur[0]  pre_res：cur[1]
            fmt.Println("pre_res: ",cur[1], " res: ", res," count: ", count)
            tmp := res
            for count > 1 {
                res += tmp
                count--
            }
            res = cur[1] + res
        } else {  //当s[i]为字母时，在 res 尾部添加s[i]
            res += string(s[i])
        }
    }
    return res
}

func main() {
    s := "3[a2[c]]"
    fmt.Printf("input: s = %v\n", s)
    fmt.Println("output: ", decodeString(s))

}