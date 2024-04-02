/*
 * @lc app=leetcode.cn id=delete-nodes-and-return-forest lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1110] 删点成林
 *
 * https://leetcode.cn/problems/delete-nodes-and-return-forest/description/
 *
 * algorithms
 * Medium (64.89%)
 * Total Accepted:    27.8K
 * Total Submissions: 40.8K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n[3,5]'
 *
 * 给出二叉树的根节点 root，树上每个节点都有一个不同的值。
 * 
 * 如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
 * 
 * 返回森林中的每棵树。你可以按任意顺序组织答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
 * 输出：[[1,2,null,4],[6],[7]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [1,2,4,null,3], to_delete = [3]
 * 输出：[[1,2,4]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中的节点数最大为 1000。
 * 每个节点都有一个介于 1 到 1000 之间的值，且各不相同。
 * to_delete.length <= 1000
 * to_delete 包含一些从 1 到 1000、各不相同的值。
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
public class Solution
{
    public IList<TreeNode> DelNodes(TreeNode root, int[] toDelete)
    {
        var ans = new List<TreeNode>();
        var S = toDelete.ToHashSet();
        void traversal(TreeNode node, TreeNode parent, bool addToList)
        {
            if (node is not null)
            {
                var nodeToDelete = S.Contains(node.val);
                if (addToList && !nodeToDelete)
                {
                    ans.Add(node);
                }
                if (nodeToDelete && parent is not null)
                {
                    if (parent.left == node)
                    {
                        parent.left = null;
                    }
                    else
                    {
                        parent.right = null;
                    }
                }
                traversal(node.left, node, nodeToDelete);
                traversal(node.right, node, nodeToDelete);
            }
        }
        traversal(root, null, true);
        return ans;
    }
}
