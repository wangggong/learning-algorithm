#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 3e5;
ll n, m, A[N + 10][2];

bool check(ll x, ll y) {
    for (ll i = 0; i < m; i++) {
        if (A[i][0] == x || A[i][1] == x || A[i][0] == y || A[i][1] == y)
            continue;
        return y == 0 ? (check(x, A[i][0]) || check(x, A[i][1])) : false;
    }
    return true;
}

int main() {
    scanf("%lld %lld", &n, &m);
    for (ll i = 0; i < m; i++)
        scanf("%lld %lld", A[i], A[i] + 1);
    puts((check(A[0][0], 0) || check(A[0][1], 0)) ? "YES" : "NO");
    return 0;
}
