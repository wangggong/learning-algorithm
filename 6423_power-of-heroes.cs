/*
 * @lc app=leetcode.cn id=power-of-heroes lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6423] 英雄的力量
 *
 * https://leetcode.cn/problems/power-of-heroes/description/
 *
 * algorithms
 * Hard (30.03%)
 * Total Accepted:    1.2K
 * Total Submissions: 3.8K
 * Testcase Example:  '[2,1,4]'
 *
 * 给你一个下标从 0 开始的整数数组 nums ，它表示英雄的能力值。如果我们选出一部分英雄，这组英雄的 力量 定义为：
 * 
 * 
 * i0 ，i1 ，... ik 表示这组英雄在数组中的下标。那么这组英雄的力量为 max(nums[i0],nums[i1] ...
 * nums[ik])^2 * min(nums[i0],nums[i1] ... nums[ik]) 。
 * 
 * 
 * 请你返回所有可能的 非空 英雄组的 力量 之和。由于答案可能非常大，请你将结果对 10^9 + 7 取余。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,1,4]
 * 输出：141
 * 解释：
 * 第 1 组：[2] 的力量为 2^2 * 2 = 8 。
 * 第 2 组：[1] 的力量为 1^2 * 1 = 1 。
 * 第 3 组：[4] 的力量为 4^2 * 4 = 64 。
 * 第 4 组：[2,1] 的力量为 2^2 * 1 = 4 。
 * 第 5 组：[2,4] 的力量为 4^2 * 2 = 32 。
 * 第 6 组：[1,4] 的力量为 4^2 * 1 = 16 。
 * 第​ ​​​​​​7 组：[2,1,4] 的力量为 4^2​​​​​​​ * 1 = 16 。
 * 所有英雄组的力量之和为 8 + 1 + 64 + 4 + 32 + 16 + 16 = 141 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,1,1]
 * 输出：7
 * 解释：总共有 7 个英雄组，每一组的力量都是 1 。所以所有英雄组的力量之和为 7 。
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

    public int SumOfPower(int[] nums)
    {
        var n = nums.Length;
        var arr = nums.Select(n => (long)n).OrderBy(x => x % Mod).ToArray();
        long ans = arr[n - 1] * arr[n - 1] % Mod * arr[n - 1] % Mod;
        long total = 0;
        for (var i = n - 2; i >= 0; i--)
        {
            total = (total * 2 % Mod + arr[i + 1] * arr[i + 1] % Mod) % Mod;
            ans = (ans + (total + arr[i] * arr[i]) % Mod * arr[i] % Mod) % Mod;
        }
        return (int)ans;
    }
}
