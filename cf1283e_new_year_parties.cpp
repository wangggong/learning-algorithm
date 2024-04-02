#include <iostream>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <algorithm>

using namespace std;

typedef pair<int, int> pii;

int main() {
    int n, x;
    unordered_map<int, int> count;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%d", &x);
        count[x]++;
    }
    vector<pii> vec;
    for (auto c : count)
        vec.push_back(c);
    sort(vec.begin(), vec.end());
    int a = 0, b = 0;
    {
        unordered_set<int> S;
        for (auto v : vec)
            if (S.count(v.first - 1) || S.count(v.first))
                continue;
            else
                S.insert(v.first + 1);
        a = S.size();
    }
    {
        unordered_set<int> S;
        for (auto v : vec) {
            if (v.second >= 3) {
                S.insert(v.first - 1), S.insert(v.first), S.insert(v.first + 1);
                continue;
            }
            bool p = S.count(v.first - 1), c = S.count(v.first);
            if (v.second == 2) {
                if (!p) {
                    S.insert(v.first - 1), S.insert(!c ? v.first : v.first + 1);
                } else {
                    S.insert(v.first), S.insert(v.first + 1);
                }
                continue;
            }
            S.insert(!p ? v.first - 1 : (!c ? v.first : v.first + 1));
        }
        b = S.size();
    }
    printf("%d %d\n", a, b);
    return 0;
}
