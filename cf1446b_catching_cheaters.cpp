#include <iostream>

using namespace std;

const int N = 5e3;

int dp[N + 10][N + 10];

int main() {
    int n, m;
    cin >> n >> m;
    string s, t;
    cin >> s >> t;
    int ans = 0;
    for (int i = 0; i < n; i++)
        for (int j = 0; j < m; j++) {
            dp[i + 1][j + 1] = s[i] == t[j] ? dp[i][j] + 2 : max(max(dp[i + 1][j], dp[i][j + 1]) - 1, 0);
            ans = max(ans, dp[i + 1][j + 1]);
        }
    cout << ans << endl;
    return 0;
}
