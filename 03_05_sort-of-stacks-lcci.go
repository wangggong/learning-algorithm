/*
 * 栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。
 *
 * 示例1:
 *
 *  输入：
 * ["SortedStack", "push", "push", "peek", "pop", "peek"]
 * [[], [1], [2], [], [], []]
 *  输出：
 * [null,null,null,1,null,2]
 * 示例2:
 *
 *  输入：
 * ["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
 * [[], [], [], [1], [], []]
 *  输出：
 * [null,null,null,null,null,true]
 * 说明:
 *
 * 栈中的元素数目在[0, 5000]范围内。
 */
type SortedStack struct {
	Stack     []int
	TempStack []int
}

func Constructor() SortedStack {
	return SortedStack{}
}

func (this *SortedStack) Push(val int) {
	for len(this.Stack) > 0 && this.Stack[len(this.Stack)-1] < val {
		v := this.Stack[len(this.Stack)-1]
		this.Stack = this.Stack[:len(this.Stack)-1]
		this.TempStack = append(this.TempStack, v)
	}
	this.Stack = append(this.Stack, val)
	for len(this.TempStack) > 0 {
		v := this.TempStack[len(this.TempStack)-1]
		this.TempStack = this.TempStack[:len(this.TempStack)-1]
		this.Stack = append(this.Stack, v)
	}
}

func (this *SortedStack) Pop() {
	if this.IsEmpty() {
		return
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *SortedStack) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Stack[len(this.Stack)-1]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.Stack) == 0
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
