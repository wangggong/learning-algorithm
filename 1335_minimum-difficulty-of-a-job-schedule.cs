/*
 * @lc app=leetcode.cn id=minimum-difficulty-of-a-job-schedule lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1335] 工作计划的最低难度
 *
 * https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/description/
 *
 * algorithms
 * Hard (60.53%)
 * Total Accepted:    6K
 * Total Submissions: 9.9K
 * Testcase Example:  '[6,5,4,3,2,1]\n2'
 *
 * 你需要制定一份 d 天的工作计划表。工作之间存在依赖，要想执行第 i 项工作，你必须完成全部 j 项工作（ 0 <= j < i）。
 * 
 * 你每天 至少 需要完成一项任务。工作计划的总难度是这 d 天每一天的难度之和，而一天的工作难度是当天应该完成工作的最大难度。
 * 
 * 给你一个整数数组 jobDifficulty 和一个整数 d，分别代表工作难度和需要计划的天数。第 i 项工作的难度是
 * jobDifficulty[i]。
 * 
 * 返回整个工作计划的 最小难度 。如果无法制定工作计划，则返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：jobDifficulty = [6,5,4,3,2,1], d = 2
 * 输出：7
 * 解释：第一天，您可以完成前 5 项工作，总难度 = 6.
 * 第二天，您可以完成最后一项工作，总难度 = 1.
 * 计划表的难度 = 6 + 1 = 7 
 * 
 * 
 * 示例 2：
 * 
 * 输入：jobDifficulty = [9,9,9], d = 4
 * 输出：-1
 * 解释：就算你每天完成一项工作，仍然有一天是空闲的，你无法制定一份能够满足既定工作时间的计划表。
 * 
 * 
 * 示例 3：
 * 
 * 输入：jobDifficulty = [1,1,1], d = 3
 * 输出：3
 * 解释：工作计划为每天一项工作，总难度为 3 。
 * 
 * 
 * 示例 4：
 * 
 * 输入：jobDifficulty = [7,1,7,1,7,1], d = 3
 * 输出：15
 * 
 * 
 * 示例 5：
 * 
 * 输入：jobDifficulty = [11,111,22,222,33,333,44,444], d = 6
 * 输出：843
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= jobDifficulty.length <= 300
 * 0 <= jobDifficulty[i] <= 1000
 * 1 <= d <= 10
 * 
 * 
 */
public class Solution
{
    private const int Maxn = 0x3f3f3f3f;

    public int MinDifficulty(int[] difficulties, int d)
    {
        var n = difficulties.Length;
        if (n < d)
        {
            return -1;
        }
        var dp = Enumerable.Range(0, 2).Select(_ => new int[n]).ToArray();
        for (var (i, j) = (0, 0); i < n; i++)
        {
            j = Math.Max(j, difficulties[i]);
            dp[0][i] = j;
        }
        for (var i = 1; i < d; i++)
        {
            Array.Fill(dp[i & 1], Maxn);
            for (var j = i; j < n; j++)
            {
                for (var (k, max) = (j, 0); k >= i; k--)
                {
                    max = Math.Max(max, difficulties[k]);
                    dp[i & 1][j] = Math.Min(
                        dp[i & 1][j], dp[(i - 1) & 1][k - 1] + max);
                }
            }
        }
        return dp[(d - 1) & 1][n - 1];
    }
}

/*
 * public class Solution
 * {
 *     private const int Maxn = 0x3f3f3f3f;
 * 
 *     public int MinDifficulty(int[] difficulties, int d)
 *     {
 *         var n = difficulties.Length;
 *         if (n < d)
 *         {
 *             return -1;
 *         }
 *         var mx = Enumerable.Range(0, n).Select(_ => new int[n]).ToArray();
 *         for (var i = 0; i < n; i++)
 *         {
 *             mx[i][i] = difficulties[i];
 *             for (var j = i + 1; j < n; j++)
 *             {
 *                 mx[i][j] = Math.Max(mx[i][j - 1], difficulties[j]);
 *             }
 *         }
 *         var dp = Enumerable.Range(0, d + 1).Select(_ =>
 *         {
 *             var ans = new int[n + 1];
 *             Array.Fill(ans, Maxn);
 *             return ans;
 *         }).ToArray();
 *         dp[0][0] = 0;
 *         for (var i = 1; i <= d; i++)
 *         {
 *             for (var j = 0; j < n; j++)
 *             {
 *                 for (var k = 0; k <= j; k++)
 *                 {
 *                     dp[i][j + 1] = Math.Min(dp[i][j + 1], dp[i - 1][k] + mx[k][j]);
 *                 }
 *             }
 *         }
 *         return dp[d][n];
 *     }
 * }
 */
