package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 

// 方法1 递归
// func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
//     if t1 == nil { return t2 }
//     if t2 == nil { return t1 }
//     t1.Val += t2.Val
//     t1.Left = mergeTrees(t1.Left, t2.Left)
//     t1.Right = mergeTrees(t1.Right, t2.Right)
//     return t1
// }

// 方法2 迭代：栈的每层存放两棵树对应位置的节点，将所有累加在t1上返回
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil { return t2 }
    if t2 == nil { return t1 }
    stack := [][]*TreeNode{{t1, t2}}
    list := []*TreeNode{}
    for len(stack) != 0 {
        list, stack = stack[len(stack)-1], stack[:len(stack)-1]
        list[0].Val += list[1].Val
        // 左子树的处理
        if list[0].Left != nil && list[1].Left != nil {
            stack = append(stack, []*TreeNode{list[0].Left, list[1].Left})
        } else if list[0].Left == nil {
            list[0].Left = list[1].Left
        }
        // 右子树的处理
        if list[0].Right != nil && list[1].Right != nil {
            stack = append(stack, []*TreeNode{list[0].Right, list[1].Right})
        } else if list[0].Right == nil {
            list[0].Right = list[1].Right
        }
    }
    return t1
}

func main() {
    tree1 := &TreeNode{1, &TreeNode{3,&TreeNode{5,nil,nil},nil}, &TreeNode{2, nil,nil}}
    tree2 := &TreeNode{2, &TreeNode{1,nil,&TreeNode{4,nil,nil}},&TreeNode{3,nil,&TreeNode{7,nil,nil}}}
    fmt.Print("input: tree1 = [")
    printTree(tree1)
    fmt.Printf("] tree2 = [")
    printTree(tree2)
    fmt.Printf("]\n")
    fmt.Print("output: [")
    printTree(mergeTrees(tree1, tree2))
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