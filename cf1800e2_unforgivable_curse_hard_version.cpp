#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 2e5;
const ll A = 26;
ll T, n, k, c[A + 10];
char s[N + 10], t[N + 10];

/*
 * bool check() {
 *     memset(c, 0, sizeof c);
 *     if (n <= k) {
 *         for (ll i = 0; i < n; i++)
 *             if (s[i] != t[i])
 *                 return false;
 *         return true;
 *     }
 *     if (n - 1 - k >= k) {
 *         for (ll i = 0; i < n; i++)
 *             c[s[i] - 'a']++, c[t[i] - 'a']--;
 *         for (ll i = 0; i < A; i++)
 *             if (c[i])
 *                 return false;
 *         return true;
 *     }
 *     for (ll i = 0; i + k < n; i++)
 *         c[s[i] - 'a']++, c[t[i] - 'a']--;
 *     for (ll i = n - k; i < k; i++)
 *         if (s[i] != t[i])
 *             return false;
 *     for (ll i = k; i < n; i++)
 *         c[s[i] - 'a']++, c[t[i] - 'a']--;
 *         for (ll i = 0; i < A; i++)
 *             if (c[i])
 *                 return false;
 *     return true;
 * }
 */

// 参考: https://codeforces.com/contest/1800/submission/239865964
bool check() {
    memset(c, 0, sizeof c);
    for (ll i = 0; i < n; i++) {
        if (i >= n - k && i < k && s[i] != t[i])
            return false;
        c[s[i] - 'a']++, c[t[i] - 'a']--;
    }
    for (ll i = 0; i < A; i++)
        if (c[i])
            return false;
    return true;
}

int main() {
    for (scanf("%lld", &T); T; T--) {
        scanf("%lld%lld%s%s", &n, &k, s, t);
        puts(check() ? "YES" : "NO");
    }
    return 0;
}
