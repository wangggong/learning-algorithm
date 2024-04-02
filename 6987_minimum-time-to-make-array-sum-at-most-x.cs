/*
 * @lc app=leetcode.cn id=minimum-time-to-make-array-sum-at-most-x lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6987] 使数组和小于等于 x 的最少时间
 *
 * https://leetcode.cn/problems/minimum-time-to-make-array-sum-at-most-x/description/
 *
 * algorithms
 * Hard (24.75%)
 * Total Accepted:    433
 * Total Submissions: 1.7K
 * Testcase Example:  '[1,2,3]\n[1,2,3]\n4'
 *
 * 给你两个长度相等下标从 0 开始的整数数组 nums1 和 nums2 。每一秒，对于所有下标 0 <= i < nums1.length
 * ，nums1[i] 的值都增加 nums2[i] 。操作 完成后 ，你可以进行如下操作：
 * 
 * 
 * 选择任一满足 0 <= i < nums1.length 的下标 i ，并使 nums1[i] = 0 。
 * 
 * 
 * 同时给你一个整数 x 。
 * 
 * 请你返回使 nums1 中所有元素之和 小于等于 x 所需要的 最少 时间，如果无法实现，那么返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums1 = [1,2,3], nums2 = [1,2,3], x = 4
 * 输出：3
 * 解释：
 * 第 1 秒，我们对 i = 0 进行操作，得到 nums1 = [0,2+2,3+3] = [0,4,6] 。
 * 第 2 秒，我们对 i = 1 进行操作，得到 nums1 = [0+1,0,6+3] = [1,0,9] 。
 * 第 3 秒，我们对 i = 2 进行操作，得到 nums1 = [1+1,0+2,0] = [2,2,0] 。
 * 现在 nums1 的和为 4 。不存在更少次数的操作，所以我们返回 3 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums1 = [1,2,3], nums2 = [3,3,3], x = 4
 * 输出：-1
 * 解释：不管如何操作，nums1 的和总是会超过 x 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums1.length <= 10^3
 * 1 <= nums1[i] <= 10^3
 * 0 <= nums2[i] <= 10^3
 * nums1.length == nums2.length
 * 0 <= x <= 10^6
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-time-to-make-array-sum-at-most-x/solution/jiao-ni-yi-bu-bu-si-kao-ben-ti-by-endles-2eho/
//
// 从贪心策略引申出的背包问题. 有点意思.
public class Solution
{
    public int MinimumTime(IList<int> nums1, IList<int> nums2, int target)
    {
        var n = nums1.Count();
        var dp = new int[n + 1];
        foreach (var (n1, n2) in nums1.Zip(nums2)
            .OrderBy(x => x.Item2))
        {
            for (var i = n; i > 0; i--)
            {
                dp[i] = Math.Max(dp[i], dp[i - 1] + (n1 + i * n2));
            }
        }
        var (s1, s2) = (nums1.Sum(), nums2.Sum());
        return dp.Select((v, t) => (v, t))
            .Where(x => (s1 + x.t * s2) - x.v <= target)
            .FirstOrDefault((-1, -1))
            .Item2;
    }
}
