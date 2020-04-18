package main
import "fmt"

// 2020.4.15 滑动窗口 时间O(m) 空间O(n)
func findAnagrams(s string, p string) []int {
    hash := map[byte]int{}
    res := []int{}
    for _, char := range p {
        hash[byte(char)]++
    }

    left, right := 0, 0
    tmp := cloneMap(hash)
    for right < len(s) {
        if _, ok := tmp[s[right]]; ok {
            tmp[s[right]]--
            if tmp[s[right]] == 0 {
                delete(tmp, s[right])
            }
        } else {  // 此时字符不在剩余可用字符tmp中，需要调整窗口，且包含两种情况
            if _, ok := hash[s[right]]; !ok {  // 情况1: 若此字符不在原字符串p中，则直接从它后一位开始重新增长窗口
                right++
                left = right
                tmp = cloneMap(hash)
                continue
            }
            tmp[s[left]]++  //情况2: 原字符串p中有该字符，但此时个数已用尽，被tmp中删去了，此时恢复tmp中左边界字符的剩余数量
            left++          // 滑动窗口左边界向内收缩，重新增长窗口
            continue
        }
        right++
        if len(tmp) == 0 {
            res = append(res, left)
            tmp[s[left]]++
            left++
        }
    }
    return res
}

// 2020.4.15 暴力遍历 用map判断是否元素一样 超时
// 时间O(m*n)  空间O(n) , m、n分别为s、p两串长度
// func findAnagrams(s string, p string) []int {
//     hash := map[byte]int{}
//     res := []int{}
//     for _, char := range p {
//         hash[byte(char)]++
//     }
//     for i := 0; i <= len(s) - len(p); i++ {
//         tmp := cloneMap(hash)
//         for j := i; j < i + len(p); j++ {
//             if _, ok := tmp[s[j]]; !ok{
//                 if  _, ok := hash[s[j]]; !ok {  // 剪枝  然而还是超时
//                     i = j
//                 }
//                 break
//             } else {
//                 tmp[s[j]]--
//                 if tmp[s[j]] == 0 {
//                     delete(tmp, s[j])
//                 }
//             }
            
//             if len(tmp) == 0 {
//                 res = append(res, i)
//             }
//         }
//     }
//     return res
// }
func cloneMap(m map[byte]int) map[byte]int {
    res := make(map[byte]int, len(m))
    for k, v := range m {
        res[k] = v
    }
    return res
}


func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Printf("input: s = %s, p = %s\n", s, p)
	fmt.Println("output: ", findAnagrams(s, p))
}
