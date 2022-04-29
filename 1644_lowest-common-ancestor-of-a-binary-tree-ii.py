# Hard
# 
# 给定一棵二叉树的根节点 root，返回给定节点 p 和 q 的最近公共祖先（LCA）节点。如果 p 或 q 之一 不存在 于该二叉树中，返回 null。树中的每个节点值都是互不相同的。
# 
# 根据维基百科中对最近公共祖先节点的定义：“两个节点 p 和 q 在二叉树 T 中的最近公共祖先节点是 后代节点 中既包括 p 又包括 q 的最深节点（我们允许 一个节点为自身的一个后代节点 ）”。一个节点 x 的 后代节点 是节点 x 到某一叶节点间的路径中的节点 y。
# 
#  
# 
# 示例 1:
# 
# 
# 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
# 输出： 3
# 解释： 节点 5 和 1 的共同祖先节点是 3。
# 示例 2:
# 
# 
# 
# 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
# 输出： 5
# 解释： 节点 5 和 4 的共同祖先节点是 5。根据共同祖先节点的定义，一个节点可以是自身的后代节点。
# 示例 3:
# 
# 
# 
# 输入： root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 10
# 输出： null
# 解释： 节点 10 不存在于树中，所以返回 null。
#  
# 
# 提示:
# 
# 树中节点个数的范围是 [1, 104]
# -109 <= Node.val <= 109
# 所有节点的值 Node.val 互不相同
# p != q
#  
# 
# 进阶： 在不检查节点是否存在的情况下，你可以遍历树找出最近公共祖先节点吗？
# 
# 通过次数2,090
# 提交次数3,844

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def getPath(self, root: 'TreeNode', node: 'TreeNode') -> List:
        if root is None:
            return []
        if root == node:
            return [node]
        left, right = self.getPath(root.left, node), self.getPath(root.right, node)
        if left or right:
            return left + [root] if left else right + [root]
        return []

    '''
    # 现场做的, 就直接把路径打出来了. 很玩赖.
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if root is None or p is None or q is None:
            return None
        p_path, q_path = self.getPath(root, p), self.getPath(root, q)
        if not p_path or not q_path:
            return None
        n, m = len(p_path), len(q_path)
        if n < m:
            q_path = q_path[m-n:]
        elif n > m:
            p_path = p_path[n-m:]
        for node_p, node_q in zip(p_path, q_path):
            if node_p == node_q:
                return node_p
        return None
    '''

    def __init__(self):
        self.ans = None

    '''树形 DP, 返回值代表当前根节点存在 `p` 或者 `q`.'''
    def traversal(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> bool:
        if root is None:
            return False
        l = self.traversal(root.left, p, q)
        r = self.traversal(root.right, p, q)
        # case 1: 如果一个在左子树一个在右子树, 那当前根节点就是 LCA.
        if l and r:
            self.ans = root
        # case 2: 如果当前根节点为其中一个节点:
        #
        # 不妨设当前根节点为 `p`, 同时 `q` 在另一个子树中 (左子树右子树无所谓), 当前根节点也是 LCA.
        if (l or r) and (root == p or root == q):
            self.ans = root
        # 返回当前根节点下的子树是否存在 `p` 或者 `q`.
        # case 1: 左子树 / 右子树中存在 `p` 或者 `q`;
        # case 2: 当前节点本身是 `p` 或者 `q`;
        return l or r or root == p or root == q

    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        self.traversal(root, p, q)
        return self.ans

