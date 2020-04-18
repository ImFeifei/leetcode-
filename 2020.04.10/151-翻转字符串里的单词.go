package main
import (
    "fmt"
    "strings"
)

// 使用拼接API：先处理掉首尾空格 再保存单词到slice，最后用空格拼接
func reverseWords(s string) string {
    if len(s) == 0 { return s }
    words := []string{}
    start, end := 0, len(s)-1
    for start < len(s) - 1 && s[start] == ' ' { start++ }
    for end >= 0 && s[end] == ' '{ end-- }
    l, r := end, end
    for l >= start {
        for l >= start && s[l] != ' ' { l-- }
        words = append(words, s[l+1:r+1])
        for l >= start && s[l] == ' ' { l-- }
        r = l
    }
    //fmt.Println(words)
    return strings.Join(words, " ")
}

// 2020.4.10 自己的暴力解法，从后向前遍历
// hold用来控制右指针是否前移，first代表当前是否是第一个单词（用来控制空格)
// func reverseWords(s string) string {
//     if len(s) == 0 { return s }
//     res := ""
//     l, r := len(s) - 1, len(s) - 1
//     hold, first := false, true
//     for l >= 0 {
//         if s[l] != ' ' && hold == false{
//             if l == 0 {
//                 if first == false { res += " "}
//                 res += string(s[l])
//             } else {
//                 r = l
//                 hold = true
//             }  
//         } else if s[l] != ' ' && hold == true {
//             if l == 0 {
//                 if first == false { res += " " }
//                 res += s[l:r+1]
//                 first = false
//             }
//         } else if s[l] == ' ' && hold == true {
//             if first == false { res += " " }
//             res += s[l+1:r+1]
//             first = false
//             hold = false  
//         } 
//         l--
//     }
//     return res
// }

func main() {
    s := "a good   example"    
    fmt.Println("input: s = ", s)
    fmt.Println("output: ", reverseWords(s))
}
