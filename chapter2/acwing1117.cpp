// comment this line when upload

// #define ACM_DEBUG


#pragma GCC optimize(2)

#include <bits/stdc++.h>
using namespace std;

#define LL long long
#define ULL unsigned long long

#define PII pair<int,int>
#define all(a) a.begin(), a.end()

#define umap unordered_map
#define pq priority_queue

#define vi vector<int>
#define vvi vector<vector<int>>
auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};

const int N = 25;

int n, ans;

string s[N];
int st[N];


void dfs(string x, int y) {
    st[y]++;
    int len = x.size();
    ans = max(ans, len);

    string t;
    for(int i = 0; i < n; i++) {
        for(int j = len-1, k=1; j > 0 && k < s[i].size(); j--, k++) {
            if(st[i] < 2 && x.substr(j) == s[i].substr(0, k)) {
                t = x.substr(0, len-k) + s[i];
                dfs(t, i);
            }
        }
    }
    st[y]--;
}


int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(0);
    std::cout.tie(0);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif
    
    cin >> n;
    for(int i = 0; i < n; i++) {
        cin >> s[i];
    }

    string start;
    cin >> start;
    start = " " + start;
    dfs(start, n);

    cout << ans -1 << endl;


    return 0;
}
