#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
ll n, k, x;
vector<ll> a, miss;

int main() {
    scanf("%lld %lld", &n, &k);
    for (ll i = 0; i < k; i++)
        scanf("%lld", &x), a.push_back(x);
    for (ll i = 1, j = 0; i <= n; i++)
        if (j < k && i == a[j])
            j++;
        else
            miss.push_back(i);
    ll i = 0;
    for (ll j = 0; j + 1 < k; j++) {
        printf("%lld ", a[j]);
        if (i < n - k && miss[i] < a[j])
            printf("%lld ", miss[i++]);
    }
    ll j = n - k - 1;
    for (; j >= i && miss[j] > a[k - 1]; j--)
        printf("%lld ", miss[j]);
    printf("%lld ", a[k - 1]);
    for (; j >= i; j--)
        printf("%lld ", miss[j]);
    return 0;
}
