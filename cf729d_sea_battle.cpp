#include <iostream>

using namespace std;
typedef long long ll;

const ll N = 2e5;
ll n, a, b, k, ans[N + 10], idx;
string s;

int main() {
    cin >> n >> a >> b >> k >> s;
    for (ll i = 0; i < n; i++) {
        if (s[i] - '0')
            continue;
        ll p = i;
        for (; i < n && s[i] == '0'; i++);
        for (p += b - 1; p < i; p += b)
            ans[idx++] = p;
    }
    cout << idx - (a - 1) << endl;
    for (ll i = a - 1; i < idx; i++)
        cout << (ans[i] + 1) << ' ';
    return 0;
}
