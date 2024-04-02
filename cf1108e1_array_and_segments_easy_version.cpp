#include <iostream>
#include <vector>
#include <cstring>

using namespace std;
typedef long long ll;

const ll N = 300;
ll a[N + 10], b[N + 10], n, m, I[N + 10][2], maxVal, ans[N + 10], cnt, cur[N + 10], curCnt;

int main() {
    cin >> n >> m;
    for (ll i = 1; i <= n; i++)
        cin >> a[i];
    for (ll i = 1; i <= m; i++)
        cin >> I[i][0] >> I[i][1];
    for (ll i = 1; i <= n; i++) {
        curCnt = 0, memcpy(b, a, sizeof b);
        for (ll j = 1; j <= m; j++) {
            if (I[j][0] <= i && i <= I[j][1])
                continue;
            cur[curCnt++] = j;
            for (ll k = I[j][0]; k <= I[j][1]; b[k]--, k++);
        }
        ll curMin = b[1];
        for (ll j = 1; j <= n; j++)
            curMin = min(curMin, b[j]);
        if (maxVal < a[i] - curMin)
            maxVal = a[i] - curMin, memcpy(ans, cur, sizeof ans), cnt = curCnt;
    }
    cout << maxVal << endl << cnt << endl;
    for (ll i = 0; i < cnt; i++)
        cout << ans[i] << " ";
    return 0;
}
