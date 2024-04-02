/*
 * @lc app=leetcode.cn id=special-permutations lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6893] 特别的排列
 *
 * https://leetcode.cn/problems/special-permutations/description/
 *
 * algorithms
 * Medium (20.71%)
 * Total Accepted:    1.2K
 * Total Submissions: 5.7K
 * Testcase Example:  '[2,3,6]'
 *
 * 给你一个下标从 0 开始的整数数组 nums ，它包含 n 个 互不相同 的正整数。如果 nums
 * 的一个排列满足以下条件，我们称它是一个特别的排列：
 * 
 * 
 * 对于 0 <= i < n - 1 的下标 i ，要么 nums[i] % nums[i+1] == 0 ，要么 nums[i+1] % nums[i]
 * == 0 。
 * 
 * 
 * 请你返回特别排列的总数目，由于答案可能很大，请将它对 10^9 + 7 取余 后返回。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,3,6]
 * 输出：2
 * 解释：[3,6,2] 和 [2,6,3] 是 nums 两个特别的排列。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [1,4,3]
 * 输出：2
 * 解释：[3,1,4] 和 [4,1,3] 是 nums 两个特别的排列。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= nums.length <= 14
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */
public class Solution
{
    private const long Mod = (long)1e9 + 7;

    public int SpecialPerm(int[] nums)
    {
        var n = nums.Length;
        var G = Enumerable
            .Range(0, n)
            .Select(i => Enumerable
                .Range(0, n)
                .Where(j => j != i
                    && nums[i] % nums[j] is 0 || nums[j] % nums[i] is 0)
                .ToArray())
            .ToArray();
        var memo = new Dictionary<(int, int), long>();
        long dfs(int u, int mask)
        {
            var key = (u, mask);
            if (!memo.ContainsKey(key))
            {
                memo[key] = mask + 1 == 1 << n
                    ? 1
                    : G[u]
                        .Where(v => (mask & (1 << v)) == 0)
                        .Select(v => dfs(v, mask | 1 << v))
                        .Sum();
            }
            return memo[key];
        }
        return (int)Enumerable
            .Range(0, n)
            .Select(u => dfs(u, 1 << u))
            .Aggregate((x, y) => (x + y) % Mod);
    }
}
