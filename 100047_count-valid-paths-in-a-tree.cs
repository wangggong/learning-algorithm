/*
 * @lc app=leetcode.cn id=count-valid-paths-in-a-tree lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100047] 统计树中的合法路径数目
 *
 * https://leetcode.cn/problems/count-valid-paths-in-a-tree/description/
 *
 * algorithms
 * Hard (23.72%)
 * Total Accepted:    1K
 * Total Submissions: 4.3K
 * Testcase Example:  '5\n[[1,2],[1,3],[2,4],[2,5]]'
 *
 * 给你一棵 n 个节点的无向树，节点编号为 1 到 n 。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[i]
 * = [ui, vi] 表示节点 ui 和 vi 在树中有一条边。
 * 
 * 请你返回树中的 合法路径数目 。
 * 
 * 如果在节点 a 到节点 b 之间 恰好有一个 节点的编号是质数，那么我们称路径 (a, b) 是 合法的 。
 * 
 * 注意：
 * 
 * 
 * 路径 (a, b) 指的是一条从节点 a 开始到节点 b 结束的一个节点序列，序列中的节点 互不相同 ，且相邻节点之间在树上有一条边。
 * 路径 (a, b) 和路径 (b, a) 视为 同一条 路径，且只计入答案 一次 。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：n = 5, edges = [[1,2],[1,3],[2,4],[2,5]]
 * 输出：4
 * 解释：恰好有一个质数编号的节点路径有：
 * - (1, 2) 因为路径 1 到 2 只包含一个质数 2 。
 * - (1, 3) 因为路径 1 到 3 只包含一个质数 3 。
 * - (1, 4) 因为路径 1 到 4 只包含一个质数 2 。
 * - (2, 4) 因为路径 2 到 4 只包含一个质数 2 。
 * 只有 4 条合法路径。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：n = 6, edges = [[1,2],[1,3],[2,4],[3,5],[3,6]]
 * 输出：6
 * 解释：恰好有一个质数编号的节点路径有：
 * - (1, 2) 因为路径 1 到 2 只包含一个质数 2 。
 * - (1, 3) 因为路径 1 到 3 只包含一个质数 3 。
 * - (1, 4) 因为路径 1 到 4 只包含一个质数 2 。
 * - (1, 6) 因为路径 1 到 6 只包含一个质数 3 。
 * - (2, 4) 因为路径 2 到 4 只包含一个质数 2 。
 * - (3, 6) 因为路径 3 到 6 只包含一个质数 3 。
 * 只有 6 条合法路径。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * edges.length == n - 1
 * edges[i].length == 2
 * 1 <= ui, vi <= n
 * 输入保证 edges 形成一棵合法的树。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/count-valid-paths-in-a-tree/solutions/2456716/tu-jie-on-xian-xing-zuo-fa-pythonjavacgo-tjz2/
//
// 忘了记忆化搜索了.
// 同时, `dfs` 的入口选择哪个也很重要.
public class Solution
{
    private const int N = (int)1e5;
    private static bool[] isPrimes;

    static Solution()
    {
        isPrimes = new bool[N + 1];
        Array.Fill(isPrimes, true);
        isPrimes[1] = false;
        for (var i = 2; i <= N; i++)
        {
            if (!isPrimes[i]) { continue; }
            for (var j = i; j <= N / i; j++) { isPrimes[i * j] = false; }
        }
    }

    public long CountPaths(int n, int[][] edges)
    {
        var G = new List<int>[n + 1];
        for (var i = 1; i <= n; i++) { G[i] = new(); }
        foreach (var edge in edges)
        {
            var (u, v) = (edge[0], edge[1]);
            G[u].Add(v);
            G[v].Add(u);
        }
        var memo = new Dictionary<(int, int), long>();
        long dfs(int u, int p)
        {
            var key = (u, p);
            if (memo.ContainsKey(key)) { return memo[key]; }
            if (isPrimes[u]) { return memo[key] = 0; }
            memo[key] = 1;
            foreach (var v in G[u])
            {
                if (v == p) { continue; }
                memo[key] += dfs(v, u);
            }
            return memo[key];
        }
        var ans = 0L;
        for (var u = 1; u <= n; u++)
        {
            if (!isPrimes[u]) { continue; }
            var counts = G[u]
                .Select(v => dfs(v, -1))
                .ToList();
            var total = counts.Sum() + 1;
            ans += total * (total - 1) / 2 - counts
                .Select(c => c * (c - 1) / 2)
                .Sum();
        }
        return ans;
    }
}
