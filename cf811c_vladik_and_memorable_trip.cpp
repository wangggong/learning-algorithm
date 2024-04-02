#include <iostream>
#include <cstring>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;

const ll N = 5000;
ll a[N + 10], n, dp[N + 10], fr[N + 10], la[N + 10];
bool vis[N + 10];

int main() {
    __FAST__;
    memset(fr, -1, sizeof fr);
    memset(la, -1, sizeof la);
    cin >> n;
    for (ll i = 0; i < n; i++)
        cin >> a[i], fr[a[i]] = fr[a[i]] < 0 ? i : fr[a[i]], la[a[i]] = i;
    for (ll i = 0; i < n; i++) {
        memset(vis, false, sizeof vis);
        ll cur = 0;
        dp[i + 1] = dp[i];
        for (ll j = i, le = fr[a[i]]; j >= 0; j--) {
            if (!vis[a[j]]) {
                if (la[a[j]] > i)
                    break;
                vis[a[j]] = true, le = min(le, fr[a[j]]), cur ^= a[j];
            }
            if (le == j)
                dp[i + 1] = max(dp[i + 1], dp[le] + cur);
        }
    }
    cout << dp[n] << endl;
    return 0;
}
