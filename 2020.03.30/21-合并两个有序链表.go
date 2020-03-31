package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil { return l2 }
    if l2 == nil { return l1 }
    
    if l1.Val <= l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2
    }
}

 // 迭代
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//     dummy := &ListNode{}
//     head := dummy
//     for l1 != nil && l2 != nil {
//         if l1.Val <= l2.Val {
//             head.Next = l1
//             l1 = l1.Next
//         } else {
//             head.Next = l2
//             l2 = l2.Next
//         }
//         head = head.Next
//     }
//     if l1 != nil {
//         head.Next = l1
//     }
//     if l2 != nil {
//         head.Next = l2
//     }
//     return dummy.Next
// }

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4,nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4,nil}}}
	fmt.Print("input: ")
	printLinkList(l1)
	fmt.Print("\t")
	printLinkList(l2)
	//zsh打印若末尾无换行符会显示一个高亮百分号，与程序无关
	fmt.Print("\noutput: ")
	printLinkList(mergeTwoLists(l1, l2))

}

func printLinkList(l *ListNode) {
	for l != nil {
		if l.Next == nil {
			fmt.Printf("%d", l.Val)
		} else {
			fmt.Printf("%d -> ", l.Val)
		}
		l = l.Next
	}
}
