// 参考: https://codeforces.com/contest/864/submission/228761298
#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll a[N + 10], c[N + 10], n, ans;
bool skip[N + 10];

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i), c[a[i]]++;
    vector<ll> to_add;
    for (ll i = 1; i <= n; i++)
        if (c[i] == 0)
            to_add.push_back(i), ans++;
    for (ll i = 0, j = 0; i < n; i++) {
        if (c[a[i]] <= 1)
            continue;
        if (to_add[j] < a[i])
            c[a[i]]--, a[i] = to_add[j], j++;
        else {
            if (!skip[a[i]])
                skip[a[i]] = true;
            else
                c[a[i]]--, a[i] = to_add[j], j++;
        }
    }
    printf("%lld\n", ans);
    for (ll i = 0; i < n; i++)
        printf("%lld ", a[i]);
    return 0;
}
