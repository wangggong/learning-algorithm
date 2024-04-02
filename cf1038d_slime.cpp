#include <iostream>

using namespace std;

typedef long long ll;

const int N = 5e5;
int a[N + 10];

int main() {
    int n;
    scanf("%d", &n);
    for (int i = 0; i < n; i++)
        scanf("%d", &a[i]);
    if (n == 1) {
        printf("%d\n", a[0]);
        return 0;
    }
    ll tot = 0;
    int minAbs = 0x3f3f3f3f, minVal = a[0], maxVal = a[0];
    for (int i = 0; i < n; i++)
        tot += abs(a[i]), minAbs = min(minAbs, abs(a[i])), minVal = min(minVal, a[i]), maxVal = max(maxVal, a[i]);
    printf("%lld\n", minVal >= 0 || maxVal < 0 ? tot - 2 * minAbs : tot);
    return 0;
}
