#include <iostream>
#include <algorithm>
#include <set>

using namespace std;

const int N = 1e5;
int n, m, k, a[N + 10], b[N + 10];

int get(int arr[], int n, int k) {
    int p = -1, q = n;
    while (p < q) {
        int mid = (p + q + 1) >> 1;
        if (arr[mid] >= k)
            p = mid;
        else
            q = mid - 1;
    }
    return p;
}

int main() {
    cin >> n >> m >> k;
    for (int i = 0; i < n; i++)
        cin >> a[i];
    for (int i = 0; i < m; i++)
        cin >> b[i];
    if (n > m) {
        cout << "YES" << endl;
        return 0;
    }
    sort(a, a + n, [](int &p, int &q) { return p > q; });
    sort(b, b + m, [](int &p, int &q) { return p > q; });
    set<int> S;
    for (int i = 0; i < n; i++)
        S.insert(a[i]);
    for (int v : S)
        if (get(a, n, v) > get(b, m, v)) {
            cout << "YES" << endl;
            return 0;
        }
    cout << "NO" << endl;
    return 0;
}
