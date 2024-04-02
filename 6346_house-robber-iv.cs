/*
 * @lc app=leetcode.cn id=house-robber-iv lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6346] 打家劫舍 IV
 *
 * https://leetcode.cn/problems/house-robber-iv/description/
 *
 * algorithms
 * Medium (36.48%)
 * Total Accepted:    2.7K
 * Total Submissions: 6.1K
 * Testcase Example:  '[2,3,5,9]\n2'
 *
 * 沿街有一排连续的房屋。每间房屋内都藏有一定的现金。现在有一位小偷计划从这些房屋中窃取现金。
 * 
 * 由于相邻的房屋装有相互连通的防盗系统，所以小偷 不会窃取相邻的房屋 。
 * 
 * 小偷的 窃取能力 定义为他在窃取过程中能从单间房屋中窃取的 最大金额 。
 * 
 * 给你一个整数数组 nums 表示每间房屋存放的现金金额。形式上，从左起第 i 间房屋中放有 nums[i] 美元。
 * 
 * 另给你一个整数数组 k ，表示窃贼将会窃取的 最少 房屋数。小偷总能窃取至少 k 间房屋。
 * 
 * 返回小偷的 最小 窃取能力。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,3,5,9], k = 2
 * 输出：5
 * 解释：
 * 小偷窃取至少 2 间房屋，共有 3 种方式：
 * - 窃取下标 0 和 2 处的房屋，窃取能力为 max(nums[0], nums[2]) = 5 。
 * - 窃取下标 0 和 3 处的房屋，窃取能力为 max(nums[0], nums[3]) = 9 。
 * - 窃取下标 1 和 3 处的房屋，窃取能力为 max(nums[1], nums[3]) = 9 。
 * 因此，返回 min(5, 9, 9) = 5 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,7,9,3,1], k = 2
 * 输出：2
 * 解释：共有 7 种窃取方式。窃取能力最小的情况所对应的方式是窃取下标 0 和 4 处的房屋。返回 max(nums[0], nums[4]) = 2
 * 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 1 <= k <= (nums.length + 1)/2
 * 
 * 
 */

// 看到 "最小化最大" / "最大化最小" 就要想二分!
public class Solution
{
    private bool CanRobKHouses(int[] arr, int k, int n)
    {
        var dp = new int[3];
        foreach (var a in arr)
        {
            dp[2] = dp[1];
            if (a <= n)
            {
                dp[2] = Math.Max(dp[2], dp[0] + 1);
            }
            dp[0] = dp[1];
            dp[1] = dp[2];
        }
        return dp[2] >= k;
    }

    public int MinCapability(int[] nums, int k)
    {
        var p = 0;
        var q = nums.Max();
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (CanRobKHouses(nums, k, mid))
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        return p;
    }
}
