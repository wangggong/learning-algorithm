/*
 * @lc app=leetcode.cn id=count-number-of-possible-root-nodes lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6314] 统计可能的树根数目
 *
 * https://leetcode.cn/problems/count-number-of-possible-root-nodes/description/
 *
 * algorithms
 * Hard (53.35%)
 * Total Accepted:    1.4K
 * Total Submissions: 2.5K
 * Testcase Example:  '[[0,1],[1,2],[1,3],[4,2]]\n[[1,3],[0,1],[1,0],[2,4]]\n3'
 *
 * Alice 有一棵 n 个节点的树，节点编号为 0 到 n - 1 。树用一个长度为 n - 1 的二维整数数组 edges 表示，其中
 * edges[i] = [ai, bi] ，表示树中节点 ai 和 bi 之间有一条边。
 * 
 * Alice 想要 Bob 找到这棵树的根。她允许 Bob 对这棵树进行若干次 猜测 。每一次猜测，Bob 做如下事情：
 * 
 * 
 * 选择两个 不相等 的整数 u 和 v ，且树中必须存在边 [u, v] 。
 * Bob 猜测树中 u 是 v 的 父节点 。
 * 
 * 
 * Bob 的猜测用二维整数数组 guesses 表示，其中 guesses[j] = [uj, vj] 表示 Bob 猜 uj 是 vj 的父节点。
 * 
 * Alice 非常懒，她不想逐个回答 Bob 的猜测，只告诉 Bob 这些猜测里面 至少 有 k 个猜测的结果为 true 。
 * 
 * 给你二维整数数组 edges ，Bob 的所有猜测和整数 k ，请你返回可能成为树根的 节点数目 。如果没有这样的树，则返回 0。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：edges = [[0,1],[1,2],[1,3],[4,2]], guesses = [[1,3],[0,1],[1,0],[2,4]], k
 * = 3
 * 输出：3
 * 解释：
 * 根为节点 0 ，正确的猜测为 [1,3], [0,1], [2,4]
 * 根为节点 1 ，正确的猜测为 [1,3], [1,0], [2,4]
 * 根为节点 2 ，正确的猜测为 [1,3], [1,0], [2,4]
 * 根为节点 3 ，正确的猜测为 [1,0], [2,4]
 * 根为节点 4 ，正确的猜测为 [1,3], [1,0]
 * 节点 0 ，1 或 2 为根时，可以得到 3 个正确的猜测。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：edges = [[0,1],[1,2],[2,3],[3,4]], guesses = [[1,0],[3,4],[2,1],[3,2]], k
 * = 1
 * 输出：5
 * 解释：
 * 根为节点 0 ，正确的猜测为 [3,4]
 * 根为节点 1 ，正确的猜测为 [1,0], [3,4]
 * 根为节点 2 ，正确的猜测为 [1,0], [2,1], [3,4]
 * 根为节点 3 ，正确的猜测为 [1,0], [2,1], [3,2], [3,4]
 * 根为节点 4 ，正确的猜测为 [1,0], [2,1], [3,2]
 * 任何节点为根，都至少有 1 个正确的猜测。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * edges.length == n - 1
 * 2 <= n <= 10^5
 * 1 <= guesses.length <= 10^5
 * 0 <= ai, bi, uj, vj <= n - 1
 * ai != bi
 * uj != vj
 * edges 表示一棵有效的树。
 * guesses[j] 是树中的一条边。
 * 0 <= k <= guesses.length
 * 
 * 
 */
public class Solution
{
    public int RootCount(int[][] edges, int[][] guesses, int k)
    {
        var n = edges.Length + 1;
        var G = new List<int>[n];
        for (var i = 0; i < n; i++)
        {
            G[i] = new();
        }
        foreach (var e in edges)
        {
            G[e[0]].Add(e[1]);
            G[e[1]].Add(e[0]);
        }
        var count = new Dictionary<(int, int), int>();
        foreach (var g in guesses)
        {
            var key = (g[0], g[1]);
            count.TryGetValue(key, out var v);
            count[key] = v + 1;
        }
        var init = 0;
        void initDfs(int u, int p)
        {
            foreach (var v in G[u])
            {
                if (v != p)
                {
                    count.TryGetValue((u, v), out var c);
                    init += c;
                    initDfs(v, u);
                }
            }
        }
        initDfs(0, -1);
        var tot = 0;
        void dfs(int u, int p, int d)
        {
            if (init + d >= k)
            {
                tot++;
            }
            foreach (var v in G[u])
            {
                if (v != p)
                {
                    count.TryGetValue((u, v), out var c1);
                    count.TryGetValue((v, u), out var c2);
                    dfs(v, u, d - c1 + c2);
                }
            }
        }
        dfs(0, -1, 0);
        return tot;
    }
}
