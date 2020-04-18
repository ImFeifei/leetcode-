package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
// 2020.4.6 迭代
 func preorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    stack := []*TreeNode{root}
    var cur *TreeNode
    res := []int{}
    for len(stack) != 0 {
        cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
        res = append(res, cur.Val)
        // 先入栈右孩子，再入栈左孩子
        if cur.Right != nil { stack = append(stack, cur.Right) }
        if cur.Left != nil { stack = append(stack, cur.Left)}
    }
    return res
}
// 2020.4.6 递归
// func preorderTraversal(root *TreeNode) []int {
//     if root == nil { return nil }
//     res := []int{}
//     preOrder(root, &res)
//     return res
// }
// func preOrder(root *TreeNode, res *[]int) {
//     if root == nil { return }
//     *res = append(*res, root.Val)
//     preOrder(root.Left, res)
//     preOrder(root.Right, res)
// }

func main() {
    //tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3,nil,nil},nil}}
    tree := &TreeNode{10, &TreeNode{5,&TreeNode{3,&TreeNode{3,nil,nil},&TreeNode{-2,nil,nil}},&TreeNode{2,nil,&TreeNode{1,nil,nil}}}, &TreeNode{-3,nil,&TreeNode{11,nil,nil}}}
    fmt.Print("input: tree = [")
    printTree(tree)
    fmt.Printf("]\n")
    fmt.Println("output: ", preorderTraversal(tree))
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
