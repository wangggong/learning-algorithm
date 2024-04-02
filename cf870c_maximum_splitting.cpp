// 参考: https://codeforces.com/problemset/submission/870/250444080
//
// 想破头也没想到分 `4`, 太菜了...
#include <iostream>

using namespace std;
typedef long long ll;
ll t, n;

ll get(ll n) {
    return (n == 1 || n == 2 || n == 3 || n == 5 || n == 7 || n == 11)
        ? -1
        : (n % 4 == 0 || n % 4 == 2)
        ? n / 4
        : (n / 4 - 1);
}

int main() {
    scanf("%lld", &t); while (t--)
        scanf("%lld", &n), printf("%lld\n", get(n));
    return 0;
}
