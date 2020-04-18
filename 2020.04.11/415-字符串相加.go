package main
import "fmt"

// 从尾部开始计算，维护进位carry,低位补零
func addStrings(num1 string, num2 string) string {
    if num1 == "0" { return num2 }
    if num2 == "0" { return num1 }
    var res string
    i, j := len(num1)-1, len(num2)-1
    carry := 0
    for i >= 0 || j >= 0 {
        var a, b int
        if i >= 0 {
            a = int(num1[i] - '0')
        } else {
            a = 0
        }
        if j >= 0 {
            b = int(num2[j] - '0')
        } else {
            b = 0
        }
        sum := ((a + b + carry) % 10) + '0'
        res = string(sum) + res
        carry = (a + b + carry) / 10
        i--
        j--
    }

    if carry == 1 {
        res = "1" + res
    }
    return res
}

func main() {
    num1 := "999"
    num2 := "1"
    fmt.Printf("input: num1 = %s, num2 = %s\n", num1, num2)
    fmt.Println("output: ", addStrings(num1, num2))
}
