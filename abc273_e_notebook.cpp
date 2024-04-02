#include <iostream>
#include <unordered_map>

using namespace std;

typedef long long ll;

struct ListNode {
    ll val;
    ListNode *prev, *next;

    ListNode(ll val) {
        this->val = val;
    }
};

ll q, a;
char op[8];

int main() {
    unordered_map<ll, ListNode*> m;
    auto dummy = new ListNode(-1);
    auto tail = dummy;
    auto node = dummy;
    scanf("%lld", &q);
    while (q--) {
        scanf("%s %lld", op, &a);
        switch (op[0]) {
        case 'A':
            node = new ListNode(a);
            node->prev = tail, tail->next = node, tail = tail->next;
            break;
        case 'D':
            if (tail != dummy)
                tail = tail->prev;
            break;
        case 'S':
            m[a] = tail;
            break;
        case 'L':
            tail = m.count(a) ? m[a] : dummy;
            break;
        default:
            break;
        }
        printf("%lld ", tail == dummy ? -1 : tail->val);
    }
    return 0;
}
