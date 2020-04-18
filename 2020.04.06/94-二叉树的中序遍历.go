package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
// 2020.4.6 迭代
func inorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    stack, res := []*TreeNode{}, []int{}
    for len(stack) != 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root, stack = stack[len(stack)-1], stack[:len(stack)-1]
        res = append(res, root.Val)
        root = root.Right
    }
    return res
}

// 2020.4.6 递归
// func inorderTraversal(root *TreeNode) []int {
//     if root == nil { return nil }
//     res := []int{}
//     res = append(res, inorderTraversal(root.Left)...)
//     res = append(res, root.Val)
//     res = append(res, inorderTraversal(root.Right)...)
//     return res
// }

func main() {
    //tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3,nil,nil},nil}}
    tree := &TreeNode{10, &TreeNode{5,&TreeNode{3,&TreeNode{3,nil,nil},&TreeNode{-2,nil,nil}},&TreeNode{2,nil,&TreeNode{1,nil,nil}}}, &TreeNode{-3,nil,&TreeNode{11,nil,nil}}}
    fmt.Print("input: tree = [")
    printTree(tree)
    fmt.Printf("]\n")
    fmt.Println("output: ", inorderTraversal(tree))
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
