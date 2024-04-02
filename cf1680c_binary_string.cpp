// 参考: https://www.luogu.com.cn/blog/endlesscheng/solution-cf1680c
//
// 二分容易, 直接滑窗考虑一下?
#include <iostream>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)

using namespace std;
typedef long long ll;
ll t, n, tot1;
string s;

bool check(ll k) {
    ll count0 = 0, count1 = 0;
    for (ll p = 0, q = 0; p < n; p++) {
        for (; q < n && count0 <= k; q++)
            if (s[q] == '0')
                count0++;
            else
                count1++;
        if (tot1 - count1 <= k)
            return true;
        if (s[p] == '0')
            count0--;
        else
            count1--;
    }
    return false;
}

int main() {
    __FAST__;
    for (cin >> t; t; t--) {
        cin >> s;
        n = s.size(), tot1 = 0;
        for (ll i = 0; i < n; i++)
            tot1 += s[i] - '0';
        ll p = 0, q = n;
        while (p < q) {
            ll mid = (p + q) >> 1;
            if (check(mid))
                q = mid;
            else
                p = mid + 1;
        }
        cout << p << endl;
    }
    return 0;
}
