#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;

string s;
int n;
ll dp[2][10];

int main() {
    cin >> s; n = s.size();
    for (int i = 0; i <= 9; i++)
        dp[0][i] = 1;
    for (int i = 1; i < n; i++) {
        int pre = (i - 1) % 2, cur = i % 2;
        memset(dp[cur], 0, sizeof dp[cur]);
        int k = s[i] - '0';
        for (int j = 0; j <= 9; j++) {
            dp[cur][(j + k) / 2] += dp[pre][j];
            if ((j + k) % 2)
                dp[cur][(j + k) / 2 + 1] += dp[pre][j];
        }
    }
    ll ans = 0;
    for (int i = 0; i <= 9; i++)
        ans += dp[(n - 1) % 2][i];
    bool canbuild = true;
    for (int curr = s[0] - '0', next = s[1] - '0', i = 1; i < n; curr = next, i++, next = s[i] - '0')
        if (!((curr + next) / 2 == next || ((curr + next) % 2 != 0 && (curr + next) / 2 + 1 == next)))
            canbuild = false;
    cout << (ans + (canbuild ? -1 : 0)) << endl;
    return 0;
}
