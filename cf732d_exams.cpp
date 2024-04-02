#include <iostream>
#include <cstring>

using namespace std;

const int N = 1e5;

int n, m;
int d[N + 10], a[N + 10], last[N + 10];

bool check(int k) {
    memset(last, 0, sizeof last);
    for (int i = 0; i <= k; i++)
        if (d[i])
            last[d[i] - 1] = i + 1;
    for (int i = 0; i < m; i++)
        if (!last[i])
            return false;
    int cnt = 0;
    for (int i = 0; i <= k; i++)
        if (d[i] && i + 1 == last[d[i] - 1]) {
            if (cnt < a[d[i] - 1])
                return false;
            cnt -= a[d[i] - 1];
        } else
            cnt++;
    return true;
}

int main() {
    scanf("%d%d", &n, &m);
    for (int i = 0; i < n; i++)
        scanf("%d", &d[i]);
    for (int i = 0; i < m; i++)
        scanf("%d", &a[i]);
    int p = 0, q = n + 1;
    while (p < q) {
        int mid = p + q >> 1;
        if (check(mid))
            q = mid;
        else
            p = mid + 1;
    }
    printf("%d\n", p > n ? -1 : p + 1);
    return 0;
}
