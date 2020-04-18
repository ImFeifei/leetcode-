package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
    Pre *ListNode
}

func main() {
	// m := map[int]*ListNode{1:&ListNode{1,nil}, 2:&ListNode{2,nil}, 3:&ListNode{3,nil}}
	// delete(m, 2)
	// fmt.Println(m)
	dummy := &ListNode{-1, nil, nil}
	a := &ListNode{2, nil, nil}
	b := &ListNode{4, nil, nil}
	c := &ListNode{3, nil, nil}
	a.Next = b
	b.Next = c
	b.Pre = a
	c.Pre = b
	dummy.Next = a
	a.Pre = dummy
	tail := c


	last := tail
    tail = a
    last.Next = a
    a.Pre.Next = a.Next
    a.Next.Pre = a.Pre
    a.Next = nil 
    if last.Pre == a {
        last.Pre = a.Pre
    }
    a.Pre = last  

    p := dummy
    for p != nil {
        fmt.Printf("%d ", p.Val)
        p = p.Next
    }
    fmt.Printf("\n")  

}
