package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }

 // 递归
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q { return root }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil { return root }
    if left == nil { return right }
    return left
  
}

func main() {
    q := &TreeNode{7,nil,nil}
    p := &TreeNode{5,&TreeNode{6,nil,nil},&TreeNode{2,q,&TreeNode{4,nil,nil}}}
    tree := &TreeNode{3, p, &TreeNode{1,&TreeNode{0,nil,nil},&TreeNode{8,nil,nil}}}
    fmt.Print("input: tree = [")
    printTree(tree)
    fmt.Printf("]  p = %d  q = %d\n", p.Val, q.Val)
    fmt.Printf("output: %d\n", lowestCommonAncestor(tree, p, q).Val)
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