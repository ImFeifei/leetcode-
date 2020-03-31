package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
// 链表拼接
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p, q := headA, headB
    for p != q {
        if p != nil {
            p = p.Next
        } else {
            p = headB
        }
        if q != nil {
            q = q.Next
        } else {
            q = headA
        }
    }
    return p
}
// 常规暴力：遍历两长度 -> 作差 -> 将长的那个链表移至和短链表到末尾长度一致处开始遍历
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
//     lengthA, lengthB := 0, 0
//     p, q := headA, headB
//     for p != nil {
//         lengthA++
//         p = p.Next
//     }
//     for q != nil {
//         lengthB++
//         q = q.Next
//     }
//     if lengthA >= lengthB {
//         for i := 0; i < lengthA - lengthB; i++ {
//             headA = headA.Next
//         }     
//     } else {
//         for i := 0; i < lengthB - lengthA; i++ {
//             headB = headB.Next
//         } 
//     }
//     for headA != nil {
//         if headA == headB{
//             return headB
//         }
//         headA = headA.Next
//         headB = headB.Next
//     }
//     return nil
// }
func main() {
    common := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
    headA := &ListNode{4, &ListNode{1, common}}
    headB := &ListNode{5, &ListNode{0, &ListNode{1, common}}}
    fmt.Print("input: ")
    printLinkList(headA)
    fmt.Print("\t")
    printLinkList(headB)
    fmt.Print("\noutput: ")
    printLinkList(getIntersectionNode(headA, headB))
    // 防止zshell输出一个高亮%
    fmt.Print("\n")

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