/*
 * @lc app=leetcode.cn id=split-array-into-maximum-number-of-subarrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100019] 将数组分割成最多数目的子数组
 *
 * https://leetcode.cn/problems/split-array-into-maximum-number-of-subarrays/description/
 *
 * algorithms
 * Medium (41.50%)
 * Total Accepted:    1.5K
 * Total Submissions: 3.6K
 * Testcase Example:  '[1,0,2,0,1,2]'
 *
 * 给你一个只包含 非负 整数的数组 nums 。
 * 
 * 我们定义满足 l <= r 的子数组 nums[l..r] 的分数为 nums[l] AND nums[l + 1] AND ... AND
 * nums[r] ，其中 AND 是按位与运算。
 * 
 * 请你将数组分割成一个或者更多子数组，满足：
 * 
 * 
 * 每个 元素都 只 属于一个子数组。
 * 子数组分数之和尽可能 小 。
 * 
 * 
 * 请你在满足以上要求的条件下，返回 最多 可以得到多少个子数组。
 * 
 * 一个 子数组 是一个数组中一段连续的元素。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,0,2,0,1,2]
 * 输出：3
 * 解释：我们可以将数组分割成以下子数组：
 * - [1,0] 。子数组分数为 1 AND 0 = 0 。
 * - [2,0] 。子数组分数为 2 AND 0 = 0 。
 * - [1,2] 。子数组分数为 1 AND 2 = 0 。
 * 分数之和为 0 + 0 + 0 = 0 ，是我们可以得到的最小分数之和。
 * 在分数之和为 0 的前提下，最多可以将数组分割成 3 个子数组。所以返回 3 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,7,1,3]
 * 输出：1
 * 解释：我们可以将数组分割成一个子数组：[5,7,1,3] ，分数为 1 ，这是可以得到的最小总分数。
 * 在总分数为 1 的前提下，最多可以将数组分割成 1 个子数组。所以返回 1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^6
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int MaxSubarrays(int[] nums)
 *     {
 *         var ans = 0;
 *         for (var (p, q, n) = (0, 0, nums.Length); p < n; p = q)
 *         {
 *             var cur = nums[p];
 *             for (q = p + 1; q < n && cur != 0; q++) { cur &= nums[q]; }
 *             if (cur is 0) { ans++; }
 *         }
 *         return Math.Max(ans, 1);
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/split-array-into-maximum-number-of-subarrays/solutions/2464645/li-yong-and-xing-zhi-yi-ci-bian-li-pytho-p3bj/
//
// 直接用 `-1` 当作二进制全 "一".
public class Solution
{
    public int MaxSubarrays(int[] nums)
    {
        var (ans, cur) = (0, -1);
        foreach (var n in nums)
        {
            cur &= n;
            if (cur is 0)
            {
                ans++;
                cur = -1;
            }
        }
        return Math.Max(ans, 1);
    }
}
