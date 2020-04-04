package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if n == 0 { return head }
    dummy := &ListNode{0, head}
    slow, fast := dummy, dummy
    for i:=0; i < n; i++ {
        fast = fast.Next
    }
    for fast.Next != nil {
        slow, fast = slow.Next, fast.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}

func main() {
    n := 2
    l := &ListNode{1, &ListNode{2, &ListNode{3,&ListNode{4, &ListNode{5,nil}}}}}
    fmt.Print("input: ")
    printLinkList(l)
    fmt.Printf("   n =%d\n", n)
    //zsh打印若末尾无换行符会显示一个高亮百分号，与程序无关
    fmt.Print("output: ")
    printLinkList(removeNthFromEnd(l,n))
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

