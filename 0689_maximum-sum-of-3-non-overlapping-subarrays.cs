/*
 * @lc app=leetcode.cn id=maximum-sum-of-3-non-overlapping-subarrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [689] 三个无重叠子数组的最大和
 *
 * https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays/description/
 *
 * algorithms
 * Hard (56.08%)
 * Total Accepted:    23.5K
 * Total Submissions: 41.8K
 * Testcase Example:  '[1,2,1,2,6,7,5,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，找出三个长度为 k 、互不重叠、且全部数字和（3 * k 项）最大的子数组，并返回这三个子数组。
 * 
 * 以下标的数组形式返回结果，数组中的每一项分别指示每个子数组的起始位置（下标从 0 开始）。如果有多个结果，返回字典序最小的一个。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,1,2,6,7,5,1], k = 2
 * 输出：[0,3,5]
 * 解释：子数组 [1, 2], [2, 6], [7, 5] 对应的起始下标为 [0, 3, 5]。
 * 也可以取 [2, 1], 但是结果 [1, 3, 5] 在字典序上更大。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,1,2,1,2,1,2,1], k = 2
 * 输出：[0,2,4]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 2 * 10^4
 * 1 <= nums[i] < 2^16
 * 1 <= k <= floor(nums.length / 3)
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays/solutions/1145272/gong-shui-san-xie-jie-he-qian-zhui-he-de-ancx/?envType=daily-question&envId=2023-11-19
public class Solution
{
    public int[] MaxSumOfThreeSubarrays(int[] nums, int k)
    {
        var n = nums.Length;
        var S = new int[n + 1];
        for (var i = 0; i < n; i++) { S[i + 1] = S[i] + nums[n - 1 - i]; }
        var dp = new int[n + 1][];
        for (var i = 0; i <= n; i++) { dp[i] = new int[4]; }
        for (var i = k; i <= n; i++)
        {
            for (var j = 1; j <= 3; j++)
            {
                dp[i][j] = Math.Max(dp[i - 1][j],
                    dp[i - k][j - 1] + (S[i] - S[i - k]));
            }
        }
        var ans = new int[3];
        for (var (i, j, idx) = (n, 3, 0); j > 0; j--)
        {
            for (; dp[i - 1][j] > dp[i - k][j - 1] + (S[i] - S[i - k]); i--) { }
            ans[idx] = n - i;
            idx++;
            i -= k;
        }
        return ans;
    }
}
