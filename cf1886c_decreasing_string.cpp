// 参考: https://codeforces.com/contest/1886/submission/235784116
#include <iostream>
#include <cstring>
#include <cmath>

using namespace std;
typedef long long ll;
const ll N = 1e6;
ll t, p;
char s[N + 10], S[N + 10];

char f() {
    ll n = strlen(s);
    if (p < n)
        return s[p];
    double m = n * 2 + 1;
    ll k = (m - sqrt(m * m - 8 * p)) / 2;
    p -= (n * 2 - k + 1) * k / 2;
    for (ll i = 0, idx = 0; ; S[idx++] = s[i++])
        while (idx && s[i] < S[idx - 1]) {
            k--, idx--;
            if (!k)
                return p < idx ? S[p] : s[i + p - idx];
        }
}

int main() {
    for (scanf("%lld", &t); t; t--)
        scanf("%s%lld", s, &p), p--, printf("%c", f());
    return 0;
}
