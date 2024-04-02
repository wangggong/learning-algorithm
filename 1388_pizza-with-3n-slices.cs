/*
 * @lc app=leetcode.cn id=pizza-with-3n-slices lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1388] 3n 块披萨
 *
 * https://leetcode.cn/problems/pizza-with-3n-slices/description/
 *
 * algorithms
 * Hard (56.56%)
 * Total Accepted:    8K
 * Total Submissions: 13.8K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给你一个披萨，它由 3n 块不同大小的部分组成，现在你和你的朋友们需要按照如下规则来分披萨：
 * 
 * 
 * 你挑选 任意 一块披萨。
 * Alice 将会挑选你所选择的披萨逆时针方向的下一块披萨。
 * Bob 将会挑选你所选择的披萨顺时针方向的下一块披萨。
 * 重复上述过程直到没有披萨剩下。
 * 
 * 
 * 每一块披萨的大小按顺时针方向由循环数组 slices 表示。
 * 
 * 请你返回你可以获得的披萨大小总和的最大值。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：slices = [1,2,3,4,5,6]
 * 输出：10
 * 解释：选择大小为 4 的披萨，Alice 和 Bob 分别挑选大小为 3 和 5 的披萨。然后你选择大小为 6 的披萨，Alice 和 Bob
 * 分别挑选大小为 2 和 1 的披萨。你获得的披萨总大小为 4 + 6 = 10 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：slices = [8,9,8,6,1,1]
 * 输出：16
 * 解释：两轮都选大小为 8 的披萨。如果你选择大小为 9 的披萨，你的朋友们就会选择大小为 8 的披萨，这种情况下你的总和不是最大的。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= slices.length <= 500
 * slices.length % 3 == 0
 * 1 <= slices[i] <= 1000
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/pizza-with-3n-slices/solutions/163091/hua-wei-bu-xiang-lin-de-zi-shu-lie-de-zui-da-he-we/
public class Solution
{
    public int MaxSizeSlices(int[] slices) => Math.Max(
        Dp(slices[1 ..].ToArray()), Dp(slices[.. ^1].ToArray()));

    private int Dp(int[] arr)
    {
        const int K = 3;
        var N = arr.Length;
        var n = (N / K) + 1;
        var dp = Enumerable.Range(0, 2).Select(_ => new int[N]).ToArray();
        for (var i = 0; i < n; i++)
        {
            for (var (j, max) = (0, 0); j < N; j++)
            {
                if (j > 1) { max = Math.Max(max, dp[(i + 1) % 2][j - 2]); }
                dp[i % 2][j] = arr[j] + max;
            }
        }
        return dp[(n - 1) % 2].Max();
    }
}
