package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil { return true }
    slow, fast := head, head
    var pre *ListNode
    var cur *ListNode
    // 找链表中点同时反转前半部分
    for fast != nil && fast.Next != nil {
        cur = slow
        slow, fast = slow.Next, fast.Next.Next
        cur.Next = pre
        pre = cur
    }
    // 比较pre和slow, 注意奇偶问题带来的比较方式
    if fast != nil { //奇数链，slow比pre多一个中间元素
        slow = slow.Next
    } 
    for pre != nil && slow != nil {
        if pre.Val != slow.Val { return false }
            pre, slow = pre.Next, slow.Next  
    }
    
    if pre != nil || slow != nil { return false }
    return true
}

func main() {
    l := &ListNode{1, &ListNode{0, &ListNode{1,nil}}}
    fmt.Print("input: ")
    printLinkList(l)
    //zsh打印若末尾无换行符会显示一个高亮百分号，与程序无关
    fmt.Println("\noutput: ", isPalindrome(l))
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
