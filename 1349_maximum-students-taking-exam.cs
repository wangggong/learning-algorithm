/*
 * @lc app=leetcode.cn id=maximum-students-taking-exam lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1349] 参加考试的最大学生数
 *
 * https://leetcode.cn/problems/maximum-students-taking-exam/description/
 *
 * algorithms
 * Hard (54.98%)
 * Total Accepted:    13.2K
 * Total Submissions: 21.1K
 * Testcase Example:  '[["#",".","#","#",".","#"],[".","#","#","#","#","."],["#",".","#","#",".","#"]]'
 *
 * 给你一个 m * n 的矩阵 seats 表示教室中的座位分布。如果座位是坏的（不可用），就用 '#' 表示；否则，用 '.' 表示。
 * 
 * 
 * 学生可以看到左侧、右侧、左上、右上这四个方向上紧邻他的学生的答卷，但是看不到直接坐在他前面或者后面的学生的答卷。请你计算并返回该考场可以容纳的同时参加考试且无法作弊的
 * 最大 学生人数。
 * 
 * 学生必须坐在状况良好的座位上。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：seats = [["#",".","#","#",".","#"],
 * [".","#","#","#","#","."],
 * ["#",".","#","#",".","#"]]
 * 输出：4
 * 解释：教师可以让 4 个学生坐在可用的座位上，这样他们就无法在考试中作弊。 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：seats = [[".","#"],
 * ["#","#"],
 * ["#","."],
 * ["#","#"],
 * [".","#"]]
 * 输出：3
 * 解释：让所有学生坐在可用的座位上。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：seats = [["#",".",".",".","#"],
 * [".","#",".","#","."],
 * [".",".","#",".","."],
 * [".","#",".","#","."],
 * ["#",".",".",".","#"]]
 * 输出：10
 * 解释：让学生坐在第 1、3 和 5 列的可用座位上。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * seats 只包含字符 '.' 和'#'
 * m == seats.length
 * n == seats[i].length
 * 1 <= m <= 8
 * 1 <= n <= 8
 * 
 * 
 */
public class Solution
{
    public int MaxStudents(char[][] seats)
    {
        const int N = 8;
        int index(int i, int j) => i * N + j;
        bool setBit(long k, int d) => ((k >> d) & 1) is not 0;
        var (n, m) = (seats.Length, seats[0].Length);
        int lastSetBit(long k) => Enumerable.Range(0, N * N)
            .First(d => setBit(k, d));
        long next(long k, int d, bool sit)
        {
            k -= 1l << d;
            if (!sit) { return k; }
            var (i, j) = (d / N, d % N);
            if (j > 0)
            {
                if (setBit(k, index(i, j - 1))) { k -= 1l << index(i, j - 1); }
                if (i + 1 < n && setBit(k, index(i + 1, j - 1)))
                { k -= 1l << index(i + 1, j - 1); }
            }
            if (j + 1 < m)
            {
                if (setBit(k, index(i, j + 1))) { k -= 1l << index(i, j + 1); }
                if (i + 1 < n && setBit(k, index(i + 1, j + 1)))
                { k -= 1l << index(i + 1, j + 1); }
            }
            return k;
        }
        var memo = new Dictionary<long, int>();
        int dfs(long k) => memo[k] = memo.ContainsKey(k)
            ? memo[k]
            : (k is 0 ? 0 : Math.Max(
                dfs(next(k, lastSetBit(k), false)),
                dfs(next(k, lastSetBit(k), true)) + 1));
        return dfs(seats
            .SelectMany((row, i) => row
                .Select((c, j) => (i, j, c)))
            .Where(x => x.c is '.')
            .Select(x => 1l << index(x.i, x.j))
            .Sum());
    }
}
