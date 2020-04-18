package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
// 按定义迭代，用栈辅助同时记录一个前序节点pre
func postorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    res := []int{}
    stack := []*TreeNode{}
    pre := &TreeNode{}
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 若此时栈顶元素没有右子树或右子树已经遍历过，则弹出并遍历
        if stack[len(stack)-1].Right == nil || stack[len(stack)-1].Right == pre {
            res = append(res, stack[len(stack)-1].Val)
            pre, stack = stack[len(stack)-1], stack[:len(stack)-1]
            root = nil
        } else if stack[len(stack)-1].Right != nil{
            root = stack[len(stack)-1].Right
        }
    }
    return res
}

// 前序： 根-左-右  后序：左-右-根 == reverse(根-右-左)
// 将前序遍历左右孩子顺序互换，再反转得后续遍历
// func postorderTraversal(root *TreeNode) []int {
//     if root == nil { return nil }
//     res := []int{}
//     stack := []*TreeNode{root}
//     var cur *TreeNode
//     for len(stack) != 0 {
//         cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
//         res = append(res, cur.Val)
//         if cur.Left != nil { stack = append(stack, cur.Left) }
//         if cur.Right != nil { stack = append(stack, cur.Right) }
//     }
//     // 反转res
//     for i := 0; i < len(res) / 2; i++ {
//         res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
//     }
//     return res
// }

func main() {
    //tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3,nil,nil},nil}}
    tree := &TreeNode{10, &TreeNode{5,&TreeNode{3,&TreeNode{3,nil,nil},&TreeNode{-2,nil,nil}},&TreeNode{2,nil,&TreeNode{1,nil,nil}}}, &TreeNode{-3,nil,&TreeNode{11,nil,nil}}}
    fmt.Print("input: tree = [")
    printTree(tree)
    fmt.Printf("]\n")
    fmt.Println("output: ", postorderTraversal(tree))
}
// 层次遍历输出tree节点值
func printTree(t *TreeNode) {
    if t == nil { 
        fmt.Println("输入的是空树！")
        return
    }
    queue := []*TreeNode{t}
    var top *TreeNode
    for len(queue) != 0 {
        top , queue = queue[0], queue[1:]
        if top == nil {
            fmt.Print("nil ")
        } else {
            fmt.Print(top.Val, " ")
            if top.Left != nil || top.Right != nil {
                queue = append(queue, top.Left)
                queue = append(queue, top.Right)
            }
        }       
    }

}
