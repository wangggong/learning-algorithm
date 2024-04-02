/*
 * @lc app=leetcode.cn id=minimum-adjacent-swaps-for-k-consecutive-ones lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1703] 得到连续 K 个 1 的最少相邻交换次数
 *
 * https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/description/
 *
 * algorithms
 * Hard (37.99%)
 * Total Accepted:    4.5K
 * Total Submissions: 9.5K
 * Testcase Example:  '[1,0,0,1,0,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k 。 nums 仅包含 0 和 1 。每一次移动，你可以选择 相邻 两个数字并将它们交换。
 * 
 * 请你返回使 nums 中包含 k 个 连续 1 的 最少 交换次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [1,0,0,1,0,1], k = 2
 * 输出：1
 * 解释：在第一次操作时，nums 可以变成 [1,0,0,0,1,1] 得到连续两个 1 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [1,0,0,0,0,0,1,1], k = 3
 * 输出：5
 * 解释：通过 5 次操作，最左边的 1 可以移到右边直到 nums 变为 [0,0,0,0,0,1,1,1] 。
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums = [1,1,0,1], k = 2
 * 输出：0
 * 解释：nums 已经有连续 2 个 1 了。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * nums[i] 要么是 0 ，要么是 1 。
 * 1 <= k <= sum(nums)
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/solution/tu-jie-zhuan-huan-cheng-zhong-wei-shu-ta-iz4v/
public class Solution
{
    public int MinMoves(int[] nums, int k)
    {
        var n = nums.Length;
        var ones = new List<int>();
        for (var i = 0; i < n; i++)
        {
            if (nums[i] != 0)
            {
                ones.Add(i - ones.Count);
            }
        }
        var m = ones.Count;
        var prefixSum = new int[m + 1];
        for (var i = 0; i < m; i++)
        {
            prefixSum[i + 1] = prefixSum[i] + ones[i];
        }
        var ans = Int32.MaxValue;
        for (var i = 0; i + k - 1 < m; i++)
        {
            var cur = prefixSum[i] + prefixSum[i + k] - 2 * prefixSum[i + k / 2] - (k % 2) * ones[i + k / 2];
            ans = Math.Min(ans, cur);
        }
        return ans;
    }
}
