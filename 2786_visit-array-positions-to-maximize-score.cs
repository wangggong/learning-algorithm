/*
 * @lc app=leetcode.cn id=visit-array-positions-to-maximize-score lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2786] 访问数组中的位置使分数最大
 *
 * https://leetcode.cn/problems/visit-array-positions-to-maximize-score/description/
 *
 * algorithms
 * Medium (37.33%)
 * Total Accepted:    2.3K
 * Total Submissions: 6.1K
 * Testcase Example:  '[2,3,6,1,9,2]\n5'
 *
 * 给你一个下标从 0 开始的整数数组 nums 和一个正整数 x 。
 * 
 * 你 一开始 在数组的位置 0 处，你可以按照下述规则访问数组中的其他位置：
 * 
 * 
 * 如果你当前在位置 i ，那么你可以移动到满足 i < j 的 任意 位置 j 。
 * 对于你访问的位置 i ，你可以获得分数 nums[i] 。
 * 如果你从位置 i 移动到位置 j 且 nums[i] 和 nums[j] 的 奇偶性 不同，那么你将失去分数 x 。
 * 
 * 
 * 请你返回你能得到的 最大 得分之和。
 * 
 * 注意 ，你一开始的分数为 nums[0] 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,3,6,1,9,2], x = 5
 * 输出：13
 * 解释：我们可以按顺序访问数组中的位置：0 -> 2 -> 3 -> 4 。
 * 对应位置的值为 2 ，6 ，1 和 9 。因为 6 和 1 的奇偶性不同，所以下标从 2 -> 3 让你失去 x = 5 分。
 * 总得分为：2 + 6 + 1 + 9 - 5 = 13 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [2,4,6,8], x = 3
 * 输出：20
 * 解释：数组中的所有元素奇偶性都一样，所以我们可以将每个元素都访问一次，而且不会失去任何分数。
 * 总得分为：2 + 4 + 6 + 8 = 20 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= nums.length <= 10^5
 * 1 <= nums[i], x <= 10^6
 * 
 * 
 */
public class Solution
{
    private const long Maxn = 0x3f3f3f3f;

    public long MaxScore(int[] nums, int x)
    {
        var dp = new long[2];
        Array.Fill(dp, -Maxn);
        dp[nums[0] % 2] = nums[0];
        foreach (var n in nums[1 ..])
        {
            var curr = (long)n + Math.Max(dp[n % 2], dp[(n - 1) % 2] - x);
            dp[n % 2] = Math.Max(dp[n % 2], curr);
        }
        return dp.Max();
    }
}
