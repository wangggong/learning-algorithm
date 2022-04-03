/*
 * @lc app=leetcode.cn id=recover-binary-search-tree lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [99] 恢复二叉搜索树
 *
 * https://leetcode-cn.com/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Medium (60.80%)
 * Total Accepted:    83.6K
 * Total Submissions: 137.6K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * 给你二叉搜索树的根节点 root ，该树中的两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [1,3,null,null,2]
 * 输出：[3,1,null,null,2]
 * 解释：3 不能是 1 左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [3,1,4,null,null,2]
 * 输出：[2,1,4,null,null,3]
 * 解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树上节点的数目在范围 [2, 1000] 内
 * -2^31 <= Node.val <= 2^31 - 1
 * 
 * 
 * 
 * 
 * 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？
 * 
 */
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    void inorder(TreeNode* r, vector<int>& N) {
        if (r == nullptr) return;
        inorder(r->left, N);
        N.push_back(r->val);
        inorder(r->right, N);
    }

    pair<int, int> findUnsorted(vector<int> N) {
        auto n = N.size();
        auto p = -1;
        auto q = -1;
        for (auto i = 0; i + 1 < n; i++) {
            if (N[i + 1] < N[i]) {
                q = i + 1;
                if (p == -1) {
                    p = i;
                } else {
                    break;
                }
            }
        }
        return make_pair(N[p], N[q]);
    }

    void recover(TreeNode* r, int n, int p, int q) {
        if (r == nullptr) return;

        if (r->val == p || r->val == q) {
            r->val = r->val == p ? q : p;
            n--;
            if (n == 0) return;

        }

        recover(r->left, n, p, q);
        recover(r->right, n, p, q);
    }

    void recoverTree(TreeNode* r) {
        vector<int> N;
        inorder(r, N);
        auto p = findUnsorted(N);
        recover(r, 2, p.first, p.second);
    }
};
