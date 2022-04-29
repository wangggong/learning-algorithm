#
# @lc app=leetcode.cn id=diameter-of-binary-tree lang=python3
#
# NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
#
# [543] 二叉树的直径
#
# https://leetcode-cn.com/problems/diameter-of-binary-tree/description/
#
# algorithms
# Easy (56.59%)
# Total Accepted:    202.3K
# Total Submissions: 357.1K
# Testcase Example:  '[1,2,3,4,5]'
#
# 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
# 
# 
# 
# 示例 :
# 给定二叉树
# 
# ⁠         1
# ⁠        / \
# ⁠       2   3
# ⁠      / \     
# ⁠     4   5    
# 
# 
# 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
# 
# 
# 
# 注意：两结点之间的路径长度是以它们之间边的数目表示。
# 
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def __init__(self):
        self.ans = 0

    def traversal(self, root: TreeNode) -> int:
        if root is None:
            return 0
        left = self.traversal(root.left)
        right = self.traversal(root.right)
        self.ans = max(self.ans, left + right)
        return max(left, right) + 1

    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        self.traversal(root)
        return self.ans

