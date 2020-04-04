package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil { return l2 }
    if l2 == nil { return l1 }

    addFlag := 0
    p, q := l1, l2

    for p.Next != nil && q.Next != nil { //加在l1 (p)上
        p.Val, addFlag = (p.Val + q.Val + addFlag)%10, (p.Val + q.Val + addFlag)/10
        p, q = p.Next, q.Next
    }
    p.Val, addFlag = (p.Val + q.Val + addFlag)%10, (p.Val + q.Val + addFlag)/10
    if p.Next == nil && q.Next == nil {
        if addFlag == 1 {
            p.Next = &ListNode{1, nil}
        }
    } else if p.Next != nil && q.Next == nil {
        for p.Next != nil {
            p = p.Next
            p.Val, addFlag = (p.Val + addFlag)%10, (p.Val + addFlag)/10
        }
        if addFlag == 1 {
            p.Next = &ListNode{1, nil}
        }
    } else if p.Next == nil && q.Next != nil {
        p.Next = q.Next
        for p.Next != nil {
            p = p.Next
            p.Val, addFlag = (p.Val + addFlag)%10, (p.Val + addFlag)/10
        }
        if addFlag == 1 {
            p.Next = &ListNode{1, nil}
        }
    }
    return l1
}

func main() {
    // l := &ListNode{1, &ListNode{0, &ListNode{1,nil}}}
    l1 := &ListNode{5, nil}
    l2 := &ListNode{5, nil}
    fmt.Print("input: l1: ")
    printLinkList(l1)
    fmt.Print("    l2:")
    printLinkList(l2)
    //zsh打印若末尾无换行符会显示一个高亮百分号，与程序无关
    fmt.Print("\noutput: ")
    printLinkList(addTwoNumbers(l1,l2))
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
