/*
 * @lc app=leetcode.cn id=ways-to-make-a-fair-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1664] 生成平衡数组的方案数
 *
 * https://leetcode.cn/problems/ways-to-make-a-fair-array/description/
 *
 * algorithms
 * Medium (55.87%)
 * Total Accepted:    10.6K
 * Total Submissions: 17.3K
 * Testcase Example:  '[2,1,6,4]'
 *
 * 给你一个整数数组 nums 。你需要选择 恰好 一个下标（下标从 0 开始）并删除对应的元素。请注意剩下元素的下标可能会因为删除操作而发生改变。
 * 
 * 比方说，如果 nums = [6,1,7,4,1] ，那么：
 * 
 * 
 * 选择删除下标 1 ，剩下的数组为 nums = [6,7,4,1] 。
 * 选择删除下标 2 ，剩下的数组为 nums = [6,1,4,1] 。
 * 选择删除下标 4 ，剩下的数组为 nums = [6,1,7,4] 。
 * 
 * 
 * 如果一个数组满足奇数下标元素的和与偶数下标元素的和相等，该数组就是一个 平衡数组 。
 * 
 * 请你返回删除操作后，剩下的数组 nums 是 平衡数组 的 方案数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,1,6,4]
 * 输出：1
 * 解释：
 * 删除下标 0 ：[1,6,4] -> 偶数元素下标为：1 + 4 = 5 。奇数元素下标为：6 。不平衡。
 * 删除下标 1 ：[2,6,4] -> 偶数元素下标为：2 + 4 = 6 。奇数元素下标为：6 。平衡。
 * 删除下标 2 ：[2,1,4] -> 偶数元素下标为：2 + 4 = 6 。奇数元素下标为：1 。不平衡。
 * 删除下标 3 ：[2,1,6] -> 偶数元素下标为：2 + 6 = 8 。奇数元素下标为：1 。不平衡。
 * 只有一种让剩余数组成为平衡数组的方案。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,1,1]
 * 输出：3
 * 解释：你可以删除任意元素，剩余数组都是平衡数组。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：0
 * 解释：不管删除哪个元素，剩下数组都不是平衡数组。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int WaysToMakeFair(int[] nums)
 *     {
 *         var n = nums.Length;
 *         var oddLeft = 0;
 *         var evenLeft = 0;
 *         var oddRight = 0;
 *         var evenRight = 0;
 *         for (var i = 0; i < n; i++)
 *         {
 *             oddLeft += (i & 1) * nums[i];
 *             evenLeft += ((i + 1) & 1) * nums[i];
 *         }
 *         var ans = 0;
 *         for (var i = 0; i < n; i++)
 *         {
 *             oddLeft -= (i & 1) * nums[i];
 *             evenLeft -= ((i + 1) & 1) * nums[i];
 *             if (oddLeft + evenRight == oddRight + evenLeft)
 *             {
 *                 ans++;
 *             }
 *             oddRight += (i & 1) * nums[i];
 *             evenRight += ((i + 1) & 1) * nums[i];
 *         }
 *         return ans;
 *     }
 * }
 */

public class Solution
{
    public int WaysToMakeFair(int[] nums)
    {
        var n = nums.Length;
        var S = new List<int>[2];
        for (var i = 0; i < 2; i++)
        {
            S[i] = new();
            S[i].Add(0);
        }
        for (var i = 0; i < n; i++)
        {
            S[i & 1].Add(S[i & 1][^1] + nums[i]);
        }
        var ans = 0;
        for (var i = 0; i < n; i++)
        {
            var even = S[0][(i + 1) >> 1] + S[1][^1] - S[1][(i + 1) >> 1];
            var odd = S[1][i >> 1] + S[0][^1] - S[0][(i >> 1) + 1];
            if (even == odd)
            {
                ans++;
            }
        }
        return ans;
    }
}
