/*
 * @lc app=leetcode.cn id=kth-ancestor-of-a-tree-node lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1483] 树节点的第 K 个祖先
 *
 * https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/description/
 *
 * algorithms
 * Hard (34.24%)
 * Total Accepted:    12.5K
 * Total Submissions: 29.1K
 * Testcase Example:  '["TreeAncestor","getKthAncestor","getKthAncestor","getKthAncestor"]\n' +
  '[[7,[-1,0,0,1,1,2,2]],[3,1],[5,2],[6,3]]'
 *
 * 给你一棵树，树上有 n 个节点，按从 0 到 n-1 编号。树以父节点数组的形式给出，其中 parent[i] 是节点 i 的父节点。树的根节点是编号为
 * 0 的节点。
 * 
 * 树节点的第 k 个祖先节点是从该节点到根节点路径上的第 k 个节点。
 * 
 * 实现 TreeAncestor 类：
 * 
 * 
 * TreeAncestor（int n， int[] parent） 对树和父数组中的节点数初始化对象。
 * getKthAncestor(int node, int k) 返回节点 node 的第 k 个祖先节点。如果不存在这样的祖先节点，返回 -1
 * 。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：
 * ["TreeAncestor","getKthAncestor","getKthAncestor","getKthAncestor"]
 * [[7,[-1,0,0,1,1,2,2]],[3,1],[5,2],[6,3]]
 * 
 * 输出：
 * [null,1,0,-1]
 * 
 * 解释：
 * TreeAncestor treeAncestor = new TreeAncestor(7, [-1, 0, 0, 1, 1, 2, 2]);
 * 
 * treeAncestor.getKthAncestor(3, 1);  // 返回 1 ，它是 3 的父节点
 * treeAncestor.getKthAncestor(5, 2);  // 返回 0 ，它是 5 的祖父节点
 * treeAncestor.getKthAncestor(6, 3);  // 返回 -1 因为不存在满足要求的祖先节点
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= k <= n <= 5 * 10^4
 * parent[0] == -1 表示编号为 0 的节点是根节点。
 * 对于所有的 0 < i < n ，0 <= parent[i] < n 总成立
 * 0 <= node < n
 * 至多查询 5 * 10^4 次
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/
//
// 今天学习倍增
public class TreeAncestor
{
    private const int N = 16;
    private int[][] Tree;

    public TreeAncestor(int n, int[] parent)
    {
        Tree = Enumerable
            .Range(0, N)
            .Select(_ => Enumerable
                .Range(0, n)
                .Select(_ => -1)
                .ToArray())
            .ToArray();
        for (var i = 0; i < n; i++)
        {
            Tree[0][i] = parent[i];
        }
        for (var i = 1; i < N; i++)
        {
            for (var j = 0; j < n; j++)
            {
                if (Tree[i - 1][j] is not -1)
                {
                    Tree[i][j] = Tree[i - 1][Tree[i - 1][j]];
                }
            }
        }
    }

    public int GetKthAncestor(int node, int k)
    {
        for (var i = 0; i < N; i++)
        {
            if (((k >> i) & 1) is not 0)
            {
                node = Tree[i][node];
            }
            if (node is -1)
            {
                break;
            }
        }
        return node;
    }
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * TreeAncestor obj = new TreeAncestor(n, parent);
 * int param_1 = obj.GetKthAncestor(node,k);
 */
