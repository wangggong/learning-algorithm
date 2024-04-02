#include <iostream>

using namespace std;

int main() {
    // 1*2*3*4 = 24
    // 5*4+ = 24
    // 5*6-(1-(2-3)+4) = 24
    // 5*6-7+(3-4)*(1-2) = 24
    int n;
    cin >> n;
    if (n < 4) {
        cout << "NO" << endl;
        return 0;
    }
    cout << "YES" << endl;
    int k = (n % 4) + 4;
    switch (k) {
    case 4:
        cout << "1 * 2 = 2" << endl;
        cout << "2 * 3 = 6" << endl;
        cout << "6 * 4 = 24" << endl;
        break;
    case 5:
        cout << "5 * 4 = 20" << endl;
        cout << "2 - 1 = 1" << endl;
        cout << "1 + 3 = 4" << endl;
        cout << "20 + 4 = 24" << endl;
        break;
    case 6:
        cout << "1 + 2 = 3" << endl;
        cout << "3 + 4 = 7" << endl;
        cout << "7 - 5 = 2" << endl;
        cout << "2 + 6 = 8" << endl;
        cout << "3 * 8 = 24" << endl;
        break;
    case 7:
        cout << "5 * 6 = 30" << endl;
        cout << "30 - 7 = 23" << endl;
        cout << "1 - 2 = -1" << endl;
        cout << "3 - 4 = -1" << endl;
        cout << "-1 * -1 = 1" << endl;
        cout << "1 + 23 = 24" << endl;
        break;
    default:
        // unreachable
        break;
    }
    if (n == k)
        return 0;
    int cnt1 = 0;
    for (int i = k + 1; i < n; i += 4) {
        cout << i << " - " << (i + 1) << " = -1" << endl;
        cout << (i + 2) << " - " << (i + 3) << " = -1" << endl;
        cout << "-1 * -1 = 1" << endl;
        cnt1++;
    }
    while (--cnt1)
        cout << "1 * 1 = 1" << endl;
    if (n > k)
        cout << "24 * 1 = 24" << endl;
    return 0;
}

