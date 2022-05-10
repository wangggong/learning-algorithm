/*
 * Medium
 * 
 * 多项式链表是一种特殊形式的链表，每个节点表示多项式的一项。
 * 
 * 每个节点有三个属性：
 * 
 * coefficient：该项的系数。项 9x4 的系数是 9 。
 * power：该项的指数。项 9x4 的指数是 4 。
 * next：指向下一个节点的指针（引用），如果当前节点为链表的最后一个节点则为 null 。
 * 例如，多项式 5x3 + 4x - 7 可以表示成如下图所示的多项式链表：
 * 
 * 
 * 
 * 多项式链表必须是标准形式的，即多项式必须 严格 按指数 power 的递减顺序排列（即降幂排列）。另外，系数 coefficient 为 0 的项需要省略。
 * 
 * 给定两个多项式链表的头节点 poly1 和 poly2，返回它们的和的头节点。
 * 
 * PolyNode 格式：
 * 
 * 输入/输出格式表示为 n 个节点的列表，其中每个节点表示为 [coefficient, power] 。例如，多项式 5x3 + 4x - 7 表示为： [[5,3],[4,1],[-7,0]] 。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：poly1 = [[1,1]], poly2 = [[1,0]]
 * 输出：[[1,1],[1,0]]
 * 解释：poly1 = x. poly2 = 1. 和为 x + 1.
 * 示例 2：
 * 
 * 输入：poly1 = [[2,2],[4,1],[3,0]], poly2 = [[3,2],[-4,1],[-1,0]]
 * 输出：[[5,2],[2,0]]
 * 解释：poly1 = 2x2 + 4x + 3. poly2 = 3x2 - 4x - 1. 和为 5x2 + 2. 注意，我们省略 "0x" 项。
 * 示例 3：
 * 
 * 输入：poly1 = [[1,2]], poly2 = [[-1,2]]
 * 输出：[]
 * 解释：和为 0。我们返回空链表。
 *  
 * 
 * 提示：
 * 
 * 0 <= n <= 104
 * -109 <= PolyNode.coefficient <= 109
 * PolyNode.coefficient != 0
 * 0 <= PolyNode.power <= 109
 * PolyNode.power > PolyNode.next.power
 * 通过次数1,032
 * 提交次数1,813
 */

/**
 * Definition for polynomial singly-linked list.
 * struct PolyNode {
 *     int coefficient, power;
 *     PolyNode *next;
 *     PolyNode(): coefficient(0), power(0), next(nullptr) {};
 *     PolyNode(int x, int y): coefficient(x), power(y), next(nullptr) {};
 *     PolyNode(int x, int y, PolyNode* next): coefficient(x), power(y), next(next) {};
 * };
 */

class Solution {
    PolyNode* reverse(PolyNode* node) {
        PolyNode *prev = nullptr, *curr = node, *next = nullptr;
        for (; curr != nullptr; prev = curr, curr = next) {
            next = curr->next;
            curr->next = prev;
        }
        return prev;
    }
public:
    PolyNode* addPoly(PolyNode* poly1, PolyNode* poly2) {
        poly1 = reverse(poly1);
        poly2 = reverse(poly2);
        PolyNode* dummy = new PolyNode();
        PolyNode *curr = dummy, *p1 = poly1, *p2 = poly2;
        while (p1 != nullptr && p2 != nullptr) {
            if (p1->power == p2->power) {
                if (p1->coefficient + p2->coefficient)
                    curr->next = new PolyNode(p1->coefficient + p2->coefficient, p1->power), curr = curr->next;
                p1 = p1->next, p2 = p2->next;
            } else if (p1->power < p2->power) {
                curr->next = new PolyNode(p1->coefficient, p1->power);
                curr = curr->next, p1 = p1->next;
            } else {
                curr->next = new PolyNode(p2->coefficient, p2->power);
                curr = curr->next, p2 = p2->next;
            }
        }
        while (p1 != nullptr) {
            curr->next = new PolyNode(p1->coefficient, p1->power);
            curr = curr->next, p1 = p1->next;
        }
        while (p2 != nullptr) {
            curr->next = new PolyNode(p2->coefficient, p2->power);
            curr = curr->next, p2 = p2->next;
        }
        poly1 = reverse(poly1);
        poly2 = reverse(poly2);
        dummy->next = reverse(dummy->next);
        return dummy->next;
    }

    /*
     * // 正着也可行:
     * PolyNode* addPoly(PolyNode* poly1, PolyNode* poly2) {
     *     PolyNode *dummy = new PolyNode();
     *     PolyNode *curr = dummy, *p1 = poly1, *p2 = poly2;
     *     while (p1 != nullptr && p2 != nullptr) {
     *         if (p1->power > p2->power)
     *             curr->next = new PolyNode(p1->coefficient, p1->power), curr = curr->next, p1 = p1->next;
     *         else if (p1->power < p2->power)
     *             curr->next = new PolyNode(p2->coefficient, p2->power), curr = curr->next, p2 = p2->next;
     *         else {
     *             int c = p1->coefficient + p2->coefficient;
     *             if (c)
     *                 curr->next = new PolyNode(c, p1->power), curr = curr->next;
     *             p1 = p1->next, p2 = p2->next;
     *         }
     *     }
     *     if (p1 != nullptr)
     *         curr->next = p1;
     *     else
     *         curr->next = p2;
     *     return dummy->next;
     * }
     */
};
