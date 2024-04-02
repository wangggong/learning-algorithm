#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;

const int K = 2000;
const ll mod = 998244353;
ll dp[2][3][K + 10];

int main() {
    int n, k;
    cin >> n >> k;
    dp[1][0][1] = 1, dp[1][1][2] = 2, dp[1][2][1] = 1;
    for (int i = 1; i < n; i++) {
        int curr = i % 2, next = (i + 1) % 2;
        memset(dp[next], 0, sizeof dp[next]);
        for (int k = 1; k <= i * 2; k++) {
            dp[next][0][k] = (dp[next][0][k] + dp[curr][0][k] + dp[curr][1][k]) % mod;
            dp[next][1][k] = (dp[next][1][k] + dp[curr][1][k]) % mod;
            dp[next][2][k] = (dp[next][2][k] + dp[curr][2][k] + dp[curr][1][k]) % mod;
            dp[next][0][k + 1] = (dp[next][0][k + 1] + dp[curr][2][k]) % mod;
            dp[next][1][k + 1] = (dp[next][1][k + 1] + (dp[curr][0][k] + dp[curr][2][k]) * 2) % mod;
            dp[next][2][k + 1] = (dp[next][2][k + 1] + dp[curr][0][k]) % mod;
            dp[next][1][k + 2] = (dp[next][1][k + 2] + dp[curr][1][k]) % mod;
        }
    }
    cout << ((dp[n % 2][0][k] + dp[n % 2][1][k] + dp[n % 2][2][k]) % mod) << endl;
    return 0;
}
