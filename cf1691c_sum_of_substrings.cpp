#include <iostream>
#include <cstring>

using namespace std;

const int N = 1e5;
int dp[N + 10];

/*
 * int main() {
 *     int t;
 *     cin >> t;
 *     while (t--) {
 *         memset(dp, 0, sizeof dp);
 *         int n, k;
 *         string s;
 *         cin >> n >> k >> s;
 *         if (n == 1) {
 *             cout << (s[0] - '0') << endl;
 *             continue;
 *         }
 *         if (s.find("1") == string::npos) {
 *             cout << 0 << endl;
 *             continue;
 *         }
 *         int ans = (s[0] - '0') * 10 + (s[n - 1] - '0');
 *         for (int i = 1; i + 1 < n; i++)
 *             ans += 11 * (s[i] - '0');
 *         int first = 0, last = n - 1;
 *         for (; first < n && s[first] == '0'; first++);
 *         for (; last >= 0 && s[last] == '0'; last--);
 *         if (first == last) {
 *             if (last == n - 1) {
 *                 cout << ans << endl;
 *                 continue;
 *             }
 *             if (s[n - 1] == '0' && k >= n - 1 - last)
 *                 ans -= first == 0 ? 9 : 10;
 *             else if (s[0] == '0' && k >= first)
 *                 ans -= 1;
 *             else;
 *             cout << ans << endl;
 *             continue;
 *         }
 *         if (s[n - 1] == '0' && k >= n - 1 - last)
 *             k -= n - 1 - last, ans -= 10;
 *         if (s[0] == '0' && k >= first)
 *             k -= first, ans -= 1;
 *         cout << ans << endl;
 *     }
 *     return 0;
 * }
 */

int main() {
    int t;
    cin >> t;
    while (t--) {
        int n, k;
        string s;
        cin >> n >> k >> s;
        int ones = 0, add = 0, first = n, last = 0;
        for (int i = 0; i < n; i++)
            if (s[i] == '1')
                ones++, last = i, first = first == n ? i : first;
        if (ones > 0 && k >= n - 1 - last)
            k -= n - 1 - last, add += 1, ones--;
        if (ones > 0 && k >= first)
            k -= first, add += 10, ones--;
        cout << ones * 11 + add << endl;
    }
    return 0;
}
