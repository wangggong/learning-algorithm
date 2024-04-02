#include <iostream>

using namespace std;
typedef long long ll;

const int N = 1e5;
const int mod = 1e9 + 7;
ll a[N + 10], dp[N + 10], idx;

int main() {
    string s;
    cin >> s;
    for (int p = 0, q = 0, n = s.size(); p < n; p = q, idx++) {
        for (; q < n && s[q] == 'b'; q++);
        p = q;
        for (; q < n && s[q] != 'b'; q++)
            if (s[q] == 'a')
                a[idx]++;
    }
    for (int i = 0; i < idx; i++)
        dp[i + 1] = (dp[i] + a[i] * (1 + dp[i])) % mod;
    // for (int i = 0; i <= idx; i++)
    //     cout << a[i] << " " << dp[i] << endl;
    cout << dp[idx] << endl;
    return 0;
}
