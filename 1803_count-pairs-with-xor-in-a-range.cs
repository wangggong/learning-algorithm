/*
 * @lc app=leetcode.cn id=count-pairs-with-xor-in-a-range lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1803] 统计异或值在范围内的数对有多少
 *
 * https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/description/
 *
 * algorithms
 * Hard (43.15%)
 * Total Accepted:    5.4K
 * Total Submissions: 11.1K
 * Testcase Example:  '[1,4,2,7]\n2\n6'
 *
 * 给你一个整数数组 nums （下标 从 0 开始 计数）以及两个整数：low 和 high ，请返回 漂亮数对 的数目。
 * 
 * 漂亮数对 是一个形如 (i, j) 的数对，其中 0 <= i < j < nums.length 且 low <= (nums[i] XOR
 * nums[j]) <= high 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [1,4,2,7], low = 2, high = 6
 * 输出：6
 * 解释：所有漂亮数对 (i, j) 列出如下：
 * ⁠   - (0, 1): nums[0] XOR nums[1] = 5 
 * ⁠   - (0, 2): nums[0] XOR nums[2] = 3
 * ⁠   - (0, 3): nums[0] XOR nums[3] = 6
 * ⁠   - (1, 2): nums[1] XOR nums[2] = 6
 * ⁠   - (1, 3): nums[1] XOR nums[3] = 3
 * ⁠   - (2, 3): nums[2] XOR nums[3] = 5
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [9,8,4,2,1], low = 5, high = 14
 * 输出：8
 * 解释：所有漂亮数对 (i, j) 列出如下：
 * ​​​​​    - (0, 2): nums[0] XOR nums[2] = 13
 * - (0, 3): nums[0] XOR nums[3] = 11
 * - (0, 4): nums[0] XOR nums[4] = 8
 * - (1, 2): nums[1] XOR nums[2] = 12
 * - (1, 3): nums[1] XOR nums[3] = 10
 * - (1, 4): nums[1] XOR nums[4] = 9
 * - (2, 3): nums[2] XOR nums[3] = 6
 * - (2, 4): nums[2] XOR nums[4] = 5
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 2 * 10^4
 * 1 <= nums[i] <= 2 * 10^4
 * 1 <= low <= high <= 2 * 10^4
 * 
 * 
 */

/*
 * // 参考: https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/solution/javabao-li-jiao-xing-tong-guo-1650-ms-by-1jzs/
 * // 花式暴力
 * public class Solution
 * {
 *     public int CountPairs(int[] nums, int low, int high)
 *     {
 *         var count = nums.Count();
 *         var min = nums.Min();
 *         var max = nums.Max();
 *         foreach (var n in nums)
 *         {
 *             count[n]++;
 *         }
 *         var ans = 0;
 *         for (var i = min; i <= max; i++)
 *         {
 *             if (count[i] != 0)
 *             {
 *                 for (var j = low; j <= high; j++)
 *                 {
 *                     var k = i ^ j;
 *                     if (min <= k && k <= max)
 *                     {
 *                         ans += count[i] * count[k];
 *                     }
 *                 }
 *             }
 *         }
 *         return ans >> 1;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/solution/bu-hui-zi-dian-shu-zhi-yong-ha-xi-biao-y-p2pu/
// 用哈希表模拟静态统计下的 0-1 字典树. 灵佬 yyds
public class Solution
{
    public int CountPairs(int[] nums, int low, int high)
    {
        var count = nums.GroupBy(n => n).ToDictionary(g => g.Key, g => g.Count());
        var ans = 0;
        for (high++; high > 0; low >>= 1, high >>= 1)
        {
            var nextCount = new Dictionary<int, int>();
            foreach ((var k, var v) in count)
            {
                count.TryGetValue((high - 1) ^ k, out var highCount);
                count.TryGetValue((low - 1) ^ k, out var lowCount);
                nextCount.TryGetValue(k >> 1, out var kCount);
                ans += ((high & 1) * highCount - (low & 1) * lowCount) * v;
                nextCount[k >> 1] = kCount + v;
            }
            count = nextCount;
        }
        return ans >> 1;
    }
}
