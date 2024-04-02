#include <iostream>
#include <unordered_set>

using namespace std;
typedef long long ll;

// 看了答案, 被爆了一次 long long, 被卡了一次 unordered_set. 真有它的 (
const ll N = 1e5;
ll a[N + 10], n, tot, cur;

int main() {
    cin >> n;
    for (ll i = 0; i < n; i++)
        cin >> a[i], tot += a[i];
    if (n == 1 || tot % 2 != 0) {
        cout << "NO" << endl;
        return 0;
    }
    unordered_set<ll> S(N + 10);
    S.clear(); S.insert(0); cur = 0;
    for (ll i = 0; i < n; i++) {
        cur += a[i];
        if (S.count(cur - tot / 2)) {
            cout << "YES" << endl;
            return 0;
        }
        S.insert(a[i]);
    }
    S.clear(); S.insert(0); cur = 0;
    for (ll i = n - 1; i >= 0; i--) {
        cur += a[i];
        if (S.count(cur - tot / 2)) {
            cout << "YES" << endl;
            return 0;
        }
        S.insert(a[i]);
    }
    cout << "NO" << endl;
    return 0;
}
