/*
 * @lc app=leetcode.cn id=minimum-cost-to-make-array-equalindromic lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100151] 使数组成为等数数组的最小代价
 *
 * https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/description/
 *
 * algorithms
 * Medium (18.67%)
 * Total Accepted:    2.2K
 * Total Submissions: 12K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你一个长度为 n 下标从 0 开始的整数数组 nums 。
 * 
 * 你可以对 nums 执行特殊操作 任意次 （也可以 0 次）。每一次特殊操作中，你需要 按顺序 执行以下步骤：
 * 
 * 
 * 从范围 [0, n - 1] 里选择一个下标 i 和一个 正 整数 x 。
 * 将 |nums[i] - x| 添加到总代价里。
 * 将 nums[i] 变为 x 。
 * 
 * 
 * 如果一个正整数正着读和反着读都相同，那么我们称这个数是 回文数 。比方说，121 ，2552 和 65756 都是回文数，但是 24 ，46 ，235
 * 都不是回文数。
 * 
 * 如果一个数组中的所有元素都等于一个整数 y ，且 y 是一个小于 10^9 的 回文数 ，那么我们称这个数组是一个 等数数组 。
 * 
 * 请你返回一个整数，表示执行任意次特殊操作后使 nums 成为 等数数组 的 最小 总代价。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,4,5]
 * 输出：6
 * 解释：我们可以将数组中所有元素变为回文数 3 得到等数数组，数组变成 [3,3,3,3,3] 需要执行 4 次特殊操作，代价为 |1 - 3| + |2
 * - 3| + |4 - 3| + |5 - 3| = 6 。
 * 将所有元素变为其他回文数的总代价都大于 6 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [10,12,13,14,15]
 * 输出：11
 * 解释：我们可以将数组中所有元素变为回文数 11 得到等数数组，数组变成 [11,11,11,11,11] 需要执行 5 次特殊操作，代价为 |10 -
 * 11| + |12 - 11| + |13 - 11| + |14 - 11| + |15 - 11| = 11 。
 * 将所有元素变为其他回文数的总代价都大于 11 。
 * 
 * 
 * 示例 3 ：
 * 
 * 
 * 输入：nums = [22,33,22,33,22]
 * 输出：22
 * 解释：我们可以将数组中所有元素变为回文数 22 得到等数数组，数组变为 [22,22,22,22,22] 需要执行 2 次特殊操作，代价为 |33 -
 * 22| + |33 - 22| = 22 。
 * 将所有元素变为其他回文数的总代价都大于 22 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/solutions/2569308/yu-chu-li-hui-wen-shu-zhong-wei-shu-tan-7j0zy/
//
// 中位数贪心 ---- 但这里的关键是枚举回文数的模版.
public class Solution
{
    private const long N = (long)1e9;
    private static List<long> palindromes = new();

    static Solution()
    {
        for (var b = 1l; b * b <= N; b *= 10)
        {
            for (var i = b; i < b * 10; i++)
            {
                var (p, s, k) = (i / 10, 0l, i);
                for (; k > 0; k /= 10) { (p, s) = (p * 10, s * 10 + k % 10); }
                palindromes.Add(p + s);
            }
            if (b * b * 100 > N) { continue; }
            for (var i = b; i < b * 10; i++)
            {
                var (p, s, k) = (i, 0l, i);
                for (; k > 0; k /= 10) { (p, s) = (p * 10, s * 10 + k % 10); }
                palindromes.Add(p + s);
            }
        }
        palindromes.Add(N + 1); // 1e9 + 1
    }

    public long MinimumCost(int[] nums)
    {
        Array.Sort(nums);
        var target = nums[nums.Length >> 1];
        var p = 0;
        for (; palindromes[p] < (long)target; p++) { }
        long total(long v) => nums.Select(n => Math.Abs((long)n - v)).Sum();
        var ans = total(palindromes[p]);
        if (p > 0) { ans = Math.Min(ans, total(palindromes[p - 1])); }
        if (p + 1 < palindromes.Count()) { ans = Math.Min(ans, total(palindromes[p + 1])); }
        return ans;
    }
}
