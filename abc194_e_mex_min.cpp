#include <iostream>
#include <unordered_map>

using namespace std;
typedef long long ll;
const ll N = 1500000;
ll n, m, a[N + 10], mex;
unordered_map<ll, ll> M;

int main() {
    scanf("%lld %lld", &n, &m);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll i = 0; i < m; i++)
        M[a[i]]++;
    for (mex = 0; M[mex]; mex++);
    for (ll i = m; i < n; i++) {
        M[a[i]]++, M[a[i - m]]--;
        if (!M[a[i - m]])
            mex = min(mex, a[i - m]);
    }
    printf("%lld", mex);
    return 0;
}
