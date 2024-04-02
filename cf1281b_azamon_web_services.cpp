#include <iostream>
#include <cstring>

using namespace std;

const int ALPHABET = 'Z' - 'A' + 1;
int t;
string s, c;
int last_occur[ALPHABET], first_larger[ALPHABET];

string getMin(string s) {
    int n = s.size();
    memset(last_occur, -1, sizeof last_occur);
    memset(first_larger, 0x3f, sizeof first_larger);
    for (int i = 0; i < n; i++) {
        last_occur[s[i] - 'A'] = i;
        for (char c = 'A'; c < s[i]; c++)
            first_larger[c - 'A'] = min(first_larger[c - 'A'], i);
    }
    for (char c = 'A'; c <= 'Z'; c++)
        if (last_occur[c - 'A'] > first_larger[c - 'A']) {
            swap(s[last_occur[c - 'A']], s[first_larger[c - 'A']]);
            break;
        }
    return s;
}

int main() {
    cin >> t;
    while (t--) {
        cin >> s >> c;
        if (s < c) {
            cout << s << endl;
            continue;
        }
        s = getMin(s);
        cout << (s < c ? s : "---") << endl;
    }
    return 0;
}
