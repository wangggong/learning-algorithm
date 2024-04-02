/*
 * @lc app=leetcode.cn id=minimum-operations-to-form-subsequence-with-target-sum lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2835] 使子序列的和等于目标的最少操作次数
 *
 * https://leetcode.cn/problems/minimum-operations-to-form-subsequence-with-target-sum/description/
 *
 * algorithms
 * Hard (29.88%)
 * Total Accepted:    2.2K
 * Total Submissions: 7.3K
 * Testcase Example:  '[1,2,8]\n7'
 *
 * 给你一个下标从 0 开始的数组 nums ，它包含 非负 整数，且全部为 2 的幂，同时给你一个整数 target 。
 * 
 * 一次操作中，你必须对数组做以下修改：
 * 
 * 
 * 选择数组中一个元素 nums[i] ，满足 nums[i] > 1 。
 * 将 nums[i] 从数组中删除。
 * 在 nums 的 末尾 添加 两个 数，值都为 nums[i] / 2 。
 * 
 * 
 * 你的目标是让 nums 的一个 子序列 的元素和等于 target ，请你返回达成这一目标的 最少操作次数 。如果无法得到这样的子序列，请你返回 -1
 * 。
 * 
 * 数组中一个 子序列 是通过删除原数组中一些元素，并且不改变剩余元素顺序得到的剩余数组。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,8], target = 7
 * 输出：1
 * 解释：第一次操作中，我们选择元素 nums[2] 。数组变为 nums = [1,2,4,4] 。
 * 这时候，nums 包含子序列 [1,2,4] ，和为 7 。
 * 无法通过更少的操作得到和为 7 的子序列。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,32,1,2], target = 12
 * 输出：2
 * 解释：第一次操作中，我们选择元素 nums[1] 。数组变为 nums = [1,1,2,16,16] 。
 * 第二次操作中，我们选择元素 nums[3] 。数组变为 nums = [1,1,2,16,8,8] 。
 * 这时候，nums 包含子序列 [1,1,2,8] ，和为 12 。
 * 无法通过更少的操作得到和为 12 的子序列。
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,32,1], target = 35
 * 输出：-1
 * 解释：无法得到和为 35 的子序列。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 2^30
 * nums 只包含非负整数，且均为 2 的幂。
 * 1 <= target < 2^31
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-operations-to-form-subsequence-with-target-sum/solutions/2413344/tan-xin-by-endlesscheng-immn/
//
// 贪心. 从小到大贪, 反了就会漏 case.
public class Solution
{
    public int MinOperations(IList<int> nums, int target)
    {
        if (nums.Select(n => (long)n).Sum() < target) { return -1; }
        const int D = 31;
        var count = nums.GroupBy(x => x)
            .ToDictionary(g => g.Key, g => (long)g.Count());
        for (var d = 0; d <= D; d++)
        {
            count.TryGetValue(1 << d, out var c);
            count[1 << d] = c;
        }
        var ans = 0;
        for (var (d, total) = (0, 0L); d <= D; d++)
        {
            total += count[1 << d] << d;
            if (((target >> d) & 1) is 0) { continue; }
            if (total >= 1L << d)
            {
                total -= 1L << d;
                continue;
            }
            for (var next = d + 1; next <= D; next++)
            {
                if (count[1 << next] is 0) { continue; }
                count[1 << next]--;
                total += 1L << next;
                total -= 1L << d;
                ans += next - d;
                break;
            }
        }
        return ans;
    }
}
