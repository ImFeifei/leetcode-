package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
 // 方法1  先将链表反转，再相加,最后再反转
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     l1 = reverseLinkList(l1)
//     l2 = reverseLinkList(l2)
//     a, b, carry := 0, 0, 0
//     dummy := &ListNode{}
//     tail := dummy
//     for l1 != nil || l2 != nil {
//         if l1 != nil {
//             a = l1.Val
//             l1 = l1.Next
//         } else {
//             a = 0
//         }
//         if l2 != nil {
//             b = l2.Val
//             l2 = l2.Next
//         } else {
//             b = 0
//         }
//         cur := &ListNode{
//             Val: (a + b + carry) % 10,
//         }
//         carry = (a + b + carry) / 10
//         tail.Next = cur
//         tail = tail.Next 
//     }
//     if carry == 1 {
//         cur := &ListNode{Val: 1}
//         tail.Next = cur
//     }
//     return reverseLinkList(dummy.Next)
// }

func reverseLinkList(head *ListNode) *ListNode {
    var dummy_1 *ListNode
    pre := dummy_1
    for head != nil {
        next := head.Next
        head.Next = pre
        pre = head
        head = next
    }
    return pre
}

// 方法2 使用栈
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    stack_1, stack_2 := []int{}, []int{}
    for l1 != nil {
        stack_1 = append(stack_1, l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        stack_2 = append(stack_2, l2.Val)
        l2 = l2.Next
    }
    a, b, carry := 0, 0, 0
    dummy := &ListNode{}
    tail := dummy
    for len(stack_1) > 0 || len(stack_2) > 0 {
        if len(stack_1) > 0 {
            a, stack_1 = stack_1[len(stack_1)-1], stack_1[:len(stack_1)-1]
        } else {
            a = 0
        }
        if len(stack_2) > 0 {
            b, stack_2 = stack_2[len(stack_2)-1], stack_2[:len(stack_2)-1]
        } else {
            b = 0
        }
        cur := &ListNode{Val: (a + b+ carry) % 10}
        carry = (a + b+ carry) / 10
        tail.Next = cur
        tail = tail.Next
    }
    if carry == 1 {
        tail.Next = &ListNode{Val:1}
    }
    return reverseLinkList(dummy.Next)
}

func main() {
    l1 := &ListNode{7, &ListNode{2, &ListNode{4,&ListNode{3,nil}}}}
    l2 := &ListNode{5, &ListNode{6, &ListNode{4,nil}}}
    fmt.Print("input: ")
    printLinkList(l1)
    fmt.Print("  ")
    printLinkList(l2)
    //zsh打印若末尾无换行符会显示一个高亮百分号，与程序无关
    res := addTwoNumbers(l1, l2)
    fmt.Print("\noutput: ")
    printLinkList(res)
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