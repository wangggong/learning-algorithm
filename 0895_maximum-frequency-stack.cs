/*
 * @lc app=leetcode.cn id=maximum-frequency-stack lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [895] 最大频率栈
 *
 * https://leetcode.cn/problems/maximum-frequency-stack/description/
 *
 * algorithms
 * Hard (59.36%)
 * Total Accepted:    24.5K
 * Total Submissions: 38.3K
 * Testcase Example:  '["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"]\n' +
  '[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]'
 *
 * 设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。
 * 
 * 实现 FreqStack 类:
 * 
 * 
 * FreqStack() 构造一个空的堆栈。
 * void push(int val) 将一个整数 val 压入栈顶。
 * int pop() 删除并返回堆栈中出现频率最高的元素。
 * 
 * 如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：
 * 
 * ["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
 * [[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
 * 输出：[null,null,null,null,null,null,null,5,7,5,4]
 * 解释：
 * FreqStack = new FreqStack();
 * freqStack.push (5);//堆栈为 [5]
 * freqStack.push (7);//堆栈是 [5,7]
 * freqStack.push (5);//堆栈是 [5,7,5]
 * freqStack.push (7);//堆栈是 [5,7,5,7]
 * freqStack.push (4);//堆栈是 [5,7,5,7,4]
 * freqStack.push (5);//堆栈是 [5,7,5,7,4,5]
 * freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
 * freqStack.pop ();//返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
 * freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
 * freqStack.pop ();//返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= val <= 10^9
 * push 和 pop 的操作数不大于 2 * 10^4。
 * 输入保证在调用 pop 之前堆栈中至少有一个元素。
 * 
 * 
 */
public class FreqStack
{
    private Dictionary<int, int> Count;
    private List<List<int>> StackList;

    public FreqStack()
    {
        Count = new();
        StackList = new();
        StackList.Add(new());
    }

    public void Push(int val)
    {
        var c = (Count.ContainsKey(val) ? Count[val] : 0) + 1;
        if (StackList.Count <= c) { StackList.Add(new()); }
        Count[val] = c;
        StackList[c].Add(val);
    }

    public int Pop()
    {
        var n = StackList.Count - 1;
        // assert n > 0;
        var m = StackList[n].Count - 1;
        var val = StackList[n][m];
        Count[val]--;
        StackList[n].RemoveAt(m);
        if (StackList[n].Count == 0) { StackList.RemoveAt(n); }
        return val;
    }
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * FreqStack obj = new FreqStack();
 * obj.Push(val);
 * int param_2 = obj.Pop();
 */
