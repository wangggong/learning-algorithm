#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 50, MOD = 998244353;
ll a[N + 10][N + 10], n, K, ans = 1;
ll pa[N + 10], cnt[N + 10];
ll F[N + 10];

void init_uf() {
    for (ll i = 0; i <= N; i++)
        pa[i] = i, cnt[i] = 0;
    return;
}

ll query(ll k) {
    if (pa[k] != k)
        pa[k] = query(pa[k]);
    return pa[k];
}

void merge(ll p, ll q) { pa[query(p)] = query(q); }

void init() {
    F[0] = 1;
    for (ll i = 1; i <= N; i++)
        F[i] = (F[i - 1] * i) % MOD;
    return;
}

void handle() {
    init_uf();
    for (ll i = 0; i < n; i++)
        for (ll j = 0; j < n; j++)
            if (query(i) != query(j)) {
                auto valid = true;
                for (ll k = 0; k < n; k++)
                    if (a[i][k] + a[j][k] > K) {
                        valid = false;
                        break;
                    }
                if (valid)
                    merge(i, j);
            }
    for (ll i = 0; i < n; i++)
        cnt[query(i)]++;
    for (ll i = 0; i < n; i++)
        if (cnt[i] != 0)
            ans = (ans * F[cnt[i]]) % MOD;
    return;
}

int main() {
    init();
    scanf("%lld %lld", &n, &K);
    for (ll i = 0; i < n; i++)
        for (ll j = 0; j < n; j++)
            scanf("%lld", &a[i][j]);
    handle();
    for (ll i = 0; i < n; i++)
        for (ll j = i + 1; j < n; j++)
            swap(a[i][j], a[j][i]);
    handle();
    printf("%lld\n", ans);
    return 0;
}
