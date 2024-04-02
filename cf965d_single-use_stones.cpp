#include <iostream>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;

const ll N = 1e5;
ll a[N + 10], ps[N + 10], w, k, ans = 0x3f3f3f3f;

int main() {
    __FAST__;
    cin >> w >> k;
    for (ll i = 0; i < w - 1; i++)
        cin >> a[i];
    for (ll i = 1; i < w; i++)
        ps[i] = ps[i - 1] + a[i - 1];
    for (ll i = 0; i + k < w; i++)
        ans = min(ans, ps[i + k] - ps[i]);
    cout << ans << endl;
    return 0;
}
