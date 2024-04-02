/*
 * @lc app=leetcode.cn id=range-module lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [715] Range 模块
 *
 * https://leetcode.cn/problems/range-module/description/
 *
 * algorithms
 * Hard (55.68%)
 * Total Accepted:    28.5K
 * Total Submissions: 51.2K
 * Testcase Example:  '["RangeModule","addRange","removeRange","queryRange","queryRange","queryRange"]\n' +
  '[[],[10,20],[14,16],[10,14],[13,15],[16,17]]'
 *
 * Range模块是跟踪数字范围的模块。设计一个数据结构来跟踪表示为 半开区间 的范围并查询它们。
 * 
 * 半开区间 [left, right) 表示所有 left <= x < right 的实数 x 。
 * 
 * 实现 RangeModule 类:
 * 
 * 
 * RangeModule() 初始化数据结构的对象。
 * void addRange(int left, int right) 添加 半开区间 [left,
 * right)，跟踪该区间中的每个实数。添加与当前跟踪的数字部分重叠的区间时，应当添加在区间 [left, right)
 * 中尚未跟踪的任何数字到该区间中。
 * boolean queryRange(int left, int right) 只有在当前正在跟踪区间 [left, right)
 * 中的每一个实数时，才返回 true ，否则返回 false 。
 * void removeRange(int left, int right) 停止跟踪 半开区间 [left, right)
 * 中当前正在跟踪的每个实数。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入
 * ["RangeModule", "addRange", "removeRange", "queryRange", "queryRange",
 * "queryRange"]
 * [[], [10, 20], [14, 16], [10, 14], [13, 15], [16, 17]]
 * 输出
 * [null, null, null, true, false, true]
 * 
 * 解释
 * RangeModule rangeModule = new RangeModule();
 * rangeModule.addRange(10, 20);
 * rangeModule.removeRange(14, 16);
 * rangeModule.queryRange(10, 14); 返回 true （区间 [10, 14) 中的每个数都正在被跟踪）
 * rangeModule.queryRange(13, 15); 返回 false（未跟踪区间 [13, 15) 中像 14, 14.03, 14.17
 * 这样的数字）
 * rangeModule.queryRange(16, 17); 返回 true （尽管执行了删除操作，区间 [16, 17) 中的数字 16
 * 仍然会被跟踪）
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= left < right <= 10^9
 * 在单个测试用例中，对 addRange 、  queryRange 和 removeRange 的调用总数不超过 10^4 次
 * 
 * 
 */

// 动态开点线段树 (懒标记) 模版.
public class SegTree
{
    public SegTree Left;
    public SegTree Right;
    public int L;
    public int R;
    public int Lazy;
    public int Count;

    public SegTree(int l, int r) => (L, R) = (l, r);

    private void Pushdown()
    {
        var mid = (L + R) >> 1;
        Left = Left ?? new SegTree(L, mid);
        Right = Right ?? new SegTree(mid + 1, R);
        if (Lazy is 0) { return; }
        if (Lazy is 1) { (Left.Count, Right.Count) = (mid - L + 1, R - mid); }
        else { (Left.Count, Right.Count) = (0, 0); }
        (Left.Lazy, Right.Lazy, Lazy) = (Lazy, Lazy, 0);
    }

    public void Add(int l, int r, int v)
    {
        if (l <= L && R <= r)
        {
            Count = v is 1
                ? R - L + 1
                : 0;
            Lazy = v;
            return;
        }
        Pushdown();
        var mid = (L + R) >> 1;
        if (l <= mid) { Left.Add(l, r, v); }
        if (mid + 1 <= r) { Right.Add(l, r, v); }
        Count = Left.Count + Right.Count;
    }

    public int Query(int l, int r)
    {
        if (l <= L && R <= r) { return Count; }
        Pushdown();
        var mid = (L + R) >> 1;
        return (l <= mid ? Left.Query(l, r) : 0)
            + (mid + 1 <= r ? Right.Query(l, r) : 0);
    }
}

public class RangeModule
{
    private const int N = (int)1e9;
    private SegTree tr;
    public RangeModule() => tr = new SegTree(1, N);
    public void AddRange(int left, int right) => tr.Add(left, right - 1, 1);
    public bool QueryRange(int left, int right) => tr.Query(left, right - 1) == right - left;
    public void RemoveRange(int left, int right) => tr.Add(left, right - 1, -1);
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * RangeModule obj = new RangeModule();
 * obj.AddRange(left,right);
 * bool param_2 = obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
