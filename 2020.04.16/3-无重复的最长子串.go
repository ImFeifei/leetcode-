package main
import "fmt"

// 2020.4.16 方法2：优化方法1 hashmap中存 字符 -> 出现的最大下标
func lengthOfLongestSubstring(s string) int {
    if len(s) <= 1 { return len(s) }
    hash := make(map[byte]int, 0)
    res := 0
    left, right := 0, 0
    for right < len(s) {
        cur := s[right]
        if val, ok := hash[cur]; !ok {
            hash[cur] = right
            right++
        } else {   // 该字符已存在，判断索引在left之外还是之内
            if val >= left {  // 该索引在left之内，即出现了重复
                left = val + 1
            }
            hash[cur] = right
            right++
        }
        res = max(res, right - left)
    }
    return res
}

// 2020.4.16 方法1：滑动窗口 + hashmap
// hashmap中存当前不重复的子串字符 -> 个数（必为1）
// func lengthOfLongestSubstring(s string) int {
//     if len(s) <= 1 { return len(s) }
//     hash := make(map[byte]int, 0)
//     res := 0
//     left, right := 0, 0
//     for right < len(s) {
//         cur := s[right]
//         if _, ok := hash[cur]; !ok {
//             hash[cur]++
//             right++
//         } else {
//             res = max(res, len(hash))
//             delete(hash, s[left])
//             left++
//         }
//     }
//     res = max(res, len(hash))
//     return res
// }
func max(x, y int)int {
    if x > y { return x }
    return y
}

func main() {
    s := "pwwkew"
    fmt.Printf("input: s = %v\n", s)
    fmt.Println("output: ", lengthOfLongestSubstring(s))
}
