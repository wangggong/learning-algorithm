#include <iostream>
#include <vector>

using namespace std;

int Q, m, last;
double tot;
vector<int> v;

int main() {
    scanf("%d", &Q);
    while (Q--) {
        int op, k;
        scanf("%d", &op);
        switch (op) {
        case 1:
            scanf("%d", &k);
            v.push_back(k);
            break;
        case 2:
            tot = tot - last + v.back();
            while (m < v.size() - 1 && tot / (m + 1) > (tot + v[m]) / (m + 2))
                tot += v[m], m++;
            last = v.back();
            double ans = v.back() - tot / (m + 1);
            printf("%.10f\n", ans);
            break;
        }
    }
    return 0;
}
