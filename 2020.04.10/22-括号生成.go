package main
import "fmt"

// 2020.4.10 回溯
func generateParenthesis(n int) []string {
    if n == 0 { return []string{""} }  
    res := []string{}  
    backtrack("", n, n, n, &res)
    return res
}

// 回溯1
// func backtrack(ans string, n, left, right int, res *[]string) {
//     if left == 0 && right == 0 {
//         *res = append(*res, ans)
//         return
//     }
//     // 这里其实可以改为先放左括号，再顺序放右括号 如回溯2
//     if left == right {
//         backtrack(ans + "(", n, left-1, right, res)
//     } else if left != 0 && right > left {
//         backtrack(ans + "(", n, left-1, right, res)
//         backtrack(ans + ")", n, left, right-1, res)
//     } else if left == 0 && right > left {
//         backtrack(ans + ")", n, left, right-1, res)
//     }
// }
// 回溯2
func backtrack(ans string, n, left, right int, res *[]string) {
    if left == 0 && right == 0 {
        *res = append(*res, ans)
        return
    }
    // 这样写更简洁 先放左括号，再顺序放右括号
    if left > 0 {
        backtrack(ans + "(", n, left-1, right, res)
    } 
    if right > left {
        backtrack(ans + ")", n, left, right-1, res)
    }
}

func main() {
    n := 3
    fmt.Println("input: n = ", n)
    fmt.Println("output: ", generateParenthesis(n))
}
