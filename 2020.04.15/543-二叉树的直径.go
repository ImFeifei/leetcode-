package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
// 容意想不到的点：不一定经过root, 可能在某中间节点上返回遍历其左右孩子
 // 每个节点当作折返点root求最大值
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil { return 0 }
    leftDepth := getDepth(root.Left)
    rightDepth := getDepth(root.Right)
    return max(leftDepth + rightDepth, max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)))
}

func getDepth(root *TreeNode) int {
    if root == nil { return 0 }
    depth := 1
    depth += max(getDepth(root.Left), getDepth(root.Right))
    return depth
}

func max(x, y int) int {
    if x > y { return x }
    return y
}


func main() {
    tree := &TreeNode{1, &TreeNode{2,&TreeNode{4,&TreeNode{6,&TreeNode{8,nil,nil},nil},nil},&TreeNode{5,nil,&TreeNode{7,nil,&TreeNode{9,nil,nil}}}}, &TreeNode{3, nil,nil}}
    //tree := &TreeNode{8, &TreeNode{4,&TreeNode{2,nil,&TreeNode{3,nil,nil}},&TreeNode{6,&TreeNode{5,nil,nil},&TreeNode{7,nil,nil}}}, &TreeNode{15,&TreeNode{11,nil,&TreeNode{12,nil,nil}},&TreeNode{16,nil,&TreeNode{18,&TreeNode{17,nil,nil},nil}}}}
    fmt.Print("input: tree = [ ")
    printTree(tree)
    fmt.Printf("]\n")
    fmt.Println("output: ", diameterOfBinaryTree(tree))
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