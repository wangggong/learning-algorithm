#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;

const ll N = 3000;
ll t, cnt[N + 10], a[N + 10], n;

int main() {
    cin >> t;
    while (t--) {
        cin >> n;
        for (ll i = 0; i < n; i++)
            cin >> a[i];
        ll tot = 0;
        memset(cnt, 0, sizeof cnt);
        cnt[a[0]]++;
        for (ll j = 1; j + 1 < n; j++) {
            ll c = cnt[a[j + 1]];
            for (ll l = j + 2; l < n; l++) {
                if (a[j] == a[l])
                    tot += c;
                c += cnt[a[l]];
            }
            cnt[a[j]]++;
        }
        cout << tot << endl;
    }
    return 0;
}
