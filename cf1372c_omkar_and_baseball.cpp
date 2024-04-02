#include <iostream>

using namespace std;

const int N = 2e5;
int a[N + 10], n, t;

int main() {
    scanf("%d", &t);
    while (t--) {
        scanf("%d", &n);
        for (int i = 1; i <= n; i++)
            scanf("%d", a + i);
        int p = 1, q = n;
        for (; p < q && a[p] == p; p++);
        for (; q > p && a[q] == q; q--);
        if (p == q) {
            printf("%d\n", 0);
            continue;
        }
        bool full = true;
        for (int i = p; i <= q; i++)
            if (a[i] == i)
                full = false;
        printf("%d\n", full ? 1 : 2);
    }
    return 0;
}
