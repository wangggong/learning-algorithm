/*
 * @lc app=leetcode.cn id=find-elements-in-a-contaminated-binary-tree lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1261] 在受污染的二叉树中查找元素
 *
 * https://leetcode.cn/problems/find-elements-in-a-contaminated-binary-tree/description/
 *
 * algorithms
 * Medium (73.90%)
 * Total Accepted:    21K
 * Total Submissions: 27.1K
 * Testcase Example:  '["FindElements","find","find"]\n[[[-1,null,-1]],[1],[2]]'
 *
 * 给出一个满足下述规则的二叉树：
 * 
 * 
 * root.val == 0
 * 如果 treeNode.val == x 且 treeNode.left != null，那么 treeNode.left.val == 2 * x +
 * 1
 * 如果 treeNode.val == x 且 treeNode.right != null，那么 treeNode.right.val == 2 * x
 * + 2
 * 
 * 
 * 现在这个二叉树受到「污染」，所有的 treeNode.val 都变成了 -1。
 * 
 * 请你先还原二叉树，然后实现 FindElements 类：
 * 
 * 
 * FindElements(TreeNode* root) 用受污染的二叉树初始化对象，你需要先把它还原。
 * bool find(int target) 判断目标值 target 是否存在于还原后的二叉树中并返回结果。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：
 * ["FindElements","find","find"]
 * [[[-1,null,-1]],[1],[2]]
 * 输出：
 * [null,false,true]
 * 解释：
 * FindElements findElements = new FindElements([-1,null,-1]); 
 * findElements.find(1); // return False 
 * findElements.find(2); // return True 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：
 * ["FindElements","find","find","find"]
 * [[[-1,-1,-1,-1,-1]],[1],[3],[5]]
 * 输出：
 * [null,true,true,false]
 * 解释：
 * FindElements findElements = new FindElements([-1,-1,-1,-1,-1]);
 * findElements.find(1); // return True
 * findElements.find(3); // return True
 * findElements.find(5); // return False
 * 
 * 示例 3：
 * 
 * 
 * 
 * 输入：
 * ["FindElements","find","find","find","find"]
 * [[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]
 * 输出：
 * [null,true,false,false,true]
 * 解释：
 * FindElements findElements = new FindElements([-1,null,-1,-1,null,-1]);
 * findElements.find(2); // return True
 * findElements.find(3); // return False
 * findElements.find(4); // return False
 * findElements.find(5); // return True
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * TreeNode.val == -1
 * 二叉树的高度不超过 20
 * 节点的总数在 [1, 10^4] 之间
 * 调用 find() 的总次数在 [1, 10^4] 之间
 * 0 <= target <= 10^6
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
public class FindElements
{
    private TreeNode tree;

    public FindElements(TreeNode root)
    {
        void dfs(TreeNode node, int val)
        {
            if (node is not null)
            {
                node.val = val;
                dfs(node.left, val * 2 + 1);
                dfs(node.right, val * 2 + 2);
            }
        }
        dfs(root, 0);
        tree = root;
    }

    public bool Find(int target) => Find(tree, target);
    private bool Find(TreeNode node, int val) => node is not null
        && (node.val == val
            || Find(node.left, val)
            || Find(node.right, val));
}

/**
 * Your FindElements object will be instantiated and called as such:
 * FindElements obj = new FindElements(root);
 * bool param_1 = obj.Find(target);
 */
