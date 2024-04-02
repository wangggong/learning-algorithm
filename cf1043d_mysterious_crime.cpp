#include <iostream>
#define __FAST_READ__ std::ios::sync_with_stdio(false), std::cin.tie(nullptr), std::cout.tie(nullptr);

using namespace std;
typedef long long ll;

const int N = 1e5, M = 10;
ll a[M + 10][N + 10], inv[M + 10][N + 10], idx[M + 10], n, m, k, ans;

bool check(int k) {
    for (int i = 0; i < m; i++)
        if (idx[i] + k > n)
            return false;
    for (int i = 1; i < m; i++)
        if (a[i][idx[i] + k] != a[0][idx[0] + k])
            return false;
    return true;
}

int main() {
    __FAST_READ__
    cin >> n >> m;
    for (int i = 0; i < m; i++)
        for (int j = 1; j <= n; j++)
            cin >> a[i][j], inv[i][a[i][j]] = j;
    for (int i = 1; i <= n; i += k) {
        for (int j = 0; j < m; j++)
            idx[j] = inv[j][a[0][i]];
        for (k = 0; check(k); k++);
        ans += k * (k + 1) / 2;
    }
    cout << ans << endl;
    return 0;
}
