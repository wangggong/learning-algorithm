/*
 * @lc app=leetcode.cn id=collect-coins-in-a-tree lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6356] 收集树中金币
 *
 * https://leetcode.cn/problems/collect-coins-in-a-tree/description/
 *
 * algorithms
 * Hard (26.18%)
 * Total Accepted:    708
 * Total Submissions: 2.3K
 * Testcase Example:  '[1,0,0,0,0,1]\n[[0,1],[1,2],[2,3],[3,4],[4,5]]'
 *
 * 给你一个 n 个节点的无向无根树，节点编号从 0 到 n - 1 。给你整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中
 * edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间有一条边。再给你一个长度为 n 的数组 coins ，其中 coins[i]
 * 可能为 0 也可能为 1 ，1 表示节点 i 处有一个金币。
 * 
 * 一开始，你需要选择树中任意一个节点出发。你可以执行下述操作任意次：
 * 
 * 
 * 收集距离当前节点距离为 2 以内的所有金币，或者
 * 移动到树中一个相邻节点。
 * 
 * 
 * 你需要收集树中所有的金币，并且回到出发节点，请你返回最少经过的边数。
 * 
 * 如果你多次经过一条边，每一次经过都会给答案加一。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：coins = [1,0,0,0,0,1], edges = [[0,1],[1,2],[2,3],[3,4],[4,5]]
 * 输出：2
 * 解释：从节点 2 出发，收集节点 0 处的金币，移动到节点 3 ，收集节点 5 处的金币，然后移动回节点 2 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：coins = [0,0,0,1,1,0,0,1], edges =
 * [[0,1],[0,2],[1,3],[1,4],[2,5],[5,6],[5,7]]
 * 输出：2
 * 解释：从节点 0 出发，收集节点 4 和 3 处的金币，移动到节点 2 处，收集节点 7 处的金币，移动回节点 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == coins.length
 * 1 <= n <= 3 * 10^4
 * 0 <= coins[i] <= 1
 * edges.length == n - 1
 * edges[i].length == 2
 * 0 <= ai, bi < n
 * ai != bi
 * edges 表示一棵合法的树。
 * 
 * 
 */


/*
 * // 拓扑排序.
 * // 
 * // 参考: https://leetcode.cn/problems/collect-coins-in-a-tree/solution/tuo-bu-pai-xu-ji-lu-ru-dui-shi-jian-pyth-6uli/
 * public class Solution
 * {
 *     public int CollectTheCoins(int[] coins, int[][] edges)
 *     {
 *         // 建图
 *         var n = coins.Length;
 *         var G = new List<int>[n];
 *         var degrees = new int[n];
 *         for (var i = 0; i < n; i++)
 *         {
 *             G[i] = new();
 *         }
 *         foreach (var e in edges)
 *         {
 *             var (u, v) = (e[0], e[1]);
 *             G[u].Add(v);
 *             G[v].Add(u);
 *             degrees[u]++;
 *             degrees[v]++;
 *         }
 *         // 剪枝
 *         var Q = new Queue<int>();
 *         for (var i = 0; i < n; i++)
 *         {
 *             if (degrees[i] == 1 && coins[i] == 0)
 *             {
 *                 Q.Enqueue(i);
 *             }
 *         }
 *         var left = n - 1;
 *         while (Q.Count > 0)
 *         {
 *             var u = Q.Dequeue();
 *             left--;
 *             foreach (var v in G[u])
 *             {
 *                 degrees[v]--;
 *                 if (degrees[v] == 1 && coins[v] == 0)
 *                 {
 *                     Q.Enqueue(v);
 *                 }
 *             }
 *         }
 *         // 到二层的拓扑排序
 *         for (var i = 0; i < n; i++)
 *         {
 *             if (degrees[i] == 1 && coins[i] == 1)
 *             {
 *                 Q.Enqueue(i);
 *             }
 *         }
 *         left -= Q.Count;
 *         while (Q.Count > 0)
 *         {
 *             var u = Q.Dequeue();
 *             foreach (var v in G[u])
 *             {
 *                 degrees[v]--;
 *                 if (degrees[v] == 1)
 *                 {
 *                     left--;
 *                 }
 *             }
 *         }
 *         return Math.Max(left * 2, 0);
 *     }
 * }
 */

public class Solution
{
    public int CollectTheCoins(int[] coins, int[][] edges)
    {
        // 建图
        var n = coins.Length;
        var G = new List<int>[n];
        var degrees = new int[n];
        for (var i = 0; i < n; i++)
        {
            G[i] = new();
        }
        foreach (var e in edges)
        {
            var (u, v) = (e[0], e[1]);
            G[u].Add(v);
            G[v].Add(u);
            degrees[u]++;
            degrees[v]++;
        }
        // 剪枝
        var Q = new Queue<int>();
        for (var i = 0; i < n; i++)
        {
            if (degrees[i] == 1 && coins[i] == 0)
            {
                Q.Enqueue(i);
            }
        }
        while (Q.Count > 0)
        {
            var u = Q.Dequeue();
            foreach (var v in G[u])
            {
                degrees[v]--;
                if (degrees[v] == 1 && coins[v] == 0)
                {
                    Q.Enqueue(v);
                }
            }
        }
        // 拓扑排序, 记录入队时间
        var times = new int[n];
        for (var i = 0; i < n; i++)
        {
            if (degrees[i] == 1 && coins[i] == 1)
            {
                Q.Enqueue(i);
            }
        }
        // if (Q.Count <= 1)
        // {
        //     return 0;
        // }
        while (Q.Count > 0)
        {
            var u = Q.Dequeue();
            foreach (var v in G[u])
            {
                degrees[v]--;
                if (degrees[v] == 1)
                {
                    times[v] = times[u] + 1;
                    Q.Enqueue(v);
                }
            }
        }
        return Math.Max(2 * edges.Where(e => times[e[0]] >= 2 && times[e[1]] >= 2).Count(), 0);
    }
}
