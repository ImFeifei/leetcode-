package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
// 2020.4.15 迭代：反中序遍历：右-中-左
func convertBST(root *TreeNode) *TreeNode {
    if root == nil || root.Left == nil && root.Right == nil { return root }
    stack := []*TreeNode{}
    cur, sum := root, 0
    for len(stack) != 0 || cur != nil {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Right
        }
        cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
        cur.Val, sum = cur.Val + sum, sum + cur.Val
        cur = cur.Left
    }
    return root
}

// 2020.4.15 递归
// func convertBST(root *TreeNode) *TreeNode {
//     if root == nil || root.Left == nil && root.Right == nil { return root }
//     sum := 0
//     return helper(root, &sum)
// }

// func helper(root *TreeNode, sum *int) *TreeNode{
//     if root != nil {
//         helper(root.Right, sum)
//         *sum += root.Val
//         root.Val = *sum
//         helper(root.Left, sum)
//     }
//     return root
// }

func main() {
    tree := &TreeNode{5, &TreeNode{2,nil,nil}, &TreeNode{13, nil,nil}}
    //tree := &TreeNode{8, &TreeNode{4,&TreeNode{2,nil,&TreeNode{3,nil,nil}},&TreeNode{6,&TreeNode{5,nil,nil},&TreeNode{7,nil,nil}}}, &TreeNode{15,&TreeNode{11,nil,&TreeNode{12,nil,nil}},&TreeNode{16,nil,&TreeNode{18,&TreeNode{17,nil,nil},nil}}}}
    fmt.Print("input: tree = [ ")
    printTree(tree)
    fmt.Printf("]\n")
    fmt.Print("output: [ ")
    printTree(convertBST(tree))
    fmt.Printf("]\n")
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