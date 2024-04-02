/*
 * @lc app=leetcode.cn id=minimize-or-of-remaining-elements-using-operations lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100179] 给定操作次数内使剩余元素的或值最小
 *
 * https://leetcode.cn/problems/minimize-or-of-remaining-elements-using-operations/description/
 *
 * algorithms
 * Hard (30.15%)
 * Total Accepted:    349
 * Total Submissions: 1.1K
 * Testcase Example:  '[3,5,3,2,7]\n2'
 *
 * 给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。
 * 
 * 一次操作中，你可以选择 nums 中满足 0 <= i < nums.length - 1 的一个下标 i ，并将 nums[i] 和 nums[i +
 * 1] 替换为数字 nums[i] & nums[i + 1] ，其中 & 表示按位 AND 操作。
 * 
 * 请你返回 至多 k 次操作以内，使 nums 中所有剩余元素按位 OR 结果的 最小值 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,5,3,2,7], k = 2
 * 输出：3
 * 解释：执行以下操作：
 * 1. 将 nums[0] 和 nums[1] 替换为 (nums[0] & nums[1]) ，得到 nums 为 [1,3,2,7] 。
 * 2. 将 nums[2] 和 nums[3] 替换为 (nums[2] & nums[3]) ，得到 nums 为 [1,3,2] 。
 * 最终数组的按位或值为 3 。
 * 3 是 k 次操作以内，可以得到的剩余元素的最小按位或值。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [7,3,15,14,2,8], k = 4
 * 输出：2
 * 解释：执行以下操作：
 * 1. 将 nums[0] 和 nums[1] 替换为 (nums[0] & nums[1]) ，得到 nums 为 [3,15,14,2,8] 。
 * 2. 将 nums[0] 和 nums[1] 替换为 (nums[0] & nums[1]) ，得到 nums 为 [3,14,2,8] 。
 * 3. 将 nums[0] 和 nums[1] 替换为 (nums[0] & nums[1]) ，得到 nums 为 [2,2,8] 。
 * 4. 将 nums[1] 和 nums[2] 替换为 (nums[1] & nums[2]) ，得到 nums 为 [2,0] 。
 * 最终数组的按位或值为 2 。
 * 2 是 k 次操作以内，可以得到的剩余元素的最小按位或值。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [10,7,10,3,9,14,9,4], k = 1
 * 输出：15
 * 解释：不执行任何操作，nums 的按位或值为 15 。
 * 15 是 k 次操作以内，可以得到的剩余元素的最小按位或值。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] < 2^30
 * 0 <= k < nums.length
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimize-or-of-remaining-elements-using-operations/solutions/2622658/shi-tian-fa-pythonjavacgo-by-endlesschen-ysom/
//
// - [X] 从高位到低位思考
// - [X] 转化为连续子数组的操作
// - [ ] 贪心地计算要清位需要的操作次数, 与 `k` 比较
// - [ ] 用 `mask` 将之前的计算结果 (更高位的置位逻辑) 与现有结果联系起来
// 结论: 不是字典树
public class Solution
{
    public int MinOrAfterOperations(int[] nums, int k)
    {
        const int B = 30;
        var (ans, mask) = (0, 0);
        for (var b = B; b >= 0; b--)
        {
            mask |= 1 << b;
            var (count, cur) = (0, -1);
            foreach (var n in nums)
            {
                cur &= n & mask;
                if (cur is 0) { cur = -1; }
                else { count++; }
            }
            if (count <= k) { continue; }
            mask ^= 1 << b;
            ans |= 1 << b;
        }
        return ans;
    }
}
