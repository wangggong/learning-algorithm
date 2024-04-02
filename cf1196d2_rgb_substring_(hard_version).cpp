#include <iostream>

using namespace std;

string RGB = "RGB";

int main() {
    int t;
    cin >> t;
    while (t--) {
        int k, n, ans = INT_MAX;
        string s;
        cin >> n >> k;
        cin >> s;
        for (int st = 0; st < 3; st++) {
            int cur = 0;
            for (int i = 0; i < k; i++)
                if (RGB[(st + i) % 3] != s[i])
                    cur++;
            ans = min(ans, cur);
            for (int i = k; i < n; i++) {
                if (RGB[(st + i - k) % 3] != s[i - k])
                    cur--;
                if (RGB[(st + i) % 3] != s[i])
                    cur++;
                ans = min(ans, cur);
            }
        }
        cout << ans << endl;
    }
    return 0;
}
