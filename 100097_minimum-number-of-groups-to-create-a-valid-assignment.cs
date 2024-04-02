/*
 * Medium
 * 
 * 给你一个长度为 n 下标从 0 开始的整数数组 nums 。
 * 
 * 我们想将下标进行分组，使得 [0, n - 1] 内所有下标 i 都 恰好 被分到其中一组。
 * 
 * 如果以下条件成立，我们说这个分组方案是合法的：
 * 
 * 对于每个组 g ，同一组内所有下标在 nums 中对应的数值都相等。
 * 对于任意两个组 g1 和 g2 ，两个组中 下标数量 的 差值不超过 1 。
 * 请你返回一个整数，表示得到一个合法分组方案的 最少 组数。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：nums = [3,2,3,2,3]
 * 输出：2
 * 解释：一个得到 2 个分组的方案如下，中括号内的数字都是下标：
 * 组 1 -> [0,2,4]
 * 组 2 -> [1,3]
 * 所有下标都只属于一个组。
 * 组 1 中，nums[0] == nums[2] == nums[4] ，所有下标对应的数值都相等。
 * 组 2 中，nums[1] == nums[3] ，所有下标对应的数值都相等。
 * 组 1 中下标数目为 3 ，组 2 中下标数目为 2 。
 * 两者之差不超过 1 。
 * 无法得到一个小于 2 组的答案，因为如果只有 1 组，组内所有下标对应的数值都要相等。
 * 所以答案为 2 。
 * 示例 2：
 * 
 * 输入：nums = [10,10,10,3,1,1]
 * 输出：4
 * 解释：一个得到 2 个分组的方案如下，中括号内的数字都是下标：
 * 组 1 -> [0]
 * 组 2 -> [1,2]
 * 组 3 -> [3]
 * 组 4 -> [4,5]
 * 分组方案满足题目要求的两个条件。
 * 无法得到一个小于 4 组的答案。
 * 所以答案为 4 。
 *  
 * 
 * 提示：
 * 
 * 1 <= nums.length <= 105
 * 1 <= nums[i] <= 109
 * 通过次数 2.3K
 * 提交次数 10.7K
 * 通过率 21.1%
 * 
 */


// 参考: https://leetcode.cn/problems/minimum-number-of-groups-to-create-a-valid-assignment/solutions/2493313/ben-ti-zui-jian-dan-xie-fa-pythonjavacgo-t174/
//
// 确实是最简单写法. 关键点是尽量分尽可能多的 `k + 1` 组, 且能不能分组的条件在于 `c / k >= c % k`.
public class Solution
{
    public int MinGroupsForValidAssignment(int[] nums)
    {
        var counts = nums.GroupBy(x => x)
            .Select(g => g.Count())
            .OrderBy(x => x)
            .ToList();
        (int, bool) tryCreateAssignment(int k)
        {
            var ans = 0;
            foreach (var c in counts)
            {
                if (c / k < c % k) { return (0, false); }
                ans += (c + k) / (k + 1);
            }
            return (ans, true);
        }
        for (var k = counts.First(); true; k--)
        {
            var (ans, valid) = tryCreateAssignment(k);
            if (valid) { return ans; }
        }
    }
}

/*
 * // 基本该想到的都想到了, 但最少组数没想到就 WA 了四次.
 * public class Solution
 * {
 *     public int MinGroupsForValidAssignment(int[] nums)
 *     {
 *         var counts = nums.GroupBy(n => n)
 *             .Select(g => g.Count())
 *             .OrderBy(x => x)
 *             .ToArray();
 *         for (var i = counts.First(); true; i--)
 *         {
 *             var valid = true;
 *             var ans = 0;
 *             foreach (var c in counts)
 *             {
 *                 valid = false;
 *                 var last = ans;
 *                 var k = c;
 *                 if (c % (i + 1) == 0 || c % (i + 1) == i)
 *                 {
 *                     k = Math.Min(k, (c + i) / (i + 1));
 *                     valid = true;
 *                 }
 *                 if (c / (i + 1) + c % (i + 1) >= i + 1)
 *                 {
 *                     k = Math.Min(k, (c + i) / (i + 1));
 *                     valid = true;
 *                 }
 *                 if (c / i >= c % i)
 *                 {
 *                     k = Math.Min(k, c / i);
 *                     valid = true;
 *                 }
 *                 if (!valid) { break; }
 *                 ans += k;
 *             }
 *             if (valid) { return ans; }
 *         }
 *     }
 * }
 */
