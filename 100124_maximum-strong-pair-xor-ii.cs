/*
 * @lc app=leetcode.cn id=maximum-strong-pair-xor-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100124] 找出强数对的最大异或值 II
 *
 * https://leetcode.cn/problems/maximum-strong-pair-xor-ii/description/
 *
 * algorithms
 * Hard (27.67%)
 * Total Accepted:    1.6K
 * Total Submissions: 5.8K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。如果一对整数 x 和 y 满足以下条件，则称其为 强数对 ：
 * 
 * 
 * |x - y| <= min(x, y)
 * 
 * 
 * 你需要从 nums 中选出两个整数，且满足：这两个整数可以形成一个强数对，并且它们的按位异或（XOR）值是在该数组所有强数对中的 最大值 。
 * 
 * 返回数组 nums 所有可能的强数对中的 最大 异或值。
 * 
 * 注意，你可以选择同一个整数两次来形成一个强数对。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,4,5]
 * 输出：7
 * 解释：数组 nums 中有 11 个强数对：(1, 1), (1, 2), (2, 2), (2, 3), (2, 4), (3, 3), (3,
 * 4), (3, 5), (4, 4), (4, 5) 和 (5, 5) 。
 * 这些强数对中的最大异或值是 3 XOR 4 = 7 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [10,100]
 * 输出：0
 * 解释：数组 nums 中有 2 个强数对：(10, 10) 和 (100, 100) 。
 * 这些强数对中的最大异或值是 10 XOR 10 = 0 ，数对 (100, 100) 的异或值也是 100 XOR 100 = 0 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [500,520,2500,3000]
 * 输出：1020
 * 解释：数组 nums 中有 6 个强数对：(500, 500), (500, 520), (520, 520), (2500, 2500),
 * (2500, 3000) 和 (3000, 3000) 。
 * 这些强数对中的最大异或值是 500 XOR 520 = 1020 ；另一个异或值非零的数对是 (5, 6) ，其异或值是 2500 XOR 3000 =
 * 636 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 5 * 10^4
 * 1 <= nums[i] <= 2^20 - 1
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/maximum-strong-pair-xor-ii/solutions/2523213/0-1-trie-hua-dong-chuang-kou-pythonjavac-gvv2/
//
// 哈希表做法. 参考了经典的 2-sum.
public class Solution
{
    public int MaximumStrongPairXor(int[] nums)
    {
        const int D = 20;
        Array.Sort(nums);
        var ans = 0;
        var mask = 0;
        for (var d = D; d >= 0; d--)
        {
            mask |= 1 << d;
            var next = ans | (1 << d);
            var index = new Dictionary<int, int>();
            foreach (var n in nums)
            {
                var m = n & mask;
                if (index.ContainsKey(next ^ m) && index[next ^ m] * 2 >= n)
                {
                    ans = next;
                    break;
                }
                index[m] = n;
            }
        }
        return ans;
    }
}
