package main
import (
	"fmt"
//	"reflect"
)


func maxDepthAfterSplit(seq string) []int {
    res := make([]int, len(seq))
    stack := []map[string]int{}    //栈内存储 字符到其索引的map映射
    indexMap := make(map[int]int, len(seq))
    for j, v := range seq {
        if len(stack) != 0 {
            // 取栈顶字符map, 获得此时栈顶字符的index；注：栈顶字符必为'(',因为一遍历到')'不但不进栈反而弹出栈顶
            tmp := stack[len(stack)-1]
            topIndex := 0
            for _, val := range tmp {
                topIndex = val
            }
            // 当此时遍历字符为'('，则与栈顶字符属于不同子序列否则会造成嵌套使层数加深，用亦或1进行取非
            // 再将符号与其index构造成map进栈
            if string(v) == "(" {
                indexMap[j] = indexMap[topIndex]^1
                next := map[string]int{}
                next[string(v)] = j
                stack = append(stack, next)
            } 
            // 当此时遍历字符为')'，则与栈顶字符属于同一子序列构成有效括号出栈
            if string(v) == ")" {
                indexMap[j] = indexMap[topIndex]
                stack = stack[:len(stack)-1]
            }
        } else { // 当栈为空，下一入栈的必为'('，让其属于此时长度更短的那个序子列
            first := map[string]int{}
            first[string(v)] = j
            stack = append(stack, first)
            indexMap[j] = lessOneInMap(indexMap) // 可删去这一行
        }       
    }
    for i := 0; i < len(seq); i++ {
        res[i] = indexMap[i]
    }
    return res
}
// 获得此时长度更短的那个子序列的标记, 使用这个函数可使得A，B序列长度（不是嵌套深度）尽可能平衡
// 不用也行 因为 ()和()()() 的嵌套深度一样，都为1，刚开始做这题的时候这里没有想到
func lessOneInMap(m map[int]int)int {
    count0, count1 := 0, 0
    for _, v := range m {
        if v == 0 {count0++}
        if v == 1 {count1++}
    }
    if count0 <= count1 {return 0}
    return 1
}

func main(){
	str := "()(())()"
	fmt.Println("input: ", str)
	fmt.Println("output: ", maxDepthAfterSplit(str))

}