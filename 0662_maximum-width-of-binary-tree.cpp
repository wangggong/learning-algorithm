/*
 * @lc app=leetcode.cn id=maximum-width-of-binary-tree lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [662] 二叉树最大宽度
 *
 * https://leetcode-cn.com/problems/maximum-width-of-binary-tree/description/
 *
 * algorithms
 * Medium (40.87%)
 * Total Accepted:    42.3K
 * Total Submissions: 103.2K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary
 * tree）结构相同，但一些节点为空。
 * 
 * 每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
 * 
 * 示例 1:
 * 
 * 
 * 输入: 
 * 
 * ⁠          1
 * ⁠        /   \
 * ⁠       3     2
 * ⁠      / \     \  
 * ⁠     5   3     9 
 * 
 * 输出: 4
 * 解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: 
 * 
 * ⁠         1
 * ⁠        /  
 * ⁠       3    
 * ⁠      / \       
 * ⁠     5   3     
 * 
 * 输出: 2
 * 解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: 
 * 
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2 
 * ⁠      /        
 * ⁠     5      
 * 
 * 输出: 2
 * 解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。
 * 
 * 
 * 示例 4:
 * 
 * 
 * 输入: 
 * 
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      /     \  
 * ⁠     5       9 
 * ⁠    /         \
 * ⁠   6           7
 * 输出: 8
 * 解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。
 * 
 * 
 * 注意: 答案在32位有符号整数的表示范围内。
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

// 数据范围爆炸, Go 的也存在这个问题, 开了 `int64` 都不行...
typedef pair<TreeNode*, unsigned long long> PTL;

class Solution {
public:
    int widthOfBinaryTree(TreeNode* root) {
        if (root == nullptr) return 0;
        unsigned long long ans = 0;
        queue<PTL> Q;
        Q.push({root, 0});
        while (!Q.empty()) {
            unsigned long long minVal = -1, maxVal = 0;
            for (int size = Q.size(); size > 0; size--) {
                auto q = Q.front(); Q.pop();
                auto node = q.first;
                auto val = q.second;
                minVal = min(minVal, val);
                maxVal = max(maxVal, val);
                if (node->left != nullptr) Q.push({node->left, val << 1});
                if (node->right != nullptr) Q.push({node->right, val << 1 | 1});
            }
            ans = max(ans, maxVal - minVal + 1);
        }
        return ans;
    }
};
