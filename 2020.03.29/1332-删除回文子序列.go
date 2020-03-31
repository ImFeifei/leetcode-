package main
import "fmt"

// 这脑筋急转弯题？？？怪不得easy
// 由于删除的是回文子序列而非子串，且只有a,b两种字母，那么最多就删两次，一次删全部a，再删除全部b
// 实际上就是判断一个字符串是不是回文，即正反向是否相等
func removePalindromeSub(s string) int {
    if s == "" { return 0 }
    //判断s是否为回文字符串
    if s == reverse(s) { return 1 }
    return 2
}
func reverse(str string) string{
    s := []rune(str)
    for i, j := 0, len(s)-1; i <= j; i, j = i+1,j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return string(s)
}

func main() {
	s := "ababa"
	fmt.Println("input: ", s)
	fmt.Println("output: ", removePalindromeSub(s))
}