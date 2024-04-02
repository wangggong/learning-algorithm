/*
 * @lc app=leetcode.cn id=maximize-score-after-n-operations lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1799] N 次操作后的最大分数和
 *
 * https://leetcode.cn/problems/maximize-score-after-n-operations/description/
 *
 * algorithms
 * Hard (54.32%)
 * Total Accepted:    6.4K
 * Total Submissions: 10.6K
 * Testcase Example:  '[1,2]'
 *
 * 给你 nums ，它是一个大小为 2 * n 的正整数数组。你必须对这个数组执行 n 次操作。
 * 
 * 在第 i 次操作时（操作编号从 1 开始），你需要：
 * 
 * 
 * 选择两个元素 x 和 y 。
 * 获得分数 i * gcd(x, y) 。
 * 将 x 和 y 从 nums 中删除。
 * 
 * 
 * 请你返回 n 次操作后你能获得的分数和最大为多少。
 * 
 * 函数 gcd(x, y) 是 x 和 y 的最大公约数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [1,2]
 * 输出：1
 * 解释：最优操作是：
 * (1 * gcd(1, 2)) = 1
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [3,4,6,8]
 * 输出：11
 * 解释：最优操作是：
 * (1 * gcd(3, 6)) + (2 * gcd(4, 8)) = 3 + 8 = 11
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums = [1,2,3,4,5,6]
 * 输出：14
 * 解释：最优操作是：
 * (1 * gcd(1, 5)) + (2 * gcd(2, 4)) + (3 * gcd(3, 6)) = 1 + 4 + 9 = 14
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 7
 * nums.length == 2 * n
 * 1 <= nums[i] <= 10^6
 * 
 * 
 */
public class Solution
{
    public int MaxScore(int[] nums)
    {
        int gcd(int x, int y) => y == 0 ? x : gcd(y, x % y);
        var n = nums.Length / 2;
        var gcds = new int[n * 2][];
        for (var i = 0; i < n * 2; i++)
        {
            gcds[i] = new int[n * 2];
        }
        for (var i = 0; i < n * 2; i++)
        {
            for (var j = i + 1; j < n * 2; j++)
            {
                gcds[i][j] = gcd(nums[i], nums[j]);
            }
        }
        var states = 1 << (n * 2);
        var dp = new int[n + 1][];
        for (var i = 0; i <= n; i++)
        {
            dp[i] = new int[states];
        }
        for (var k = 1; k <= n; k++)
        {
            for (var s = 0; s < states; s++)
            {
                var available = new List<int>();
                for (var i = 0; i < n * 2; i++)
                {
                    if ((s & (1 << i)) == 0)
                    {
                        available.Add(i);
                    }
                }
                var m = available.Count();
                for (var i = 0; i < m; i++)
                {
                    for (var j = i + 1; j < m; j++)
                    {
                        var ns = s | (1 << available[i]) | (1 << available[j]);
                        dp[k][ns] = Math.Max(dp[k][ns], dp[k - 1][s] + k * gcds[available[i]][available[j]]);
                    }
                }
            }
        }
        return dp[n][states - 1];
    }
}
