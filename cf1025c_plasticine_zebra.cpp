#include <iostream>

using namespace std;

int main() {
    string s;
    cin >> s;
    s = s + s;
    int ans = 0;
    for (int p = 0, q = 0, n = s.size(); p < n; p = q) {
        for (q = p + 1; q < n && s[q] != s[q - 1]; q++);
        ans = max(ans, q - p);
    }
    cout << min(ans, int(s.size() / 2)) << endl;
    return 0;
}
