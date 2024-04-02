#include <iostream>

using namespace std;

string s;
bool debug = false;

string handleFloat(string s);
string handleInt(string s);

string handleFloat(string s) {
    if (debug)
        cout << "handleFloat " << s << endl;
    if (s.empty())
        return "0";
    int i = 0;
    for (; i < s.size() && s[i] == '0'; i++);
    if (i)
        return handleFloat(s.substr(i));
    i = s.find(".");
    if (i < 0)
        return handleInt(s);
    for (; !s.empty() && s.back() == '0'; s.pop_back());
    if (s.back() == '.')
        return handleInt(s.substr(0, s.size() - 1));
    if (i == 0) {
        for (i = 1; i < s.size() && s[i] == '0'; i++);
        return s.substr(i, 1) + (s.substr(i + 1).empty() ? "" : "." + s.substr(i + 1)) + "E" + to_string(-i);
    }
    if (i == 1)
        return s;
    string floatPart = s.substr(1).erase(i - 1, 1);
    return s.substr(0, 1) + (floatPart.empty() ? "" : "." + floatPart) + "E" + to_string(i - 1);
}

string handleInt(string s) {
    if (debug)
        cout << "handleInt " << s << endl;
    if (s.empty())
        return "0";
    if (s.size() == 1)
        return s;
    return handleFloat(s.substr(0, 1) + (s.substr(1).empty() ? "" : "." + s.substr(1))) + "E" + to_string(s.size() - 1);
}

int main() {
    cin >> s;
    cout << handleFloat(s) << endl;
    return 0;
}
