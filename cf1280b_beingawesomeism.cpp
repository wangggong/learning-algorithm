#include <iostream>

using namespace std;

const int N = 60;
char M[N + 10][N + 10];
int n, m;

bool check0() {
    for (int i = 0; i < n; i++)
        for (int j = 0; j < m; j++)
            if (M[i][j] == 'P')
                return false;
    return true;
}

bool check1() {
    {
        bool all = true;
        for (int i = 0; i < n; i++) {
            if (M[i][0] == 'P') {
                all = false;
                break;
            }
        }
        if (all)
            return true;
    }
    {
        bool all = true;
        for (int i = 0; i < n; i++) {
            if (M[i][m - 1] == 'P') {
                all = false;
                break;
            }
        }
        if (all)
            return true;
    }
    {
        bool all = true;
        for (int i = 0; i < m; i++) {
            if (M[0][i] == 'P') {
                all = false;
                break;
            }
        }
        if (all)
            return true;
    }
    {
        bool all = true;
        for (int i = 0; i < m; i++) {
            if (M[n - 1][i] == 'P') {
                all = false;
                break;
            }
        }
        if (all)
            return true;
    }
    return false;
}

bool check2() {
    if (M[0][0] == 'A' || M[0][m - 1] == 'A' || M[n - 1][0] == 'A' || M[n - 1][m - 1] == 'A')
        return true;
    for (int i = 0; i < n; i++) {
        bool all = true;
        for (int j = 0; j < m; j++)
            if (M[i][j] == 'P') {
                all = false;
                break;
            }
        if (all)
            return true;
    }
    for (int i = 0; i < m; i++) {
        bool all = true;
        for (int j = 0; j < n; j++)
            if (M[j][i] == 'P') {
                all = false;
                break;
            }
        if (all)
            return true;
    }
    return false;
}

bool check3() {
    for (int i = 0; i < n; i++)
        if (M[i][0] == 'A' || M[i][m - 1] == 'A')
            return true;
    for (int i = 0; i < m; i++)
        if (M[0][i] == 'A' || M[n - 1][i] == 'A')
            return true;
    return false;
}

bool check4() {
    for (int i = 0; i < n; i++)
        for (int j = 0; j < m; j++)
            if (M[i][j] == 'A')
                return true;
    return false;
}

int main() {
    int t;
    cin >> t;
    while (t--) {
        cin >> n >> m;
        for (int i = 0; i < n; i++)
            for (int j = 0; j < m; j++)
                cin >> M[i][j];
        int v = -1;
        if (check0())
            v = 0;
        else if (check1())
            v = 1;
        else if (check2())
            v = 2;
        else if (check3())
            v = 3;
        else if (check4())
            v = 4;
        else {
            cout << "MORTAL" << endl;
            continue;
        }
        cout << v << endl;
    }
    return 0;
}
