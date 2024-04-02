#include <iostream>

using namespace std;
typedef long long ll;

const int mod = 1e9 + 7;

int main() {
    int n;
    cin >> n;
    ll dp[4] = {1, 0, 0, 0};
    while (n--) {
        ll tmp[4] = {dp[0], dp[1], dp[2], dp[3]};
        dp[0] = (tmp[1] + tmp[2] + tmp[3]) % mod;
        dp[1] = (tmp[0] + tmp[2] + tmp[3]) % mod;
        dp[2] = (tmp[0] + tmp[1] + tmp[3]) % mod;
        dp[3] = (tmp[0] + tmp[1] + tmp[2]) % mod;
    }
    cout << dp[0] << endl;
    return 0;
}
