#include <iostream>

using namespace std;

const int N = 2e5;
int prefix[N + 10], suffix[N + 10], n, m, ans;

int main() {
    string s, t;
    cin >> s >> t;
    n = s.size(), m = t.size();
    for (int i = 0, j = 0; i < n && j < m; i++)
        if (s[i] == t[j])
            prefix[j] = i, j++;
    for (int i = n - 1, j = m - 1; i >= 0 && j >= 0; i--)
        if (s[i] == t[j])
            suffix[j] = i, j--;
    // b b
    // 0 1
    // 1 3
    // for (int i = 0; i < m; i++)
    //     cout << prefix[i] << ' '; cout << endl;
    // for (int i = 0; i < m; i++)
    //     cout << suffix[i] << ' '; cout << endl;
    ans = max(n - prefix[m - 1] - 1, suffix[0]);
    for (int i = 0; i + 1 < m; i++)
        ans = max(ans, suffix[i + 1] - prefix[i] - 1);
    cout << ans << endl;
    return 0;
}
