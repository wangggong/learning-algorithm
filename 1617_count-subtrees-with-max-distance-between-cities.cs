/*
 * @lc app=leetcode.cn id=count-subtrees-with-max-distance-between-cities lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1617] 统计子树中城市之间最大距离
 *
 * https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities/description/
 *
 * algorithms
 * Hard (77.10%)
 * Total Accepted:    7.7K
 * Total Submissions: 9.9K
 * Testcase Example:  '4\n[[1,2],[2,3],[2,4]]'
 *
 * 给你 n 个城市，编号为从 1 到 n 。同时给你一个大小为 n-1 的数组 edges ，其中 edges[i] = [ui, vi] 表示城市 ui
 * 和 vi 之间有一条双向边。题目保证任意城市之间只有唯一的一条路径。换句话说，所有城市形成了一棵 树 。
 * 
 * 一棵 子树
 * 是城市的一个子集，且子集中任意城市之间可以通过子集中的其他城市和边到达。两个子树被认为不一样的条件是至少有一个城市在其中一棵子树中存在，但在另一棵子树中不存在。
 * 
 * 对于 d 从 1 到 n-1 ，请你找到城市间 最大距离 恰好为 d 的所有子树数目。
 * 
 * 请你返回一个大小为 n-1 的数组，其中第 d 个元素（下标从 1 开始）是城市间 最大距离 恰好等于 d 的子树数目。
 * 
 * 请注意，两个城市间距离定义为它们之间需要经过的边的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：n = 4, edges = [[1,2],[2,3],[2,4]]
 * 输出：[3,4,0]
 * 解释：
 * 子树 {1,2}, {2,3} 和 {2,4} 最大距离都是 1 。
 * 子树 {1,2,3}, {1,2,4}, {2,3,4} 和 {1,2,3,4} 最大距离都为 2 。
 * 不存在城市间最大距离为 3 的子树。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 2, edges = [[1,2]]
 * 输出：[1]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 3, edges = [[1,2],[2,3]]
 * 输出：[2,1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 
 * edges.length == n-1
 * edges[i].length == 2
 * 1 i, vi 
 * 题目保证 (ui, vi) 所表示的边互不相同。
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities/solution/tu-jie-on3-mei-ju-zhi-jing-duan-dian-che-am2n/
//
// 1. 二进制枚举
// 2. 树的直径
public class Solution
{
    public int[] CountSubgraphsForEachDiameter(int n, int[][] edges)
    {
        // 建图
        var G = new List<int>[n];
        for (var i = 0; i < n; i++)
        {
            G[i] = new();
        }
        foreach (var edge in edges)
        {
            var (u, v) = (edge[0] - 1, edge[1] - 1);
            G[u].Add(v);
            G[v].Add(u);
        }
        var ans = new int[n];
        for (var mask = 0; mask < 1 << n; mask++)
        {
            // 至少要两个点
            if ((mask & (mask - 1)) != 0)
            {
                var visit = 0;
                var diameter = 0;
                int dfs(int u)
                {
                    visit |= 1 << u;
                    // 子树最大长度
                    var maxLen = 0;
                    foreach (var v in G[u])
                    {
                        if ((mask & (1 << v)) != 0 && (visit & (1 << v)) == 0)
                        {
                            var cur = dfs(v) + 1;
                            // 两段不同子树的长度和更新全局直径
                            diameter = Math.Max(diameter, maxLen + cur);
                            // 当前子树长度更新子树最大长度
                            maxLen = Math.Max(maxLen, cur);
                        }
                    }
                    return maxLen;
                }
                var k = 0;
                for (; (mask & (1 << k)) == 0; k++) { }
                dfs(k);
                // 枚举状态后, 要判断这是一个 "子树" 而非 "森林".
                if (visit == mask)
                {
                    ans[diameter]++;
                }
            }
        }
        return ans[1 .. ];
    }
}
