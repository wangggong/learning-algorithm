/*
 * @lc app=leetcode.cn id=subarrays-distinct-element-sum-of-squares-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2916] 子数组不同元素数目的平方和 II
 *
 * https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-ii/description/
 *
 * algorithms
 * Hard (37.64%)
 * Total Accepted:    1.6K
 * Total Submissions: 4.3K
 * Testcase Example:  '[1,2,1]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。
 * 
 * 定义 nums 一个子数组的 不同计数 值如下：
 * 
 * 
 * 令 nums[i..j] 表示 nums 中所有下标在 i 到 j 范围内的元素构成的子数组（满足 0 <= i <= j < nums.length
 * ），那么我们称子数组 nums[i..j] 中不同值的数目为 nums[i..j] 的不同计数。
 * 
 * 
 * 请你返回 nums 中所有子数组的 不同计数 的 平方 和。
 * 
 * 由于答案可能会很大，请你将它对 10^9 + 7 取余 后返回。
 * 
 * 子数组指的是一个数组里面一段连续 非空 的元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,1]
 * 输出：15
 * 解释：六个子数组分别为：
 * [1]: 1 个互不相同的元素。
 * [2]: 1 个互不相同的元素。
 * [1]: 1 个互不相同的元素。
 * [1,2]: 2 个互不相同的元素。
 * [2,1]: 2 个互不相同的元素。
 * [1,2,1]: 2 个互不相同的元素。
 * 所有不同计数的平方和为 1^2 + 1^2 + 1^2 + 2^2 + 2^2 + 2^2 = 15 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,2]
 * 输出：3
 * 解释：三个子数组分别为：
 * [2]: 1 个互不相同的元素。
 * [2]: 1 个互不相同的元素。
 * [2,2]: 1 个互不相同的元素。
 * 所有不同计数的平方和为 1^2 + 1^2 + 1^2 = 3 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^5
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-ii/solutions/2502897/yi-bu-bu-ti-shi-ni-si-kao-ben-ti-pythonj-zhhs/
public class Solution
{
    public int SumCounts(int[] nums)
    {
        const long Mod = (long)1e9 + 7;
        const int N = (int)1e5;
        var ans = 0l;
        var cur = 0l;
        var d = new Dictionary<int, int>();
        var tr = new SegTree(1, N + 1);
        foreach (var (n, i) in nums.Select((n, i) => (n, i + 1)))
        {
            d.TryGetValue(n, out var j);
            d[n] = i;
            cur = (cur + 2 * tr.Query(j + 1, i) + (i - j)) % Mod;
            tr.Add(j + 1, i, 1);
            ans = (ans + cur) % Mod;
        }
        return (int)ans;
    }
}

public class SegTree
{
    public SegTree Left;
    public SegTree Right;
    public int L;
    public int R;
    public long Lazy;
    public long Sum;
    private int mid;

    public SegTree(int l, int r) => (L, R, mid) = (l, r, (l + r) >> 1);
    
    private void Pushdown()
    {
        Left ??= new SegTree(L, mid);
        Right ??= new SegTree(mid + 1, R);
        if (Lazy is 0) { return; }
        Left.Sum += Lazy * (mid - L + 1);
        Right.Sum += Lazy * (R - mid);
        Left.Lazy += Lazy;
        Right.Lazy += Lazy;
        Lazy = 0;
    }

    public void Add(int l, int r, long v)
    {
        if (l <= L && R <= r)
        {
            Sum += v * (R - L + 1);
            Lazy += v;
            return;
        }
        Pushdown();
        if (l <= mid) { Left.Add(l, r, v); }
        if (mid + 1 <= r) { Right.Add(l, r, v); }
        Sum = Left.Sum + Right.Sum;
    }

    public long Query(int l, int r)
    {
        if (l <= L && R <= r) { return Sum; }
        Pushdown();
        return (l <= mid ? Left.Query(l, r) : 0)
            + (mid + 1 <= r ? Right.Query(l, r) : 0);
    }
}
