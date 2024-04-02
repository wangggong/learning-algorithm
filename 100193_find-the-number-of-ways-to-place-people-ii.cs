/*
 * @lc app=leetcode.cn id=find-the-number-of-ways-to-place-people-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100193] 人员站位的方案数 II
 *
 * https://leetcode.cn/problems/find-the-number-of-ways-to-place-people-ii/description/
 *
 * algorithms
 * Hard (51.99%)
 * Total Accepted:    1.3K
 * Total Submissions: 2.5K
 * Testcase Example:  '[[1,1],[2,2],[3,3]]'
 *
 * 给你一个  n x 2 的二维数组 points ，它表示二维平面上的一些点坐标，其中 points[i] = [xi, yi] 。
 * 
 * 我们定义 x 轴的正方向为 右 （x 轴递增的方向），x 轴的负方向为 左 （x 轴递减的方向）。类似的，我们定义 y 轴的正方向为 上 （y
 * 轴递增的方向），y 轴的负方向为 下 （y 轴递减的方向）。
 * 
 * 你需要安排这 n 个人的站位，这 n 个人中包括 liupengsay 和小羊肖恩 。你需要确保每个点处 恰好 有 一个 人。同时，liupengsay
 * 想跟小羊肖恩单独玩耍，所以 liupengsay 会以 liupengsay 的坐标为 左上角 ，小羊肖恩的坐标为 右下角
 * 建立一个矩形的围栏（注意，围栏可能 不 包含任何区域，也就是说围栏可能是一条线段）。如果围栏的 内部 或者 边缘 上有任何其他人，liupengsay
 * 都会难过。
 * 
 * 请你在确保 liupengsay 不会 难过的前提下，返回 liupengsay 和小羊肖恩可以选择的 点对 数目。
 * 
 * 注意，liupengsay 建立的围栏必须确保 liupengsay 的位置是矩形的左上角，小羊肖恩的位置是矩形的右下角。比方说，以 (1, 1)
 * ，(1, 3) ，(3, 1) 和 (3, 3) 为矩形的四个角，给定下图的两个输入，liupengsay 都不能建立围栏，原因如下：
 * 
 * 
 * 图一中，liupengsay 在 (3, 3) 且小羊肖恩在 (1, 1) ，liupengsay 的位置不是左上角且小羊肖恩的位置不是右下角。
 * 图二中，liupengsay 在 (1, 3) 且小羊肖恩在 (1, 1) ，小羊肖恩的位置不是在围栏的右下角。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：points = [[1,1],[2,2],[3,3]]
 * 输出：0
 * 解释：没有办法可以让 liupengsay 的围栏以 liupengsay 的位置为左上角且小羊肖恩的位置为右下角。所以我们返回 0 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：points = [[6,2],[4,4],[2,6]]
 * 输出：2
 * 解释：总共有 2 种方案安排 liupengsay 和小羊肖恩的位置，使得 liupengsay 不会难过：
 * - liupengsay 站在 (4, 4) ，小羊肖恩站在 (6, 2) 。
 * - liupengsay 站在 (2, 6) ，小羊肖恩站在 (4, 4) 。
 * 不能安排 liupengsay 站在 (2, 6) 且小羊肖恩站在 (6, 2) ，因为站在 (4, 4) 的人处于围栏内。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 
 * 输入：points = [[3,1],[1,3],[1,1]]
 * 输出：2
 * 解释：总共有 2 种方案安排 liupengsay 和小羊肖恩的位置，使得 liupengsay 不会难过：
 * - liupengsay 站在 (1, 1) ，小羊肖恩站在 (3, 1) 。
 * - liupengsay 站在 (1, 3) ，小羊肖恩站在 (1, 1) 。
 * 不能安排 liupengsay 站在 (1, 3) 且小羊肖恩站在 (3, 1) ，因为站在 (1, 1) 的人处于围栏内。
 * 注意围栏是可以不包含任何面积的，上图中第一和第二个围栏都是合法的。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n <= 1000
 * points[i].length == 2
 * -10^9 <= points[i][0], points[i][1] <= 10^9
 * points[i] 点对两两不同。
 * 
 * 
 */
public class Solution
{
    public int NumberOfPairs(int[][] points)
    {
        const int Maxn = 0x3f3f3f3f;
        var ans = 0;
        var positions = points.GroupBy(p => p[1])
            .OrderByDescending(g => g.Key)
            .Select(g => g.Select(p => p[0])
                .OrderBy(x => x)
                .ToArray())
            .ToArray();
        var n = positions.Length;
        foreach (var (xs, i) in positions.Select((xs, i) => (xs, i)))
        {
            foreach (var (x, j) in xs.Select((x, j) => (x, j)))
            {
                var maxX = Maxn;
                if (j + 1 < xs.Length)
                {
                    maxX = xs[j + 1];
                    ans++;
                }
                for (var k = i + 1; k < n; k++)
                {
                    var (p, q) = (0, positions[k].Length);
                    while (p < q)
                    {
                        var mid = (p + q) >> 1;
                        if (positions[k][mid] >= x)
                        {
                            q = mid;
                        }
                        else
                        {
                            p = mid + 1;
                        }
                    }
                    if (p == positions[k].Length || positions[k][p] >= maxX)
                    {
                        continue;
                    }
                    maxX = positions[k][p];
                    ans++;
                    if (positions[k][p] == x)
                    {
                        break;
                    }
                }
            }
        }
        return ans;
    }
}

// 参考: https://leetcode.cn/problems/find-the-number-of-ways-to-place-people-ii/solutions/2630655/on2-you-ya-mei-ju-by-endlesscheng-z86d/

/*
 * public class Solution
 * {
 *     public int NumberOfPairs(int[][] points)
 *     {
 *         points = points
 *             .OrderByDescending(p => p[1])
 *             .ThenBy(p => p[0])
 *             .ToArray();
 *         var ans = 0;
 *         foreach (var (p, i) in points.Select((p, i) => (p, i)))
 *         {
 *             var maxX = int.MaxValue;
 *             for (var (j, n) = (i + 1, points.Length); j < n; j++)
 *             {
 *                 if (p[0] <= points[j][0] && points[j][0] < maxX)
 *                 {
 *                     maxX = points[j][0];
 *                     ans++;
 *                 }
 *             }
 *         }
 *         return ans;
 *     }
 * }
 */
