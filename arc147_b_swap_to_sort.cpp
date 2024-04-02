#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
typedef pair<char, ll> pcl;
const ll N = 400;
ll a[N + 10], n;
vector<pcl> ops;

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", a + i);
    for (ll p = 1, q = 1; p <= n; p = q + 1) {
        for (; p <= n && (a[p] - p) % 2 == 0; p++);
        if (p > n)
            break;
        for (q = p + 1; q <= n && (a[q] - q) % 2 == 0; q += 2);
        for (; q > p + 1; q -= 2) {
            swap(a[q], a[q - 2]);
            ops.push_back({'B', q - 2});
        }
        swap(a[p], a[p + 1]);
        ops.push_back({'A', p});
    }
    for (ll i = 1; i <= n; i++)
        if (a[i] != i) {
            ll j = i + 2;
            for (; j <= n && a[j] != i; j += 2);
            for (; j > i; j -= 2) {
                swap(a[j], a[j - 2]);
                ops.push_back({'B', j - 2});
            }
        }
    printf("%lu\n", ops.size());
    for (auto op : ops)
        printf("%c %lld\n", op.first, op.second);
    return 0;
}
