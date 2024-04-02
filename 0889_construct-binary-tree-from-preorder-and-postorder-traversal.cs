/*
 * @lc app=leetcode.cn id=construct-binary-tree-from-preorder-and-postorder-traversal lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [889] 根据前序和后序遍历构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (67.85%)
 * Total Accepted:    44.4K
 * Total Submissions: 65.4K
 * Testcase Example:  '[1,2,4,5,3,6,7]\n[4,5,2,6,7,3,1]'
 *
 * 给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder
 * 是同一棵树的后序遍历，重构并返回二叉树。
 * 
 * 如果存在多个答案，您可以返回其中 任何 一个。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
 * 输出：[1,2,3,4,5,6,7]
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: preorder = [1], postorder = [1]
 * 输出: [1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= preorder.length <= 30
 * 1 <= preorder[i] <= preorder.length
 * preorder 中所有值都 不同
 * postorder.length == preorder.length
 * 1 <= postorder[i] <= postorder.length
 * postorder 中所有值都 不同
 * 保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历
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
    public TreeNode ConstructFromPrePost(int[] preorder, int[] postorder)
        => ConstructFromPrePost(preorder, postorder, 0, preorder.Length, 0, postorder.Length);

    public TreeNode ConstructFromPrePost(int[] preorder, int[] postorder, int p, int q, int s, int t)
    {
        var n = q - p;
        if (n is 0) { return null; }
        var root = new TreeNode(preorder[p]);
        if (n is 1) { return root; }
        var (vLeft, vRight) = (preorder[p + 1], postorder[t - 2]);
        if (vLeft == vRight)
        {
            root.left = ConstructFromPrePost(preorder, postorder, p + 1, q, s, t - 1);
            return root;
        }
        var l = Array.IndexOf(postorder, vLeft);
        var r = Array.IndexOf(preorder, vRight);
        root.left = ConstructFromPrePost(preorder, postorder, p + 1, r, s, l + 1);
        root.right = ConstructFromPrePost(preorder, postorder, r, q, l + 1, t - 1);
        return root;
    }
}
