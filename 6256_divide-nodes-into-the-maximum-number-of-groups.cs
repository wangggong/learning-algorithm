/*
 * @lc app=leetcode.cn id=divide-nodes-into-the-maximum-number-of-groups lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6256] 将节点分成尽可能多的组
 *
 * https://leetcode.cn/problems/divide-nodes-into-the-maximum-number-of-groups/description/
 *
 * algorithms
 * Hard (21.62%)
 * Total Accepted:    654
 * Total Submissions: 2.6K
 * Testcase Example:  '6\n[[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]]'
 *
 * 给你一个正整数 n ，表示一个 无向 图中的节点数目，节点编号从 1 到 n 。
 * 
 * 同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条 双向
 * 边。注意给定的图可能是不连通的。
 * 
 * 请你将图划分为 m 个组（编号从 1 开始），满足以下要求：
 * 
 * 
 * 图中每个节点都只属于一个组。
 * 图中每条边连接的两个点 [ai, bi] ，如果 ai 属于编号为 x 的组，bi 属于编号为 y 的组，那么 |y - x| = 1 。
 * 
 * 
 * 请你返回最多可以将节点分为多少个组（也就是最大的 m ）。如果没办法在给定条件下分组，请你返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：n = 6, edges = [[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]]
 * 输出：4
 * 解释：如上图所示，
 * - 节点 5 在第一个组。
 * - 节点 1 在第二个组。
 * - 节点 2 和节点 4 在第三个组。
 * - 节点 3 和节点 6 在第四个组。
 * 所有边都满足题目要求。
 * 如果我们创建第五个组，将第三个组或者第四个组中任何一个节点放到第五个组，至少有一条边连接的两个节点所属的组编号不符合题目要求。
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 3, edges = [[1,2],[2,3],[3,1]]
 * 输出：-1
 * 解释：如果我们将节点 1 放入第一个组，节点 2 放入第二个组，节点 3 放入第三个组，前两条边满足题目要求，但第三条边不满足题目要求。
 * 没有任何符合题目要求的分组方式。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 500
 * 1 <= edges.length <= 10^4
 * edges[i].length == 2
 * 1 <= ai, bi <= n
 * ai != bi
 * 两个点之间至多只有一条边。
 * 
 * 
 */
public class Solution
{
    public int MagnificentSets(int n, int[][] edges)
    {
        var G = new List<int>[n + 1];
        for (var i = 1; i <= n; i++) { G[i] = new(); }
        foreach (var e in edges)
        {
            G[e[0]].Add(e[1]);
            G[e[1]].Add(e[0]);
        }
        var visit = new bool[n + 1];
        var ans = 0;
        int find(int k)
        {
            var Q = new Queue<int>();
            var all = new List<int>();
            Q.Enqueue(k);
            visit[k] = true;
            all.Add(k);
            while (Q.Count > 0)
            {
                var u = Q.Dequeue();
                foreach (var v in G[u])
                {
                    if (!visit[v])
                    {
                        Q.Enqueue(v);
                        visit[v] = true;
                        all.Add(v);
                    }
                }
            }
            var maxLevel = 0;
            foreach (var a in all)
            {
                var level = 0;
                var dist = new Dictionary<int, int>();
                Q.Enqueue(a);
                dist[a] = 0;
                while (Q.Count > 0)
                {
                    for (var c = Q.Count; c > 0; c--)
                    {
                        var u = Q.Dequeue();
                        foreach (var v in G[u])
                        {
                            if (dist.ContainsKey(v))
                            {
                                if (dist[v] != level + 1 && dist[v] != level - 1) { return -1; }
                                continue;
                            }
                            Q.Enqueue(v);
                            dist[v] = level + 1;
                        }
                    }
                    level++;
                }
                maxLevel = Math.Max(maxLevel, level);
            }
            return maxLevel;
        }
        for (var i = 1; i <= n; i++)
        {
            if (!visit[i])
            {
                var cur = find(i);
                if (cur == -1) { return -1; }
                ans += cur;
            }
        }
        return ans;
    }
}
