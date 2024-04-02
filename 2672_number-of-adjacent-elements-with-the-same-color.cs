/*
 * @lc app=leetcode.cn id=number-of-adjacent-elements-with-the-same-color lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2672] 有相同颜色的相邻元素数目
 *
 * https://leetcode.cn/problems/number-of-adjacent-elements-with-the-same-color/description/
 *
 * algorithms
 * Medium (57.89%)
 * Total Accepted:    3.6K
 * Total Submissions: 6.2K
 * Testcase Example:  '4\n[[0,2],[1,2],[3,1],[1,1],[2,1]]'
 *
 * 给你一个下标从 0 开始、长度为 n 的数组 nums 。一开始，所有元素都是 未染色 （值为 0 ）的。
 * 
 * 给你一个二维整数数组 queries ，其中 queries[i] = [indexi, colori] 。
 * 
 * 对于每个操作，你需要将数组 nums 中下标为 indexi 的格子染色为 colori 。
 * 
 * 请你返回一个长度与 queries 相等的数组 answer ，其中 answer[i]是前 i 个操作 之后 ，相邻元素颜色相同的数目。
 * 
 * 更正式的，answer[i] 是执行完前 i 个操作后，0 <= j < n - 1 的下标 j 中，满足 nums[j] == nums[j + 1]
 * 且 nums[j] != 0 的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 4, queries = [[0,2],[1,2],[3,1],[1,1],[2,1]]
 * 输出：[0,1,1,0,2]
 * 解释：一开始数组 nums = [0,0,0,0] ，0 表示数组中还没染色的元素。
 * - 第 1 个操作后，nums = [2,0,0,0] 。相邻元素颜色相同的数目为 0 。
 * - 第 2 个操作后，nums = [2,2,0,0] 。相邻元素颜色相同的数目为 1 。
 * - 第 3 个操作后，nums = [2,2,0,1] 。相邻元素颜色相同的数目为 1 。
 * - 第 4 个操作后，nums = [2,1,0,1] 。相邻元素颜色相同的数目为 0 。
 * - 第 5 个操作后，nums = [2,1,1,1] 。相邻元素颜色相同的数目为 2 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1, queries = [[0,100000]]
 * 输出：[0]
 * 解释：一开始数组 nums = [0] ，0 表示数组中还没染色的元素。
 * - 第 1 个操作后，nums = [100000] 。相邻元素颜色相同的数目为 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * 1 <= queries.length <= 10^5
 * queries[i].length == 2
 * 0 <= indexi <= n - 1
 * 1 <=  colori <= 10^5
 * 
 * 
 */
public class Solution
{
    public int[] ColorTheArray(int n, int[][] queries)
    {
        var colors = new int[n];
        int get(int k, int c) => c == 0
            ? 0
            : ((k > 0 && c == colors[k - 1] ? 1 : 0)
                + (k + 1 < n && c == colors[k + 1] ? 1 : 0));
        var q = queries.Length;
        var ans = new int[q];
        colors[queries[0][0]] = queries[0][1];
        for (var i = 1; i < q; i++)
        {
            var (index, color) = (queries[i][0], queries[i][1]);
            ans[i] = ans[i - 1] - get(index, colors[index]) + get(index, color);
            colors[index] = color;
        }
        return ans;
    }
}
