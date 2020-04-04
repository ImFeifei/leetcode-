package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
// 常规：用额外空间记录节点 这里用map
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return nil }
    m := map[*ListNode]int{}
    for head != nil {
        if _, ok := m[head]; ok{
            return head
        }
        m[head]++
        head = head.Next
    }
    return nil
}
 // 2020.4.2 双指针 空间O(1) 
 // 相遇点距离环首的距离 = 头节点到环首的距离
// func detectCycle(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil { return nil }

//     s_index, f_index := 0, 0
//     slow, fast := head, head
//     for fast != nil && fast.Next != nil {
//         s_index++
//         f_index += 2
//         slow, fast = slow.Next, fast.Next.Next
//         if slow == fast {
//             for head != slow {
//                 head, slow = head.Next, slow.Next
//             }
//             return head
//         }
//     }
//     return nil
// }

func main() {
    l1 := &ListNode{1, nil}
    l2 := &ListNode{0, nil}
    l3 := &ListNode{2, nil}
    l1.Next = l2
    l2.Next = l3
    l3.Next = l2
    fmt.Print("input: 1 -> 0 -> 2链回到0")
    //zsh打印若末尾无换行符会显示一个高亮百分号，与程序无关
    node := detectCycle(l1)
    pos := 0
    for l1 != nil {
        if l1 == node {
            fmt.Println("\noutput: ", pos)
            break
        }
        l1 = l1.Next
        pos++
    }
    
}
