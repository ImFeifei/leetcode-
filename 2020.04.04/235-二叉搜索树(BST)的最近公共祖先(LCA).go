package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }

// 迭代
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil{
        if p.Val < root.Val && q.Val < root.Val {
            root = root.Left
        } else if p.Val > root.Val && q.Val > root.Val {
            root = root.Right
        } else {
            return root
        }
    }
    return nil
}

 // 递归
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//      if root == nil || root == p || root == q { return root }
//      if p.Val < root.Val && q.Val > root.Val { return root }
//      if p.Val > root.Val && q.Val < root.Val { return root }

//     if p.Val < root.Val { return lowestCommonAncestor(root.Left, p, q) }
//     return lowestCommonAncestor(root.Right, p, q)
//  }

func main() {
    q := &TreeNode{3,nil,nil}
    p := &TreeNode{2,&TreeNode{0,nil,nil},&TreeNode{4,q,&TreeNode{5,nil,nil}}}
    tree := &TreeNode{6, p, &TreeNode{8,&TreeNode{7,nil,nil},&TreeNode{9,nil,nil}}}
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
