/*
 * @lc app=leetcode.cn id=number-of-ways-to-earn-points lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6310] 获得分数的方法数
 *
 * https://leetcode.cn/problems/number-of-ways-to-earn-points/description/
 *
 * algorithms
 * Hard (62.40%)
 * Total Accepted:    3K
 * Total Submissions: 4.7K
 * Testcase Example:  '6\n[[6,1],[3,2],[2,3]]'
 *
 * 考试中有 n 种类型的题目。给你一个整数 target 和一个下标从 0 开始的二维整数数组 types ，其中 types[i] = [counti,
 * marksi] 表示第 i 种类型的题目有 counti 道，每道题目对应 marksi 分。
 * 
 * 返回你在考试中恰好得到 target 分的方法数。由于答案可能很大，结果需要对 10^9 +7 取余。
 * 
 * 注意，同类型题目无法区分。
 * 
 * 
 * 比如说，如果有 3 道同类型题目，那么解答第 1 和第 2 道题目与解答第 1 和第 3 道题目或者第 2 和第 3 道题目是相同的。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：target = 6, types = [[6,1],[3,2],[2,3]]
 * 输出：7
 * 解释：要获得 6 分，你可以选择以下七种方法之一：
 * - 解决 6 道第 0 种类型的题目：1 + 1 + 1 + 1 + 1 + 1 = 6
 * - 解决 4 道第 0 种类型的题目和 1 道第 1 种类型的题目：1 + 1 + 1 + 1 + 2 = 6
 * - 解决 2 道第 0 种类型的题目和 2 道第 1 种类型的题目：1 + 1 + 2 + 2 = 6
 * - 解决 3 道第 0 种类型的题目和 1 道第 2 种类型的题目：1 + 1 + 1 + 3 = 6
 * - 解决 1 道第 0 种类型的题目、1 道第 1 种类型的题目和 1 道第 2 种类型的题目：1 + 2 + 3 = 6
 * - 解决 3 道第 1 种类型的题目：2 + 2 + 2 = 6
 * - 解决 2 道第 2 种类型的题目：3 + 3 = 6
 * 
 * 
 * 示例 2：
 * 
 * 输入：target = 5, types = [[50,1],[50,2],[50,5]]
 * 输出：4
 * 解释：要获得 5 分，你可以选择以下四种方法之一：
 * - 解决 5 道第 0 种类型的题目：1 + 1 + 1 + 1 + 1 = 5
 * - 解决 3 道第 0 种类型的题目和 1 道第 1 种类型的题目：1 + 1 + 1 + 2 = 5
 * - 解决 1 道第 0 种类型的题目和 2 道第 1 种类型的题目：1 + 2 + 2 = 5
 * - 解决 1 道第 2 种类型的题目：5
 * 
 * 
 * 示例 3：
 * 
 * 输入：target = 18, types = [[6,1],[3,2],[2,3]]
 * 输出：1
 * 解释：只有回答所有题目才能获得 18 分。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= target <= 1000
 * n == types.length
 * 1 <= n <= 50
 * types[i].length == 2
 * 1 <= counti, marksi <= 50
 * 
 * 
 */
public class Solution
{
    private const long Mod = (long)1e9 + 7;

    public int WaysToReachTarget(int target, int[][] types)
    {
        var dp = new long[target + 1];
        dp[0] = 1;
        foreach (var typ in types)
        {
            for (var i = target; i >= 0; i--)
            {
                for (var j = typ[0]; j > 0; j--)
                {
                    if (i + typ[1] * j <= target)
                    {
                        dp[i + typ[1] * j] = (dp[i + typ[1] * j] + dp[i]) % Mod;
                    }
                }
            }
        }
        return (int)dp[target];
    }
}
