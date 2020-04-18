package main
import "fmt"

// 2020.4.16 滑动窗口：
// 1.右移right直到window中满足条件，即包括t中所有字母；(寻找可行解)
// 2.再右移left直到window中恰好不满足条件，取此时最小长度；(优化可行解)
// 3.重复1、2直到right超出边界
func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 || len(s) < len(t) { return "" }
    hash := make(map[byte]int, 0)
    start, length := 0, len(s)
    // matchSize用来判断window中是否满足条件，matchFlag用来判断是否存在满足条件的子串
    matchSize, matchFlag := 0, false
    for _, val := range t {
        hash[byte(val)]++
        matchSize++
    }
    window := make(map[byte]int, 0)
    left, right, match := 0, 0, 0
    for right < len(s) {
        cur := s[right]
        if cnt, ok := hash[cur]; ok {    // 当前字符cur 是需要的hash中的字符
            if window[cur] < cnt {
                match++
            }
            window[cur]++    
        }
        right++
        for match == matchSize {  //当前窗口中已包含t的所有字符，缩减left
            matchFlag = true
            if right - left < length {
                start = left
                length = right - left
            }
            cur = s[left]
            if _, ok := window[cur]; ok {
                window[cur]--
                if window[cur] < hash[cur] {
                    match--
                }
            }
            left++
        }
    }
    if matchFlag == true { return s[start:start+length] }
    return ""
}

func main() {
    s := "ADOBECODEBANC"
    t := "ABC"
    fmt.Printf("input: s = %v, t = %v\n", s, t)
    fmt.Println("output: ", minWindow(s, t))
}
