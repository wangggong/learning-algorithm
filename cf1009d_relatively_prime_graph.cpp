#include <iostream>

using namespace std;

const int N = 1e5;
int edges[N + 10][2];

int gcd(int x, int y) { return y == 0 ? x : gcd(y, x % y); }

int main() {
    int n, m, cnt = 0;
    cin >> n >> m;
    if (m < n - 1) {
        cout << "Impossible" << endl;
        return 0;
    }
    for (int p = 1, q = 2; p <= n; ) {
        for (; q <= n && gcd(p, q) != 1; q++);
        if (q > n) {
            p++, q = p + 1;
            continue;
        }
        edges[cnt][0] = p, edges[cnt][1] = q, q++, cnt++;
        if (cnt == m)
            break;
    }
    if (cnt < m)
        cout << "Impossible" << endl;
    else {
        cout << "Possible" << endl;
        for (int i = 0; i < m; i++)
            printf("%d %d\n", edges[i][0], edges[i][1]);
    }
    return 0;
}
