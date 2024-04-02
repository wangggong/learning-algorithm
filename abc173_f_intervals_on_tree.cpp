#include <iostream>

using namespace std;
typedef long long ll;
ll n, ans, u, v;

int main() {
    scanf("%lld", &n);
    // \Sigma_{i=1}^n{\Sigma_{j=1}^i{j}}
    // = \Sigma_{i=1}^n{i * (i+1) / 2}
    // = (\Sigma_{i=1}^n{i^2} + \Sigma_{i=1}^n{i}) / 2
    // = (n * (n + 1) * (n + 1/2) / 3 + n * (n + 1) / 2) / 2
    // = n * (n + 1) * (n + 2) / 6
    ans = n * (n + 1) * (n + 2) / 6;
    for (ll i = 1; i < n; i++) {
        scanf("%lld %lld", &u, &v);
        if (u > v)
            swap(u, v);
        ans -= u * (n - v + 1);
    }
    printf("%lld\n", ans);
    return 0;
}
