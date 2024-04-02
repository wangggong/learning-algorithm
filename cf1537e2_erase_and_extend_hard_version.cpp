#include <iostream>

using namespace std;

int main() {
    int n, k;
    string s, ans;
    cin >> n >> k >> s;
    int sz = 1;
    for (int i = 0; i < s.size(); i++) {
        if (s[i] > s[i % sz])
            break;
        if (s[i] < s[i % sz])
            sz = i + 1;
    }
    for (; ans.size() < k; ans = ans + s.substr(0, sz));
    cout << ans.substr(0, k) << endl;
    return 0;
}
