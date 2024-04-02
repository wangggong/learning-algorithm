#include <iostream>
#include <vector>

using namespace std;

typedef pair<char, int> pci;

/*
 * int main() {
 *     int n;
 *     string s;
 *     cin >> n >> s;
 *     vector<pci> vec;
 *     for (int p = 0, q = 0, n = s.size(); p < n; p = q) {
 *         for (q = p + 1; q < n && s[q] == s[p]; q++);
 *         vec.push_back({s[p], q - p});
 *     }
 *     int gcnt = 0;
 *     for (auto v : vec)
 *         if (v.first == 'G')
 *             gcnt++;
 *     int ans = 0;
 *     for (int i = 0; i < vec.size(); i++) {
 *         if (vec[i].first == 'G') {
 *             ans = max(ans, vec[i].second + (gcnt == 1 ? 0 : 1));
 *             continue;
 *         }
 *         if (i > 0 && i + 1 < vec.size()) {
 *             if (vec[i].second == 1)
 *                 ans = max(ans, vec[i - 1].second + vec[i + 1].second + (gcnt == 2 ? 0 : 1));
 *         }
 *     }
 *     cout << ans << endl;
 *     return 0;
 * }
 */

int main() {
    int n;
    string s;
    cin >> n >> s;
    // `a`: current length of "G"'s;
    // `b`: last length of "G"'s;
    // `g`: total count of "G"'s;
    int a = 0, b = 0, g = 0, ans = 0;
    for (char c : s) {
        if (c == 'G')
            g++, a++;
        else
            b = a, a = 0;
        ans = max(ans, a + b + 1);
    }
    cout << min(ans, g) << endl;
    return 0;
}
