/*
 * @lc app=leetcode.cn id=maximum-balanced-subsequence-sum lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100112] 平衡子序列的最大和
 *
 * https://leetcode.cn/problems/maximum-balanced-subsequence-sum/description/
 *
 * algorithms
 * Hard (23.06%)
 * Total Accepted:    950
 * Total Submissions: 4.1K
 * Testcase Example:  '[3,3,5,6]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。
 * 
 * nums 一个长度为 k 的 子序列 指的是选出 k 个 下标 i0 < i1 < ... < ik-1 ，如果这个子序列满足以下条件，我们说它是
 * 平衡的 ：
 * 
 * 
 * 对于范围 [1, k - 1] 内的所有 j ，nums[ij] - nums[ij-1] >= ij - ij-1 都成立。
 * 
 * 
 * nums 长度为 1 的 子序列 是平衡的。
 * 
 * 请你返回一个整数，表示 nums 平衡 子序列里面的 最大元素和 。
 * 
 * 一个数组的 子序列 指的是从原数组中删除一些元素（也可能一个元素也不删除）后，剩余元素保持相对顺序得到的 非空 新数组。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,3,5,6]
 * 输出：14
 * 解释：这个例子中，选择子序列 [3,5,6] ，下标为 0 ，2 和 3 的元素被选中。
 * nums[2] - nums[0] >= 2 - 0 。
 * nums[3] - nums[2] >= 3 - 2 。
 * 所以，这是一个平衡子序列，且它的和是所有平衡子序列里最大的。
 * 包含下标 1 ，2 和 3 的子序列也是一个平衡的子序列。
 * 最大平衡子序列和为 14 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,-1,-3,8]
 * 输出：13
 * 解释：这个例子中，选择子序列 [5,8] ，下标为 0 和 3 的元素被选中。
 * nums[3] - nums[0] >= 3 - 0 。
 * 所以，这是一个平衡子序列，且它的和是所有平衡子序列里最大的。
 * 最大平衡子序列和为 13 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [-2,-1]
 * 输出：-1
 * 解释：这个例子中，选择子序列 [-1] 。
 * 这是一个平衡子序列，而且它的和是 nums 所有平衡子序列里最大的。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/maximum-balanced-subsequence-sum/solutions/2513121/shu-zhuang-shu-zu-you-hua-dp-by-endlessc-3zf4/
//
// 类似 LIS 的问题, 只不过状态转移变成了 `dp[i] = max{dp[j]} + nums[i]`.
// 这里的关键是要知道 "不止线段树可以维护区间最大值, 树状数组也可以".
public class Solution
{
    private class BIT
    {
        private long[] tr;
        private int n;

        public BIT(int n) => (this.n, tr) = (n, new long[n + 1]);
        private int LowBit(int n) => n & -n;
        public void Update(int k, long value)
        {
            for (; k <= n; k += LowBit(k)) { tr[k] = Math.Max(tr[k], value); }
        }
        public long Query(int k)
        {
            var ans = long.MinValue;
            for (; k > 0; k -= LowBit(k)) { ans = Math.Max(ans, tr[k]); }
            return ans;
        }
    }

    public long MaxBalancedSubsequenceSum(int[] nums)
    {
        var arr = nums
            .Select((n, i) => (long)n - (long)i)
            .Distinct()
            .OrderBy(n => n)
            .ToArray();
        var m = arr.Length;
        var tr = new BIT(m);
        var ans = long.MinValue;
        foreach (var (n, i) in nums.Select((n, i) => (n, i)))
        {
            var v = (long)n - (long)i;
            var (p, q) = (0, m);
            while (p < q)
            {
                var mid = (p + q) >> 1;
                if (arr[mid] >= v) { q = mid; }
                else { p = mid + 1; }
            }
            p++;
            var cur = Math.Max(tr.Query(p), 0) + n;
            ans = Math.Max(ans, cur);
            tr.Update(p, cur);
        }
        return ans;
    }
}
