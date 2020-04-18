package main
import "fmt"

// 二分
func mySqrt(x int) int {
    if x <= 1 { return x }
    left, right := 1, x / 2
    for left < right {
        mid := left + (right - left + 1) / 2
        if x - mid * mid < 0 {
            right = mid - 1
        } else {
            left = mid
        }
    }
    return left
}

func main() {
    x := 19
    fmt.Println("input: x = ", x)
    fmt.Println("output: ", mySqrt(x))
}
