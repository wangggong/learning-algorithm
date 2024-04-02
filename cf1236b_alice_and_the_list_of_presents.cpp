#include <iostream>

using namespace std;
typedef long long ll;

const ll MOD = 1e9 + 7;
ll n, m;

ll _pow(ll n, ll k) {
    ll ans = 1;
    for (; k; k >>= 1)
        ans = (k & 1) ? (n * ans % MOD) : ans, n = n * n % MOD;
    return ans;
}

int main() {
    cin >> n >> m;
    cout << _pow(_pow(2, m) - 1, n) << endl;
    return 0;
}
