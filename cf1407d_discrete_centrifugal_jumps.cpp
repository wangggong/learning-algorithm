#include <iostream>
#include <stack>
#define DEBUG false
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;

const ll N = 3e5;
ll n, a[N + 10], dp[N + 10];

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 1; i <= n; i++)
        cin >> a[i], dp[i] = n - i;
    stack<ll> S, T;
    for (ll i = n; i >= 1; i--) {
        for (; !S.empty() && a[S.top()] < a[i]; S.pop())
            dp[i] = min(dp[i], dp[S.top()] + 1);
        if (!S.empty()) {
            dp[i] = min(dp[i], dp[S.top()] + 1);
            if (a[S.top()] == a[i])
                S.pop();
        }
        for (; !T.empty() && a[T.top()] > a[i]; T.pop())
            dp[i] = min(dp[i], dp[T.top()] + 1);
        if (!T.empty()) {
            dp[i] = min(dp[i], dp[T.top()] + 1);
            if (a[T.top()] == a[i])
                T.pop();
        }
        S.push(i);
        T.push(i);
    }
    if (DEBUG) {
        for (ll i = 1; i <= n; i++)
            cout << dp[i] << ' '; cout << endl;
    }
    cout << dp[1] << endl;
    return 0;
}

