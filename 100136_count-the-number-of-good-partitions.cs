/*
 * @lc app=leetcode.cn id=count-the-number-of-good-partitions lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100136] 统计好分割方案的数目
 *
 * https://leetcode.cn/problems/count-the-number-of-good-partitions/description/
 *
 * algorithms
 * Hard (47.27%)
 * Total Accepted:    2.1K
 * Total Submissions: 4.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个下标从 0 开始、由 正整数 组成的数组 nums。
 * 
 * 将数组分割成一个或多个 连续 子数组，如果不存在包含了相同数字的两个子数组，则认为是一种 好分割方案 。
 * 
 * 返回 nums 的 好分割方案 的 数目。
 * 
 * 由于答案可能很大，请返回答案对 10^9 + 7 取余 的结果。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,4]
 * 输出：8
 * 解释：有 8 种 好分割方案 ：([1], [2], [3], [4]), ([1], [2], [3,4]), ([1], [2,3], [4]),
 * ([1], [2,3,4]), ([1,2], [3], [4]), ([1,2], [3,4]), ([1,2,3], [4]) 和
 * ([1,2,3,4]) 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,1,1,1]
 * 输出：1
 * 解释：唯一的 好分割方案 是：([1,1,1,1]) 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,2,1,3]
 * 输出：2
 * 解释：有 2 种 好分割方案 ：([1,2,1], [3]) 和 ([1,2,1,3]) 。
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
    private const long Mod = (long)1e9 + 7;

    private long Pow(long n, long k)
    {
        var ans = 1l;
        for (; k > 0; k >>= 1)
        {
            if ((k & 1) is not 0) { ans = ans * n % Mod; }
            n = n * n % Mod;
        }
        return ans;
    }

    public int NumberOfGoodPartitions(int[] nums)
    {
        var n = nums.Length;
        var d = new Dictionary<int, int>();
        for (var i = 0; i < n; i++)
        {
            d[nums[i]] = i;
        }
        var c = 0;
        for (var (i, max) = (0, 0); i < n; i++)
        {
            max = Math.Max(max, d[nums[i]]);
            if (i == max) { c++; }
        }
        return (int)Pow(2, c - 1);
    }
}
