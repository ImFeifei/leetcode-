package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
 // solution 2: 迭代归并,空间复杂度O(1)
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head }
    length := 0
    dummy := &ListNode{Next: head}
    p := head
    for p != nil {
        p = p.Next
        length++
    }  
    interval := 1
    p = dummy
    left, right, next_left := head, head, head
    for interval < length {
        p = dummy
        next_left = dummy.Next
        for next_left != nil {
            left = next_left
            right = cut(left, interval)
            next_left = cut(right, interval)
            // p即为最终得到的链表，一次次的向尾部追加完善
            p.Next = merge(left, right)
            for p.Next != nil {
                p = p.Next
            }
        }
        
        interval *= 2
    } 
    return dummy.Next
}
// 从前向后数interval个节点断开链表, 并返回链表切掉前interval个节点后的第一个节点
func cut(head *ListNode, interval int)*ListNode {
    //initial := head
    for i := 0; i < interval; i++ {
        if head != nil {
            if i == interval - 1 {
                prev := head
                head = head.Next
                prev.Next = nil
            } else {
                head = head.Next
            }  
        } else {
            return head
        }
    }  
    // fmt.Println("after cut:")
    // printLinkList(initial)
    // fmt.Print("\t")
    // printLinkList(head)
    // fmt.Print("\n")
    return head
}
// 合并两个有序链表
func merge(h1, h2 *ListNode) *ListNode{
    if h1 == nil { return h2 }
    if h2 == nil { return h1 }
    dummy := &ListNode{Next: h1}
    pre := dummy
    for h1 != nil && h2 != nil{
        if h1.Val > h2.Val {
            tail := h2.Next
            h2.Next = nil
            pre.Next = h2
            pre.Next.Next = h1
            pre = h2
            h2 = tail
        } else {
            pre = h1
            h1 = h1.Next
        }
    }
    if h2 != nil {
        pre.Next = h2
    }
    return dummy.Next
}

// solution 1: 归并，但递归空间复杂度O(logn)
// func sortList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil { return head }
//     // 快慢指针找链表中点
//     slow, fast := head, head
//     for fast.Next != nil && fast.Next.Next != nil {
//         slow = slow.Next
//         fast = fast.Next.Next
//     }
//     left, right := head, slow.Next
//     slow.Next = nil
//     return merge(sortList(left), sortList(right))
// }

// func merge(h1, h2 *ListNode) *ListNode{
//     if h1 == nil { return h2 }
//     if h2 == nil { return h1 }
//     dummy := &ListNode{Next: h1}
//     pre := dummy
//     for h1 != nil && h2 != nil{
//         if h1.Val > h2.Val {
//             tail := h2.Next
//             h2.Next = nil
//             pre.Next = h2
//             pre.Next.Next = h1
//             pre = h2
//             h2 = tail
//         } else {
//             pre = h1
//             h1 = h1.Next
//         }
//     }
//     if h2 != nil {
//         pre.Next = h2
//     }
//     return dummy.Next
// }
func main() {
    l1 := &ListNode{-1, &ListNode{5, &ListNode{3,&ListNode{4,&ListNode{0,nil}}}}}
    //l2 := &ListNode{1, &ListNode{3, &ListNode{4,nil}}}
    fmt.Print("input: ")
    printLinkList(l1)
    //zsh打印若末尾无换行符会显示一个高亮百分号，与程序无关
    fmt.Print("\noutput: ")
    printLinkList(sortList(l1))

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
