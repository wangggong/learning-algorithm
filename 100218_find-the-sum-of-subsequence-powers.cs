/*
 * @lc app=leetcode.cn id=find-the-sum-of-subsequence-powers lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100218] 求出所有子序列的能量和
 *
 * https://leetcode.cn/problems/find-the-sum-of-subsequence-powers/description/
 *
 * algorithms
 * Hard (29.41%)
 * Total Accepted:    583
 * Total Submissions: 2K
 * Testcase Example:  '[1,2,3,4]\n3'
 *
 * 给你一个长度为 n 的整数数组 nums 和一个 正 整数 k 。
 * 
 * 一个子序列的 能量 定义为子序列中 任意 两个元素的差值绝对值的 最小值 。
 * 
 * 请你返回 nums 中长度 等于 k 的 所有 子序列的 能量和 。
 * 
 * 由于答案可能会很大，将答案对 10^9 + 7 取余 后返回。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,4], k = 3
 * 
 * 输出：4
 * 
 * 解释：
 * 
 * nums 中总共有 4 个长度为 3 的子序列：[1,2,3] ，[1,3,4] ，[1,2,4] 和 [2,3,4] 。能量和为 |2 - 3| +
 * |3 - 4| + |2 - 1| + |3 - 4| = 4 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,2], k = 2
 * 
 * 输出：0
 * 
 * 解释：
 * 
 * nums 中唯一一个长度为 2 的子序列是 [2,2] 。能量和为 |2 - 2| = 0 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [4,3,-1], k = 2
 * 
 * 输出：10
 * 
 * 解释：
 * 
 * nums 总共有 3 个长度为 2 的子序列：[4,3] ，[4,-1] 和 [3,-1] 。能量和为 |4 - 3| + |4 - (-1)| +
 * |3 - (-1)| = 10 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n == nums.length <= 50
 * -10^8 <= nums[i] <= 10^8 
 * 2 <= k <= n
 * 
 * 
 */

// 参考: https://leetcode.cn/contest/biweekly-contest-127/ranking/, 小羊肖恩的提交
// 排序 + DP. 里面这个 DP 是真的没想到, 该.
public class Solution
{
    public int SumOfPowers(int[] nums, int k)
    {
        const long Mod = (long)1e9 + 7;
        var ans = 0l;
        var n = nums.Length;
        Array.Sort(nums);
        long[] getDp(int[] arr, int d)
        {
            var n = arr.Length;
            var dp = new long[n][];
            for (var i = 0; i < n; i++)
            {
                dp[i] = new long[k];
            }
            dp[0][1] = 1;
            for (var i = 0; i < n; i++)
            {
                for (var j = i + 1; j < n; j++)
                {
                    if (arr[j] - arr[i] < d)
                    {
                        continue;
                    }
                    for (var m = 0; m + 1 < k; m++)
                    {
                        dp[j][m + 1] = (dp[j][m + 1] + dp[i][m]) % Mod;
                    }
                }
            }
            var ans = new long[k];
            for (var i = 0; i < n; i++)
            {
                for (var j = 0; j < k; j++)
                {
                    ans[j] = (ans[j] + dp[i][j]) % Mod;
                }
            }
            return ans;
        }
        for (var i = 0; i < n; i++)
        {
            for (var j = i + 1; j < n; j++)
            {
                var d = nums[j] - nums[i];
                var lefts = getDp(nums[..(i + 1)]
                    .Select(n => nums[i] - n)
                    .Reverse()
                    .ToArray(), d);
                var rights = getDp(nums[j..]
                    .Select(n => n - nums[j])
                    .ToArray(), d + 1);
                for (var m = 1; m < k; m++)
                {
                    ans = (ans + lefts[m] * rights[k - m] % Mod * d % Mod) % Mod;
                }
            }
        }
        return (int)ans;
    }
}
