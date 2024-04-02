/*
 * @lc app=leetcode.cn id=continuous-subarrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6911] 不间断子数组
 *
 * https://leetcode.cn/problems/continuous-subarrays/description/
 *
 * algorithms
 * Medium (32.25%)
 * Total Accepted:    1.3K
 * Total Submissions: 4K
 * Testcase Example:  '[5,4,2,4]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。nums 的一个子数组如果满足以下条件，那么它是 不间断 的：
 * 
 * 
 * i，i + 1 ，...，j  表示子数组中的下标。对于所有满足 i <= i1, i2 <= j 的下标对，都有 0 <= |nums[i1] -
 * nums[i2]| <= 2 。
 * 
 * 
 * 请你返回 不间断 子数组的总数目。
 * 
 * 子数组是一个数组中一段连续 非空 的元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [5,4,2,4]
 * 输出：8
 * 解释：
 * 大小为 1 的不间断子数组：[5], [4], [2], [4] 。
 * 大小为 2 的不间断子数组：[5,4], [4,2], [2,4] 。
 * 大小为 3 的不间断子数组：[4,2,4] 。
 * 没有大小为 4 的不间断子数组。
 * 不间断子数组的总数目为 4 + 3 + 1 = 8 。
 * 除了这些以外，没有别的不间断子数组。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：6
 * 解释：
 * 大小为 1 的不间断子数组：[1], [2], [3] 。
 * 大小为 2 的不间断子数组：[1,2], [2,3] 。
 * 大小为 3 的不间断子数组：[1,2,3] 。
 * 不间断子数组的总数目为 3 + 2 + 1 = 6 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */
public class Solution
{
    private const int Threshold = 2;

    /*
     * public long ContinuousSubarrays(int[] nums)
     * {
     *     var ans = 0L;
     *     var minQ = new PriorityQueue<int, int>();
     *     var maxQ = new PriorityQueue<int, int>();
     *     for (var (p, q, n) = (0, 0, nums.Length); p < n; p++)
     *     {
     *         for (; minQ.Count > 0 && minQ.Peek() < p; minQ.Dequeue()) { }
     *         for (; maxQ.Count > 0 && maxQ.Peek() < p; maxQ.Dequeue()) { }
     *         for (; q < n; q++)
     *         {
     *             minQ.Enqueue(q, nums[q]);
     *             maxQ.Enqueue(q, -nums[q]);
     *             if (nums[maxQ.Peek()] - nums[minQ.Peek()] > Threshold)
     *             {
     *                 break;
     *             }
     *         }
     *         ans += (long)(q - p);
     *     }
     *     return ans;
     * }
     */

    // 滑动窗口
    public long ContinuousSubarrays(int[] nums)
    {
        var ans = 0L;
        var p = 0;
        var d = new Dictionary<int, int>();
        foreach (var (v, q) in nums.Select((x, i) => (x, i)))
        {
            d.TryGetValue(v, out var c);
            d[v] = c + 1;
            for (; d.Keys.Max() - d.Keys.Min() > Threshold; p++)
            {
                d[nums[p]]--;
                if (d[nums[p]] is 0)
                {
                    d.Remove(nums[p]);
                }
            }
            ans += q - p + 1;
        }
        return ans;
    }
}
