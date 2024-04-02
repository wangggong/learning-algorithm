/*
 * @lc app=leetcode.cn id=shortest-path-to-get-all-keys lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [864] 获取所有钥匙的最短路径
 *
 * https://leetcode.cn/problems/shortest-path-to-get-all-keys/description/
 *
 * algorithms
 * Hard (48.24%)
 * Total Accepted:    14.4K
 * Total Submissions: 25K
 * Testcase Example:  '["@.a..","###.#","b.A.B"]'
 *
 * 给定一个二维网格 grid ，其中：
 * 
 * 
 * '.' 代表一个空房间
 * '#' 代表一堵
 * '@' 是起点
 * 小写字母代表钥匙
 * 大写字母代表锁
 * 
 * 
 * 
 * 我们从起点开始出发，一次移动是指向四个基本方向之一行走一个单位空间。我们不能在网格外面行走，也无法穿过一堵墙。如果途经一个钥匙，我们就把它捡起来。除非我们手里有对应的钥匙，否则无法通过锁。
 * 
 * 假设 k 为 钥匙/锁 的个数，且满足 1 <= k <= 6，字母表中的前 k
 * 个字母在网格中都有自己对应的一个小写和一个大写字母。换言之，每个锁有唯一对应的钥匙，每个钥匙也有唯一对应的锁。另外，代表钥匙和锁的字母互为大小写并按字母顺序排列。
 * 
 * 返回获取所有钥匙所需要的移动的最少次数。如果无法获取所有钥匙，返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：grid = ["@.a.#","###.#","b.A.B"]
 * 输出：8
 * 解释：目标是获得所有钥匙，而不是打开所有锁。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：grid = ["@..aA","..B#.","....b"]
 * 输出：6
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: grid = ["@Aa"]
 * 输出: -1
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 30
 * grid[i][j] 只含有 '.', '#', '@', 'a'-'f' 以及 'A'-'F'
 * 钥匙的数目范围是 [1, 6] 
 * 每个钥匙都对应一个 不同 的字母
 * 每个钥匙正好打开一个对应的锁
 * 
 * 
 */
public class Solution
{
    private int[][] Directions = new int[4][]
    {
        new int[] { 0, 1, },
        new int[] { 0, -1, },
        new int[] { 1, 0, },
        new int[] { -1, 0, },
    };

    public int ShortestPathAllKeys(string[] G)
    {
        int n = G.Count();
        int m = G[0].Count();
        int keys = 0;
        int x = 0;
        int y = 0;
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                if ('a' <= G[i][j] && G[i][j] <= 'z') { keys |= 1 << (G[i][j] - 'a'); }
                if (G[i][j] == '@') { x = i; y = j; }
            }
        }
        HashSet<(int, int, int)> vis = new();
        Queue<(int, int, int)> Q = new();
        var start = (x, y, 0);
        Q.Enqueue(start);
        vis.Add(start);
        for (int level = 0; Q.Count() > 0; level++)
        {
            for (int cnt = Q.Count(); cnt > 0; cnt--)
            {
                var q = Q.Dequeue();
                int cx = q.Item1, cy = q.Item2, ck = q.Item3;
                if (ck == keys)
                {
                    return level;
                }
                foreach (var d in Directions)
                {
                    int nx = cx + d[0];
                    int ny = cy + d[1];
                    if (nx < 0 || nx >= n || ny < 0 || ny >= m) { continue; }
                    if (G[nx][ny] == '#') { continue; }
                    if ('A' <= G[nx][ny] && G[nx][ny] <= 'Z' && (ck & (1 << (G[nx][ny] - 'A'))) == 0) { continue; }
                    int nk = ck;
                    if ('a' <= G[nx][ny] && G[nx][ny] <= 'z') { nk |= 1 << (G[nx][ny] - 'a'); }
                    var ns = (nx, ny, nk);
                    if (vis.Contains(ns)) { continue; }
                    Q.Enqueue(ns);
                    vis.Add(ns);
                }
            }
        }
        return -1;
    }
}
