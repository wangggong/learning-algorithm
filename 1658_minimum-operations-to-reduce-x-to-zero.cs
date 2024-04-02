/*
 * @lc app=leetcode.cn id=minimum-operations-to-reduce-x-to-zero lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1658] 将 x 减到 0 的最小操作数
 *
 * https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/
 *
 * algorithms
 * Medium (32.84%)
 * Total Accepted:    17.4K
 * Total Submissions: 51.7K
 * Testcase Example:  '[1,1,4,2,3]\n5'
 *
 * 给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要
 * 修改 数组以供接下来的操作使用。
 * 
 * 如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,1,4,2,3], x = 5
 * 输出：2
 * 解释：最佳解决方案是移除后两个元素，将 x 减到 0 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,6,7,8,9], x = 4
 * 输出：-1
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [3,2,20,1,1,3], x = 10
 * 输出：5
 * 解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1e5
 * 1 <= nums[i] <= 1e4
 * 1 <= x <= 1e9
 * 
 * 
 */
public class Solution
{
    public int MinOperations(int[] nums, int x)
    {
        var n = nums.Length;
        var prefixSum = new int[n + 1];
        for (var i = 0; i < n; i++)
        {
            prefixSum[i + 1] = prefixSum[i] + nums[i];
        }
        var ans = n + 1;
        for (var i = 0; i <= n; i++)
        {
            var target = x - prefixSum[i];
            if (target < 0)
            {
                break;
            }
            var p = i;
            var q = n;
            while (p < q)
            {
                var mid = (p + q) >> 1;
                if (prefixSum[n] - prefixSum[mid] <= target)
                {
                    q = mid;
                }
                else
                {
                    p = mid + 1;
                }
            }
            if (prefixSum[n] - prefixSum[p] == target)
            {
                ans = Math.Min(ans, i + n - p);
            }
        }
        return ans > n ? -1 : ans;
    }
}
