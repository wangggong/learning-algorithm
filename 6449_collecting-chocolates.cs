/*
 * @lc app=leetcode.cn id=collecting-chocolates lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6449] 收集巧克力
 *
 * https://leetcode.cn/problems/collecting-chocolates/description/
 *
 * algorithms
 * Medium (25.97%)
 * Total Accepted:    1.5K
 * Total Submissions: 5.2K
 * Testcase Example:  '[20,1,15]\n5'
 *
 * 给你一个长度为 n 、下标从 0 开始的整数数组 nums ，表示收集不同巧克力的成本。每个巧克力都对应一个不同的类型，最初，位于下标 i
 * 的巧克力就对应第 i 个类型。
 * 
 * 在一步操作中，你可以用成本 x 执行下述行为：
 * 
 * 
 * 同时对于所有下标 0 <= i < n - 1 进行以下操作， 将下标 i 处的巧克力的类型更改为下标 (i + 1) 处的巧克力对应的类型。如果 i
 * == n - 1 ，则该巧克力的类型将会变更为下标 0 处巧克力对应的类型。
 * 
 * 
 * 假设你可以执行任意次操作，请返回收集所有类型巧克力所需的最小成本。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [20,1,15], x = 5
 * 输出：13
 * 解释：最开始，巧克力的类型分别是 [0,1,2] 。我们可以用成本 1 购买第 1 个类型的巧克力。
 * 接着，我们用成本 5 执行一次操作，巧克力的类型变更为 [2,0,1] 。我们可以用成本 1 购买第 0 个类型的巧克力。
 * 然后，我们用成本 5 执行一次操作，巧克力的类型变更为 [1,2,0] 。我们可以用成本 1 购买第 2 个类型的巧克力。
 * 因此，收集所有类型的巧克力需要的总成本是 (1 + 5 + 1 + 5 + 1) = 13 。可以证明这是一种最优方案。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3], x = 4
 * 输出：6
 * 解释：我们将会按最初的成本收集全部三个类型的巧克力，而不需执行任何操作。因此，收集所有类型的巧克力需要的总成本是 1 + 2 + 3 = 6
 * 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 10^9
 * 1 <= x <= 10^9
 * 
 * 
 */
public class Solution
{
    public long MinCost(int[] nums, int x)
    {
        var ans = long.MaxValue;
        var n = nums.Length;
        var dp = nums.Select(n => (long)n).ToArray();
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < n; j++)
            {
                dp[j] = Math.Min(dp[j], (long)nums[(j + i) % n]);
            }
            ans = Math.Min(ans, (long)x * (long)i + dp.Sum());
        }
        return ans;
    }
}
