package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
// 快慢指针 空间O(1)
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil { return false }
    slow, fast := head, head
    for fast != nil && fast.Next != nil{
        slow, fast = slow.Next, fast.Next.Next
        if slow == fast { return true}
    }
    return false
}

func main() {
    l1 := &ListNode{1, nil}
    l2 := &ListNode{0, nil}
    l3 := &ListNode{2, nil}
    l1.Next = l2
    l2.Next = l3
    l3.Next = l2
    fmt.Print("input: 1 -> 0 -> 2链回到0")
    //zsh打印若末尾无换行符会显示一个高亮百分号，与程序无关
    fmt.Println("\noutput: ", hasCycle(l1))
}

