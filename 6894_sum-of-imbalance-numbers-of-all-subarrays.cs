/*
 * @lc app=leetcode.cn id=sum-of-imbalance-numbers-of-all-subarrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6894] 所有子数组中不平衡数字之和
 *
 * https://leetcode.cn/problems/sum-of-imbalance-numbers-of-all-subarrays/description/
 *
 * algorithms
 * Hard (49.42%)
 * Total Accepted:    774
 * Total Submissions: 1.5K
 * Testcase Example:  '[2,3,1,4]'
 *
 * 一个长度为 n 下标从 0 开始的整数数组 arr 的 不平衡数字 定义为，在 sarr = sorted(arr)
 * 数组中，满足以下条件的下标数目：
 * 
 * 
 * 0 <= i < n - 1 ，和
 * sarr[i+1] - sarr[i] > 1
 * 
 * 
 * 这里，sorted(arr) 表示将数组 arr 排序后得到的数组。
 * 
 * 给你一个下标从 0 开始的整数数组 nums ，请你返回它所有 子数组 的 不平衡数字 之和。
 * 
 * 子数组指的是一个数组中连续一段 非空 的元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,3,1,4]
 * 输出：3
 * 解释：总共有 3 个子数组有非 0 不平衡数字：
 * - 子数组 [3, 1] ，不平衡数字为 1 。
 * - 子数组 [3, 1, 4] ，不平衡数字为 1 。
 * - 子数组 [1, 4] ，不平衡数字为 1 。
 * 其他所有子数组的不平衡数字都是 0 ，所以所有子数组的不平衡数字之和为 3 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [1,3,3,3,5]
 * 输出：8
 * 解释：总共有 7 个子数组有非 0 不平衡数字：
 * - 子数组 [1, 3] ，不平衡数字为 1 。
 * - 子数组 [1, 3, 3] ，不平衡数字为 1 。
 * - 子数组 [1, 3, 3, 3] ，不平衡数字为 1 。
 * - 子数组 [1, 3, 3, 3, 5] ，不平衡数字为 2 。
 * - 子数组 [3, 3, 3, 5] ，不平衡数字为 1 。
 * - 子数组 [3, 3, 5] ，不平衡数字为 1 。
 * - 子数组 [3, 5] ，不平衡数字为 1 。
 * 其他所有子数组的不平衡数字都是 0 ，所以所有子数组的不平衡数字之和为 8 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= nums.length
 * 
 * 
 */
public class Solution
{
    /*
     * public int SumImbalanceNumbers(int[] nums)
     * {
     *     var ans = 0;
     *     for (var (i, n) = (0, nums.Length); i < n; i++)
     *     {
     *         var (l, r, lMax, rMax) = (i - 1, i + 1, i - 1, i + 1);
     *         for (; l >= 0 && nums[l] != nums[i] + 1; l--) { }
     *         for (; r < n && nums[r] != nums[i] && nums[r] != nums[i] + 1; r++) { }
     *         for (; lMax >= 0 && nums[lMax] <= nums[i]; lMax--) { }
     *         for (; rMax < n && nums[rMax] < nums[i]; rMax++) { }
     *         ans += (r - i) * (i - l) - (rMax - i) * (i - lMax);
     *     }
     *     return ans;
     * }
     */

    // O(n) 贡献法. 具体巧思见注释.
    public int SumImbalanceNumbers(int[] nums)
    {
        var n = nums.Length;
        var ans = 0;
        var indexes = new int[n + 2];
        Array.Fill(indexes, n);
        // `rights`: 当前位置右侧最近的 `nums[i]` or `nums[i] + 1`.
        var rights = new int[n];
        for (var i = n - 1; i >= 0; i--)
        {
            rights[i] = Math.Min(indexes[nums[i]], indexes[nums[i] + 1]);
            indexes[nums[i]] = i;
        }
        // `lefts`: 当前位置左侧最近的 `nums[i]`; 直接用 `indexes` 代替.
        Array.Fill(indexes, -1);
        for (var i = 0; i < n; i++)
        {
            ans += (rights[i] - i) * (i - indexes[nums[i] + 1]);
            indexes[nums[i]] = i;
        }
        // 上面没有统计当前位置为最大值的不平衡数字.
        // 但其实每个子数组都有一个最大值, 所以这部分贡献就是子数组数目.
        // 整体减去即可.
        ans -= n * (n + 1) / 2;
        return ans;
    }
}
