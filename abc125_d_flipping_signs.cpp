#include <iostream>

using namespace std;
typedef long long ll;
ll n, v, ans, cnt, _min = 0x3f3f3f3f;

int main() {
    scanf("%lld", &n);
    while (n--) {
        scanf("%lld", &v);
        ans += abs(v);
        cnt += v < 0;
        _min = min(_min, abs(v));
    }
    printf("%lld", ans - (cnt & 1) * _min * 2);
    return 0;
}
