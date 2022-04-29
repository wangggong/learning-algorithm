/*
 * @lc app=leetcode.cn id=min-stack lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (58.02%)
 * Total Accepted:    366.2K
 * Total Submissions: 631.2K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 * 实现 MinStack 类:
 *
 *
 * MinStack() 初始化堆栈对象。
 * void push(int val) 将元素val推入堆栈。
 * void pop() 删除堆栈顶部的元素。
 * int top() 获取堆栈顶部的元素。
 * int getMin() 获取堆栈中的最小元素。
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * 输出：
 * [null,null,null,null,-3,null,0,-2]
 *
 * 解释：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31 <= val <= 2^31 - 1
 * pop、top 和 getMin 操作总是在 非空栈 上调用
 * push, pop, top, and getMin最多被调用 3 * 10^4 次
 *
 *
*/
type MinStack struct {
	Arr []int
	Min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	ms.Arr = append(ms.Arr, val)
	min := val
	if len(ms.Min) > 0 && ms.Min[len(ms.Min)-1] < min {
		min = ms.Min[len(ms.Min)-1]
	}
	ms.Min = append(ms.Min, min)
}

func (ms *MinStack) Pop() {
	if len(ms.Arr) == 0 {
		return
	}
	ms.Arr = ms.Arr[:len(ms.Arr)-1]
	ms.Min = ms.Min[:len(ms.Min)-1]
}

func (ms *MinStack) Top() int {
	if len(ms.Arr) == 0 {
		return -1
	}
	return ms.Arr[len(ms.Arr)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.Min) == 0 {
		return -1
	}
	return ms.Min[len(ms.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
