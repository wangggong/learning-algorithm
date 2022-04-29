/*
 * Medium
 * 
 * 将一个 二叉搜索树 就地转化为一个 已排序的双向循环链表 。
 * 
 * 对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。
 * 
 * 特别地，我们希望可以 就地 完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中最小元素的指针。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：root = [4,2,5,1,3] 
 * 
 * 
 * 输出：[1,2,3,4,5]
 * 
 * 解释：下图显示了转化后的二叉搜索树，实线表示后继关系，虚线表示前驱关系。
 * 
 * 示例 2：
 * 
 * 输入：root = [2,1,3]
 * 输出：[1,2,3]
 * 示例 3：
 * 
 * 输入：root = []
 * 输出：[]
 * 解释：输入是空树，所以输出也是空链表。
 * 示例 4：
 * 
 * 输入：root = [1]
 * 输出：[1]
 *  
 * 
 * 提示：
 * 
 * -1000 <= Node.val <= 1000
 * Node.left.val < Node.val < Node.right.val
 * Node.val 的所有值都是独一无二的
 * 0 <= Number of Nodes <= 2000
 * 通过次数10,115
 * 提交次数14,900
 */

/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/

typedef pair<Node*, Node*> PNN;

class Solution {
    PNN _treeToDoublyList(Node* root) {
        if (root == nullptr) return {nullptr, nullptr};
        auto left = root->left, right = root->right;
        root->left = nullptr, root->right = nullptr;
        auto leftAns = _treeToDoublyList(left), rightAns = _treeToDoublyList(right);
        auto leftHead = leftAns.first, leftTail = leftAns.second;
        auto rightHead = rightAns.first, rightTail = rightAns.second;
        if (leftTail != nullptr) root->left = leftTail, leftTail->right = root;
        if (rightHead != nullptr) root->right = rightHead, rightHead->left = root;
        return {leftHead == nullptr ? root : leftHead, rightTail == nullptr ? root : rightTail};
    }

    Node *head, *tail;

    void helper(Node* root) {
        if (root == nullptr) return;
        auto left = root->left;
        root->left = nullptr;
        helper(left);
        root->left = tail;
        if (tail != nullptr) tail->right = root;
        else head = root;
        tail = root;
        helper(root->right);
        return;
    }

public:
    /*
     * Node* treeToDoublyList(Node* root) {
     *     if (root == nullptr) return root;
     *     auto ans = _treeToDoublyList(root);
     *     auto head = ans.first, tail = ans.second;
     *     head->left = tail, tail->right = head;
     *     return head;
     * }
     */

    // 官解给的是个中序遍历, 基本都是利用全局变量为所欲为那种...
    Node* treeToDoublyList(Node* root) {
        if (root == nullptr) return root;
        helper(root);
        head->left = tail, tail->right = head;
        return head;
    }
};
