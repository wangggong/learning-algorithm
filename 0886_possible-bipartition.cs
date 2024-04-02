/*
 * @lc app=leetcode.cn id=possible-bipartition lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [886] 可能的二分法
 *
 * https://leetcode.cn/problems/possible-bipartition/description/
 *
 * algorithms
 * Medium (49.75%)
 * Total Accepted:    31.3K
 * Total Submissions: 62.4K
 * Testcase Example:  '4\n[[1,2],[1,3],[2,4]]'
 *
 * 给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。
 * 
 * 给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和
 * bi的人归入同一组。当可以用这种方法将所有人分进两组时，返回 true；否则返回 false。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
 * 输出：true
 * 解释：group1 [1,4], group2 [2,3]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
 * 输出：false
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 2000
 * 0 <= dislikes.length <= 10^4
 * dislikes[i].length == 2
 * 1 <= dislikes[i][j] <= n
 * ai < bi
 * dislikes 中每一组都 不同
 * 
 * 
 * 
 * 
 */
public class Solution
{
    public bool PossibleBipartition(int n, int[][] ds)
    {
        IList<int>[] g = new IList<int>[n + 1];
        for (int i = 1; i <= n; i++)
        {
            g[i] = new List<int>();
        }
        foreach (var d in ds)
        {
            g[d[0]].Add(d[1]);
            g[d[1]].Add(d[0]);
        }
        int[] cols = new int[n + 1];
        for (int i = 1; i <= n; i++)
        {
            if (cols[i] == 0 && !Bfs(g, cols, i))
            {
                return false;
            }
        }
        return true;
    }

    private bool Bfs(IList<IList<int>> g, IList<int> cols, int k)
    {
        Queue<int> Q = new ();
        cols[k] = 1;
        Q.Enqueue(k);
        while (Q.Count != 0)
        {
            var u = Q.Dequeue();
            foreach (var v in g[u])
            {
                if (cols[v] != 0)
                {
                    if (cols[u] == cols[v])
                    {
                        return false;
                    }
                    continue;
                }
                cols[v] = -cols[u];
                Q.Enqueue(v);
            }
        }
        return true;
    }
}
