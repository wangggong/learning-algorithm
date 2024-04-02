/*
 * Hard
 * 
 * 给你一个下标从 0 开始的非负整数数组 nums 和两个整数 l 和 r 。
 * 
 * 请你返回 nums 中子多重集合的和在闭区间 [l, r] 之间的 子多重集合的数目 。
 * 
 * 由于答案可能很大，请你将答案对 109 + 7 取余后返回。
 * 
 * 子多重集合 指的是从数组中选出一些元素构成的 无序 集合，每个元素 x 出现的次数可以是 0, 1, ..., occ[x] 次，其中 occ[x] 是元素 x 在数组中的出现次数。
 * 
 * 注意：
 * 
 * 如果两个子多重集合中的元素排序后一模一样，那么它们两个是相同的 子多重集合 。
 * 空 集合的和是 0 。
 *  
 * 
 * 示例 1：
 * 
 * 输入：nums = [1,2,2,3], l = 6, r = 6
 * 输出：1
 * 解释：唯一和为 6 的子集合是 {1, 2, 3} 。
 * 示例 2：
 * 
 * 输入：nums = [2,1,4,2,7], l = 1, r = 5
 * 输出：7
 * 解释：和在闭区间 [1, 5] 之间的子多重集合为 {1} ，{2} ，{4} ，{2, 2} ，{1, 2} ，{1, 4} 和 {1, 2, 2} 。
 * 示例 3：
 * 
 * 输入：nums = [1,2,1,3,5,2], l = 3, r = 5
 * 输出：9
 * 解释：和在闭区间 [3, 5] 之间的子多重集合为 {3} ，{5} ，{1, 2} ，{1, 3} ，{2, 2} ，{2, 3} ，{1, 1, 2} ，{1, 1, 3} 和 {1, 2, 2} 。
 *  
 * 
 * 提示：
 * 
 * 1 <= nums.length <= 2 * 104
 * 0 <= nums[i] <= 2 * 104
 * nums 的和不超过 2 * 104 。
 * 0 <= l <= r <= 2 * 104
 * 通过次数 1.2K
 * 提交次数 4.5K
 * 通过率 27.8%
 * 
 */

// 多重背包.
//
// 参考: https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/solutions/2482876/duo-zhong-bei-bao-fang-an-shu-cong-po-su-f5ay/

/*
 * // TLE
 * public class Solution
 * {
 *     public int CountSubMultisets(IList<int> nums, int l, int r)
 *     {
 *         const long Mod = (long)1e9 + 7;
 *         var dp = new long[r + 1];
 *         dp[0] = 1;
 *         var c0 = 1;
 *         foreach (var (i, c) in nums
 *             .GroupBy(x => x)
 *             .Select(g => (g.Key, g.Count())))
 *         {
 *             if (i is 0)
 *             {
 *                 c0 += c;
 *                 continue;
 *             }
 *             for (var j = r; j >= 0; j--)
 *             {
 *                 for (var k = 1; k <= c && j - i * k >= 0; k++)
 *                 {
 *                     dp[j] = (dp[j] + dp[j - i * k]) % Mod;
 *                 }
 *             }
 *         }
 *         return (int)((dp[l ..].Aggregate((x, y) => (x + y) % Mod) * c0) % Mod);
 *     }
 * }
 */

public class Solution
{
    public int CountSubMultisets(IList<int> nums, int l, int r)
    {
        const long Mod = (long)1e9 + 7;
        var dp = new long[r + 1];
        dp[0] = 1;
        var c0 = 1;
        foreach (var (i, c) in nums
            .GroupBy(x => x)
            .Select(g => (g.Key, g.Count())))
        {
            if (i is 0)
            {
                c0 += c;
                continue;
            }
            for (var j = 0; j <= r; j++)
            {
                dp[j] = (dp[j] + (j >= i ? dp[j - i] : 0)) % Mod;
            }
            for (var j = r; j >= 0; j--)
            {
                dp[j] = (dp[j] - (j >= (c + 1) * i ? dp[j - (c + 1) * i] : 0)) % Mod;
            }
        }
        return (int)((dp[l ..].Aggregate((x, y) => (x + y) % Mod) * c0 % Mod + Mod) % Mod);
    }
}
