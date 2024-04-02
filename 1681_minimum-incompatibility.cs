/*
 * @lc app=leetcode.cn id=minimum-incompatibility lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1681] 最小不兼容性
 *
 * https://leetcode.cn/problems/minimum-incompatibility/description/
 *
 * algorithms
 * Hard (43.68%)
 * Total Accepted:    7K
 * Total Submissions: 13.3K
 * Testcase Example:  '[1,2,1,4]\n2'
 *
 * 给你一个整数数组 nums​​​ 和一个整数 k 。你需要将这个数组划分到 k 个相同大小的子集中，使得同一个子集里面没有两个相同的元素。
 * 
 * 一个子集的 不兼容性 是该子集里面最大值和最小值的差。
 * 
 * 请你返回将数组分成 k 个子集后，各子集 不兼容性 的 和 的 最小值 ，如果无法分成分成 k 个子集，返回 -1 。
 * 
 * 子集的定义是数组中一些数字的集合，对数字顺序没有要求。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,1,4], k = 2
 * 输出：4
 * 解释：最优的分配是 [1,2] 和 [1,4] 。
 * 不兼容性和为 (2-1) + (4-1) = 4 。
 * 注意到 [1,1] 和 [2,4] 可以得到更小的和，但是第一个集合有 2 个相同的元素，所以不可行。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [6,3,8,1,3,1,2,2], k = 4
 * 输出：6
 * 解释：最优的子集分配为 [1,2]，[2,3]，[6,8] 和 [1,3] 。
 * 不兼容性和为 (2-1) + (3-2) + (8-6) + (3-1) = 6 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [5,3,3,6,3,3], k = 3
 * 输出：-1
 * 解释：没办法将这些数字分配到 3 个子集且满足每个子集里没有相同数字。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * nums.length 能被 k 整除。
 * 1 
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-incompatibility/solution/zui-xiao-bu-jian-rong-xing-by-leetcode-s-iiac/
public class Solution
{
    private const int D = 32;
    private const int Maxn = 0x3f3f3f3f;

    public int MinimumIncompatibility(int[] nums, int k)
    {
        int bitCount(int n) => Enumerable.Range(0, D)
            .Count(i => (n & (1 << i)) is not 0);
        int bitLength(int n) => Enumerable.Range(0, D)
            .First(i => (n & (1 << i)) is not 0);
        Array.Sort(nums);
        var n = nums.Length;
        if (n % k is not 0
            || nums.GroupBy(x => x).Select(g => g.Count()).Max() > k)
        {
            return -1;
        }
        var m = n / k;
        var memo = new Dictionary<(int, int), int>();
        int dfs(int left, int start)
        {
            var key = (left, start);
            if (memo.ContainsKey(key))
            {
                return memo[key];
            }
            if (left is 0)
            {
                return memo[key] = 0;
            }
            if (bitCount(left) % m is 0)
            {
                var lowBit = left & -left;
                return memo[key] = dfs(left ^ lowBit, bitLength(lowBit));
            }
            memo[key] = Maxn;
            var last = nums[start];
            for (var i = start + 1; i < n; i++)
            {
                if ((left & (1 << i)) is not 0 && nums[i] != last)
                {
                    last = nums[i];
                    memo[key] = Math.Min(memo[key],
                        (last - nums[start]) + dfs(left ^ (1 << i), i));
                }
            }
            return memo[key];
        }
        return dfs((1 << n) - 2, 0);
    }
}

/*
 * // 自己写的版本, 超时了.
 * public class Solution
 * {
 *     private const int D = 16;
 *     private const int Maxn = 0x3f3f3f3f;
 * 
 *     public int MinimumIncompatibility(int[] nums, int k)
 *     {
 *         var n = nums.Length;
 *         if (n % k != 0 || nums.GroupBy(x => x).Select(g => g.Count()).Max() > k)
 *         {
 *             return -1;
 *         }
 *         var m = n / k;
 *         var memo = new Dictionary<(int, int, int, int, int), int>();
 *         int dfs(int start, int k, int valueVisit, int indexVisit, int left)
 *         {
 *             bool hasVisit(int k, int d) => (k & (1 << d)) is not 0;
 *             int visit(int k, int d) => k | (1 << d);
 *             var key = (start, k, valueVisit, indexVisit, left);
 *             if (memo.ContainsKey(key))
 *             {
 *                 return memo[key];
 *             }
 *             if (k is 0)
 *             {
 *                 return memo[key] = 0;
 *             }
 *             if (left is 0)
 *             {
 *                 var (min, max) = (0, 0);
 *                 for (var d = 1; d <= D; d++)
 *                 {
 *                     if (hasVisit(valueVisit, d))
 *                     {
 *                         min = (min is 0) ? d : min;
 *                         max = d;
 *                     }
 *                 }
 *                 return memo[key] = dfs(0, k - 1, 0, indexVisit, m) + (max - min);
 *             }
 *             if (start >= n)
 *             {
 *                 return memo[key] = Maxn;
 *             }
 *             memo[key] = Maxn;
 *             for (var i = start; i < n; i++)
 *             {
 *                 if (!hasVisit(indexVisit, i) && !hasVisit(valueVisit, nums[i]))
 *                 {
 *                     memo[key] = Math.Min(memo[key], dfs(i + 1, k, visit(valueVisit, nums[i]), visit(indexVisit, i), left - 1));
 *                 }
 *             }
 *             return memo[key];
 *         }
 *         return dfs(0, k, 0, 0, m);
 *     }
 * }
 */
