/*
 * @lc app=leetcode.cn id=longest-arithmetic-subsequence lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1027] 最长等差数列
 *
 * https://leetcode.cn/problems/longest-arithmetic-subsequence/description/
 *
 * algorithms
 * Medium (44.07%)
 * Total Accepted:    22.3K
 * Total Submissions: 50.4K
 * Testcase Example:  '[3,6,9,12]'
 *
 * 给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。
 * 
 * 回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik] ，且 0 <= i1 < i2 < ...
 * < ik <= nums.length - 1。并且如果 seq[i+1] - seq[i]( 0 <= i < seq.length - 1)
 * 的值都相同，那么序列 seq 是等差的。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,6,9,12]
 * 输出：4
 * 解释： 
 * 整个数组是公差为 3 的等差数列。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [9,4,7,2,10]
 * 输出：3
 * 解释：
 * 最长的等差子序列是 [4,7,10]。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [20,1,15,3,10,5,8]
 * 输出：4
 * 解释：
 * 最长的等差子序列是 [20,15,10,5]。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= nums.length <= 1000
 * 0 <= nums[i] <= 500
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     private const int N = 500;
 * 
 *     public int LongestArithSeqLength(int[] nums)
 *     {
 *         var n = nums.Length;
 *         var ans = 0;
 *         for (var d = -N * 2; d <= N * 2; d++)
 *         {
 *             var memo = new Dictionary<int, int>();
 *             foreach (var num in nums)
 *             {
 *                 memo[num] = Math.Max(memo.ContainsKey(num) ? memo[num] : 0,
 *                     (memo.ContainsKey(num - d) ? memo[num - d] : 0) + 1);
 *                 ans = Math.Max(ans, memo[num]);
 *             }
 *         }
 *         return ans;
 *     }
 * }
 */

public class Solution
{
    private const int N = 500;

    public int LongestArithSeqLength(int[] nums)
    {
        var n = nums.Length;
        var ans = 0;
        var dp = new int[n][];
        for (var i = 0; i < n; i++)
        {
            dp[i] = new int[N * 2 + 1];
        }
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j <= N * 2; j++)
            {
                dp[i][j] = 1;
            }
            for (var j = 0; j < i; j++)
            {
                var d = nums[i] - nums[j] + N;
                dp[i][d] = Math.Max(dp[i][d], dp[j][d] + 1);
                ans = Math.Max(ans, dp[i][d]);
            }
        }
        return ans;
    }
}
