#include <iostream>

using namespace std;

string s;

int main() {
    cin >> s;
    int diff = 0, cnt = 0, min_diff = INT_MAX;
    for (char c : s)
        if (c == '(')
            diff++;
        else {
            diff--;
            if (diff < 0) {
                cout << -1 << endl;
                return 0;
            }
            min_diff = min(min_diff, diff);
            if (c == '#')
                cnt++, min_diff = diff;
        }
    if (min_diff < diff) {
        cout << -1 << endl;
        return 0;
    }
    for (int i = 0; i < cnt; i++)
        cout << (i == cnt - 1 ? diff + 1 : 1) << endl;
    return 0;
}
