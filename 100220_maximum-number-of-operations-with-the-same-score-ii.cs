/*
 * @lc app=leetcode.cn id=maximum-number-of-operations-with-the-same-score-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100220] 相同分数的最大操作数目 II
 *
 * https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii/description/
 *
 * algorithms
 * Medium (27.31%)
 * Total Accepted:    1.2K
 * Total Submissions: 4.4K
 * Testcase Example:  '[3,2,1,2,3,4]'
 *
 * 给你一个整数数组 nums ，如果 nums 至少 包含 2 个元素，你可以执行以下操作中的 任意 一个：
 * 
 * 
 * 选择 nums 中最前面两个元素并且删除它们。
 * 选择 nums 中最后两个元素并且删除它们。
 * 选择 nums 中第一个和最后一个元素并且删除它们。
 * 
 * 
 * 一次操作的 分数 是被删除元素的和。
 * 
 * 在确保 所有操作分数相同 的前提下，请你求出 最多 能进行多少次操作。
 * 
 * 请你返回按照上述要求 最多 可以进行的操作次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,2,1,2,3,4]
 * 输出：3
 * 解释：我们执行以下操作：
 * - 删除前两个元素，分数为 3 + 2 = 5 ，nums = [1,2,3,4] 。
 * - 删除第一个元素和最后一个元素，分数为 1 + 4 = 5 ，nums = [2,3] 。
 * - 删除第一个元素和最后一个元素，分数为 2 + 3 = 5 ，nums = [] 。
 * 由于 nums 为空，我们无法继续进行任何操作。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,2,6,1,4]
 * 输出：2
 * 解释：我们执行以下操作：
 * - 删除前两个元素，分数为 3 + 2 = 5 ，nums = [6,1,4] 。
 * - 删除最后两个元素，分数为 1 + 4 = 5 ，nums = [6] 。
 * 至多进行 2 次操作。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= nums.length <= 2000
 * 1 <= nums[i] <= 1000
 * 
 * 
 */
public class Solution
{
    public int MaxOperations(int[] nums)
    {
        var n = nums.Length;
        var memo = new Dictionary<(int, int, int), int>();
        int dfs(int l, int r, int target) => memo[(l, r, target)] =
            memo.ContainsKey((l, r, target)) ? memo[(l, r, target)] : (
                r - l < 2 ? 0 : new int[]
                {
                    nums[l] + nums[l + 1] != target ? 0 : dfs(l + 2, r, target) + 1,
                    nums[l] + nums[r - 1] != target ? 0 : dfs(l + 1, r - 1, target) + 1,
                    nums[r - 2] + nums[r - 1] != target ? 0 : dfs(l, r - 2, target) + 1,
                }.Max());
        return new int[]
        {
            dfs(0, n, nums[0] + nums[1]),
            dfs(0, n, nums[0] + nums[n - 1]),
            dfs(0, n, nums[n - 2] + nums[n - 1]),
        }.Max();
    }
}
