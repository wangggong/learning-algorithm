/*
 * @lc app=leetcode.cn id=parallel-courses-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1494] 并行课程 II
 *
 * https://leetcode-cn.com/problems/parallel-courses-ii/description/
 *
 * algorithms
 * Hard (38.20%)
 * Total Accepted:    3K
 * Total Submissions: 7.9K
 * Testcase Example:  '4\n[[2,1],[3,1],[1,4]]\n2'
 *
 * 给你一个整数 n 表示某所大学里课程的数目，编号为 1 到 n ，数组 dependencies 中， dependencies[i] = [xi,
 * yi]  表示一个先修课的关系，也就是课程 xi 必须在课程 yi 之前上。同时你还有一个整数 k 。
 * 
 * 在一个学期中，你 最多 可以同时上 k 门课，前提是这些课的先修课在之前的学期里已经上过了。
 * 
 * 请你返回上完所有课最少需要多少个学期。题目保证一定存在一种上完所有课的方式。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：n = 4, dependencies = [[2,1],[3,1],[1,4]], k = 2
 * 输出：3 
 * 解释：上图展示了题目输入的图。在第一个学期中，我们可以上课程 2 和课程 3 。然后第二个学期上课程 1 ，第三个学期上课程 4 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：n = 5, dependencies = [[2,1],[3,1],[4,1],[1,5]], k = 2
 * 输出：4 
 * 解释：上图展示了题目输入的图。一个最优方案是：第一学期上课程 2 和 3，第二学期上课程 4 ，第三学期上课程 1 ，第四学期上课程 5 。
 * 
 * 
 * 示例 3：
 * 
 * 输入：n = 11, dependencies = [], k = 2
 * 输出：6
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 15
 * 1 <= k <= n
 * 0 <= dependencies.length <= n * (n-1) / 2
 * dependencies[i].length == 2
 * 1 <= xi, yi <= n
 * xi != yi
 * 所有先修关系都是不同的，也就是说 dependencies[i] != dependencies[j] 。
 * 题目输入的图是个有向无环图。
 * 
 * 
 */
public class Solution
{
    private const int Maxn = 0x3f3f3f3f;

    public int MinNumberOfSemesters(int n, int[][] relations, int k)
    {
        int count1(int n)
        {
            var ans = 0;
            for (; n > 0; n &= n - 1)
            {
                ans++;
            }
            return ans;
        }
        var pre = new int[n];
        Array.ForEach(relations, r => pre[r[0] - 1] |= 1 << (r[1] - 1));
        var limit = 1 << n;
        var dp = Enumerable
            .Range(0, limit)
            .Select(_ => Maxn)
            .ToArray();
        dp[0] = 0;
        for (var s = 0; s < limit; s++)
        {
            var canLearn = Enumerable
                .Range(0, n)
                .Where(x => (s & pre[x]) == pre[x])
                .Select(x => 1 << x)
                .Sum() & (limit - 1 - s);
            for (var learn = canLearn; learn > 0; learn = (learn - 1) & canLearn)
            {
                if (count1(learn) <= k)
                {
                    dp[s | learn] = Math.Min(dp[s | learn], dp[s] + 1);
                }
            }
        }
        return dp[limit - 1];
    }
}
