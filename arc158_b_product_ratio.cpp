#include <iostream>
#include <unordered_set>
#include <vector>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll n, k, a[N + 10];
double mn = 3, mx = -3;

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    sort(a, a + n);
    unordered_set<ll> S;
    for (k = 0; k < n && a[k] <= 0; k++);
    if (k > 0) {
        for (ll i = 0; i < 3 && i < n; i++)
            S.insert(i);
        for (ll i = 0; i < 3 && k - 1 - i >= 0; i++)
            S.insert(k - 1 - i);
    }
    if (k < n) {
        for (ll i = 0; i < 3 && k + i < n; i++)
            S.insert(k + i);
        for (ll i = 0; i < 3 && n - 1 - i >= 0; i++)
            S.insert(n - 1 - i);
    }
    vector<double> vec;
    for (auto it = S.begin(); it != S.end(); ++it)
        vec.push_back((double)a[*it]);
    for (ll i = 0, m = vec.size(); i < m; i++)
        for (ll j = i + 1; j < m; j++)
            for (ll k = j + 1; k < m; k++) {
                double v = (vec[i] + vec[j] + vec[k]) / (vec[i] * vec[j] * vec[k]);
                mn = min(mn, v);
                mx = max(mx, v);
            }
    printf("%.15f\n%.15f\n", mn, mx);
    return 0;
}
