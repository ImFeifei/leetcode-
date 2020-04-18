package main
import "fmt"

type MinStack struct {
    Stack []int
    Helper []int   // 辅助栈，单调递减，非严格，用于取栈顶最小值
}


/** initialize your data structure here. */
func Constructor() MinStack {
    m := MinStack{
        Stack: make([]int,0),
        Helper: make([]int,0),
    }
    return m
}


func (this *MinStack) Push(x int)  {
    if len(this.Helper) == 0 {
        this.Helper = append(this.Helper, x)
    } else {
        if x <= this.Helper[len(this.Helper)-1] {
           this.Helper = append(this.Helper, x)
        }
    }
    this.Stack = append(this.Stack, x)    
}


func (this *MinStack) Pop()  {
    val := this.Stack[len(this.Stack)-1]
    this.Stack = this.Stack[:len(this.Stack)-1]
    if this.Helper[len(this.Helper)-1] == val {
        this.Helper = this.Helper[:len(this.Helper)-1]
    }
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.Helper[len(this.Helper)-1]
}

func main() {
    
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */