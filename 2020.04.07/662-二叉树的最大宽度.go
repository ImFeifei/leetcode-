package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }

// 用了两个队列，一个保存层次遍历的树节点，一个对应保存节点的编号（类似于堆排序中的编号）
func widthOfBinaryTree(root *TreeNode) int {
    if root == nil { return 0 }
    width := 0
    nodeQueue, indexQueue := []*TreeNode{root}, []int{0}
    var node *TreeNode
    var index int
    for len(nodeQueue) != 0 {
        count := len(nodeQueue)
        left, right := 0, 0
        for i := 0; i < count; i++ {
            if i == 0 { left = indexQueue[0] }
            if i == count-1 { right = indexQueue[0] }
            node, nodeQueue = nodeQueue[0], nodeQueue[1:]
            index, indexQueue = indexQueue[0], indexQueue[1:]
            if node.Left != nil {
                nodeQueue = append(nodeQueue, node.Left)
                indexQueue = append(indexQueue, index*2+1)
            }
            if node.Right != nil {
                nodeQueue = append(nodeQueue, node.Right)
                indexQueue = append(indexQueue, index*2+2)
            }
            //fmt.Printf("left=%d, right=%d, node=%d, index=%d\n", left, right, node.Val, index)
            width = max(width, right - left + 1)
        }
    }
    return width
}

func max(x,y int)int {
    if x < y { return y }
    return x
}

func main() {
    p := &TreeNode{1,&TreeNode{3,&TreeNode{5,nil,nil},&TreeNode{4,nil,nil}},&TreeNode{2,nil,&TreeNode{9,nil,nil}}}
    fmt.Print("input: tree = [")
    printTree(p)
    fmt.Printf("]\n")
    fmt.Printf("output: %d\n", widthOfBinaryTree(p))
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
