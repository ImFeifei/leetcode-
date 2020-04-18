package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }

// 2020.4.16 思路：dfs遍历树, 同时用前缀和 + hashmap
func pathSum(root *TreeNode, target int) int {
    if root == nil { return 0 }
    hash := make(map[int]int, 0)
    hash[0] = 1
    res := 0
    dfs(root, 0, target, hash, &res)
    return res
}

func dfs(root *TreeNode, sum, target int, hash map[int]int, res *int) {
    if root == nil {
        return
    }
    *res = *res + hash[sum + root.Val - target]
    hash[sum + root.Val]++
    dfs(root.Left, sum + root.Val, target, hash, res)
    dfs(root.Right, sum + root.Val, target, hash, res)
    hash[sum + root.Val]--  // 重点 确保包括左孩子的前缀和不影响到右孩子，因此撤销本root作为左孩子时加入的和
                            // 回溯到父节点去计算右兄弟
}

// 2020.4.2 prefixSum是map: 前序节点的路径和 -> 对应出现的次数
// func pathSum(root *TreeNode, sum int) int {
//     if root == nil { return 0 }
//     prefixSum := make(map[int]int)
//     // 查找前序路径curSum - sum = 0时，代表 根节点到当前节点 就是一条路径，赋值为1
//     // 即这里的赋值代表任何节点回溯到根节点都只有唯一的一条路径
//     prefixSum[0] = 1         
//     return prefix(root, sum, 0, prefixSum)
// }
// // 
// func prefix(root *TreeNode, sum int, curSum int, prefixSum map[int]int) int {
//     if root == nil { return 0 }
//     count := 0
//     curSum += root.Val
//     if v, ok := prefixSum[curSum - sum]; ok {
//         count += v
//     }
//     prefixSum[curSum]++
//     count += (prefix(root.Left, sum, curSum, prefixSum) + prefix(root.Right, sum, curSum, prefixSum))
//     prefixSum[curSum]--  //类似回溯，因为map传指针，会使回溯回去的map也变了,影响同一层左孩子算完后对右孩子的计算
//     return count
// }

 // 2020.4.1 双重递归，效率较低
// func pathSum(root *TreeNode, sum int) int {
//     if root == nil { return 0 }
//     // if root.Left == nil && root.Right == nil {
//     //     if root.Val == sum { return 1 }
//     //     return 0
//     // }
//     res := 0
//     if root.Val == sum { res++ }
//     // 不包含当前root的路径数量
//     res += (pathSum(root.Left, sum) + pathSum(root.Right, sum))
//     // 再加上包含当前root的路径数量
//     res += (pathSumIncludeRoot(root.Left, sum - root.Val) + pathSumIncludeRoot(root.Right, sum - root.Val))
    
//     return res
// }
// func pathSumIncludeRoot(root *TreeNode, sum int) int {
//     if root == nil { return 0 }
//     res := 0
//     if root.Val == sum { res++ }
//     res += (pathSumIncludeRoot(root.Left, sum - root.Val) + pathSumIncludeRoot(root.Right, sum - root.Val))
//     return res
// }

func main() {
	tree := &TreeNode{1, &TreeNode{-2, nil, nil}, &TreeNode{-3, nil, nil}}
	sum := -1
	//sum := 8
	//tree := &TreeNode{10, &TreeNode{5,&TreeNode{3,&TreeNode{3,nil,nil},&TreeNode{-2,nil,nil}},&TreeNode{2,nil,&TreeNode{1,nil,nil}}}, &TreeNode{-3,nil,&TreeNode{11,nil,nil}}}
	fmt.Print("input: tree = [")
	printTree(tree)
	fmt.Printf("]  sum = %d\n", sum)
	fmt.Println("output: ", pathSum(tree, sum))
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



