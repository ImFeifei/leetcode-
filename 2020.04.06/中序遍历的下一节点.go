package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
    Parent *TreeNode
 }
 
func nextNodeOfInorderTraversal(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil { return nil }
	// p节点有右子树，则中序的下一节点是其右子树的最左节点
	if p.Right != nil {
		q := p.Right
		for q.Left != nil {
			q = q.Left
		}
		return q
	}
	// p节点没有右子树，若p为父亲的左孩子，则中序下一节点为父节点
	if p.Parent != nil && p == p.Parent.Left { return p.Parent}

	// p节点没有右子树，且p为父亲的右孩子，则中序下一节点是其父亲所在子树作为左孩子的父节点
	for p.Parent != nil && p == p.Parent.Right {
		p = p.Parent
	}
	return p.Parent
}

func main() {
	node_10 := &TreeNode{4, nil, nil, nil}
	node_9 := &TreeNode{1, node_10, nil, nil}
	node_8 := &TreeNode{-2, nil, nil, nil}
	node_7 := &TreeNode{0, nil, nil, nil}
	node_6 := &TreeNode{11, nil, nil, nil}
	node_3 := &TreeNode{-3, nil, node_6, nil}
	node_4 := &TreeNode{3, node_7, node_8, nil}
	node_5 := &TreeNode{2, nil, node_9, nil}
	node_2 := &TreeNode{5, node_4, node_5, nil}
	node_1 := &TreeNode{10, node_2, node_3, nil}
	node_2.Parent = node_1
	node_3.Parent = node_1
	node_4.Parent = node_2
	node_5.Parent = node_2
	node_6.Parent = node_3
	node_7.Parent = node_4
	node_8.Parent = node_4
	node_9.Parent = node_5
	node_10.Parent = node_9
	
    //tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3,nil,nil},nil}}
    //tree := &TreeNode{10, &TreeNode{5,&TreeNode{3,&TreeNode{3,nil,nil},&TreeNode{-2,nil,nil}},&TreeNode{2,nil,&TreeNode{1,nil,nil}}}, &TreeNode{-3,nil,&TreeNode{11,nil,nil}}}
    fmt.Print("input: tree = [")
    printTree(node_1)
    fmt.Printf("]\n")
    fmt.Println("output: ", nextNodeOfInorderTraversal(node_1, node_5))
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
