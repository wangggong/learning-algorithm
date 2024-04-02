/*
 * @lc app=leetcode.cn id=count-visited-nodes-in-a-directed-graph lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100075] 有向图访问计数
 *
 * https://leetcode.cn/problems/count-visited-nodes-in-a-directed-graph/description/
 *
 * algorithms
 * Hard (34.72%)
 * Total Accepted:    1.1K
 * Total Submissions: 3K
 * Testcase Example:  '[1,2,0,0]'
 *
 * 现有一个有向图，其中包含 n 个节点，节点编号从 0 到 n - 1 。此外，该图还包含了 n 条有向边。
 * 
 * 给你一个下标从 0 开始的数组 edges ，其中 edges[i] 表示存在一条从节点 i 到节点 edges[i] 的边。
 * 
 * 想象在图上发生以下过程：
 * 
 * 
 * 你从节点 x 开始，通过边访问其他节点，直到你在 此过程 中再次访问到之前已经访问过的节点。
 * 
 * 
 * 返回数组 answer 作为答案，其中 answer[i] 表示如果从节点 i 开始执行该过程，你可以访问到的不同节点数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：edges = [1,2,0,0]
 * 输出：[3,3,3,4]
 * 解释：从每个节点开始执行该过程，记录如下：
 * - 从节点 0 开始，访问节点 0 -> 1 -> 2 -> 0 。访问的不同节点数是 3 。
 * - 从节点 1 开始，访问节点 1 -> 2 -> 0 -> 1 。访问的不同节点数是 3 。
 * - 从节点 2 开始，访问节点 2 -> 0 -> 1 -> 2 。访问的不同节点数是 3 。
 * - 从节点 3 开始，访问节点 3 -> 0 -> 1 -> 2 -> 0 。访问的不同节点数是 4 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：edges = [1,2,3,4,0]
 * 输出：[5,5,5,5,5]
 * 解释：无论从哪个节点开始，在这个过程中，都可以访问到图中的每一个节点。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == edges.length
 * 2 <= n <= 10^5
 * 0 <= edges[i] <= n - 1
 * edges[i] != i
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/count-visited-nodes-in-a-directed-graph/solutions/2464852/nei-xiang-ji-huan-shu-pythonjavacgo-by-e-zrzh/
//
// 基环树.
// 基础不牢, 地动山摇
public class Solution
{
    public int[] CountVisitedNodes(IList<int> edges)
    {
        var n = edges.Count();
        var G = new List<int>[n];
        for (var i = 0; i < n; i++) { G[i] = new(); }
        var degrees = new int[n];
        foreach (var (v, u) in edges
            .Select((e, i) => (e, i)))
        {
            G[v].Add(u);
            degrees[v]++;
        }
        var Q = new Queue<int>();
        for (var i = 0; i < n; i++) { if (degrees[i] is 0) { Q.Enqueue(i); } }
        while (Q.Count > 0)
        {
            var u = Q.Dequeue();
            var v = edges[u];
            degrees[v]--;
            if (degrees[v] is 0) { Q.Enqueue(v); }
        }
        var ans = new int[n];
        void dfs(int u, int k)
        {
            ans[u] = k;
            foreach (var v in G[u])
            {
                if (degrees[v] is not 0) { continue; }
                dfs(v, k + 1);
            }
        }
        for (var i = 0; i < n; i++)
        {
            if (degrees[i] <= 0) { continue; }
            degrees[i] = -1;
            var nodes = new List<int> { i, };
            for (var u = edges[i]; u != i; u = edges[u])
            {
                degrees[u] = -1;
                nodes.Add(u);
            }
            foreach (var u in nodes) { dfs(u, nodes.Count()); }
        }
        return ans;
    }
}
