/*
 * @lc app=leetcode.cn id=minimum-size-subarray-in-infinite-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100076] 无限数组的最短子数组
 *
 * https://leetcode.cn/problems/minimum-size-subarray-in-infinite-array/description/
 *
 * algorithms
 * Medium (30.19%)
 * Total Accepted:    1.9K
 * Total Submissions: 6.3K
 * Testcase Example:  '[1,2,3]\n5'
 *
 * 给你一个下标从 0 开始的数组 nums 和一个整数 target 。
 * 
 * 下标从 0 开始的数组 infinite_nums 是通过无限地将 nums 的元素追加到自己之后生成的。
 * 
 * 请你从 infinite_nums 中找出满足 元素和 等于 target 的 最短 子数组，并返回该子数组的长度。如果不存在满足条件的子数组，返回
 * -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3], target = 5
 * 输出：2
 * 解释：在这个例子中 infinite_nums = [1,2,3,1,2,3,1,2,...] 。
 * 区间 [1,2] 内的子数组的元素和等于 target = 5 ，且长度 length = 2 。
 * 可以证明，当元素和等于目标值 target = 5 时，2 是子数组的最短长度。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,1,1,2,3], target = 4
 * 输出：2
 * 解释：在这个例子中 infinite_nums = [1,1,1,2,3,1,1,1,2,3,1,1,...].
 * 区间 [4,5] 内的子数组的元素和等于 target = 4 ，且长度 length = 2 。
 * 可以证明，当元素和等于目标值 target = 4 时，2 是子数组的最短长度。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [2,4,6,8], target = 3
 * 输出：-1
 * 解释：在这个例子中 infinite_nums = [2,4,6,8,2,4,6,8,...] 。
 * 可以证明，不存在元素和等于目标值 target = 3 的子数组。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^5
 * 1 <= target <= 10^9
 * 
 * 
 */
public class Solution
{
    public int MinSizeSubarray(int[] nums, int target)
    {
        var n = nums.Length;
        var total = nums.Select(n => (long)n).Sum();
        var d = new Dictionary<long, int>();
        var S = 0l;
        d[S] = 0;
        for (var i = 0; i < n; i++)
        {
            S = (S + (long)nums[i]) % total;
            d[S] = i + 1;
        }
        S = 0;
        var ans = n + 1;
        for (var i = 0; i < n; i++)
        {
            var key = (S - (long)target % total + total) % total;
            if (d.ContainsKey(key))
            {
                ans = Math.Min(ans, (i - d[key] + n) % n);
            }
            S = (S + (long)nums[i]) % total;
        }
        return ans > n ? -1 : (ans + (int)((long)target / total * (long)n));
    }
}
