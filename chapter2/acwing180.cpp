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


const int N = 15;

int q[N];
int w[5][N];

int t;
int n;


int f() {
    int ans = 0;
    for(int i = 0; i < n-1; i++) {
        if(q[i+1] != q[i]+1) {
            ans++;
        }
    }
    return (ans + 2)/3;
}

bool check() {
    for(int i = 0; i < n; i++) {    // 判断是否有序
        if(q[i] != i+1) {
            return false;
        }
    }
    return true;
}

bool dfs(int depth, int max_depth) {
    if(depth + f() > max_depth) {  // 超过估计的最大层数
        return false;
    }
    if(check()) return true;

    for(int l = 0; l < n; l++) {
        for(int r = l; r < n; r++) {
            for(int k = r+1; k < n; k++) {
                memcpy(w[depth], q, sizeof(q));
                int x, y;
                for(x = r+1, y = l; x <= k; x++, y++) q[y] = w[depth][x];
                for(x = l; x <= r; x++, y++) q[y] = w[depth][x];
                if(dfs(depth+1, max_depth)) return true;
                memcpy(q, w[depth], sizeof(q));
            }
        }
    }
    return false;
}


int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> t;
    while(t--) {
        cin >> n;
        
        for(int i = 0; i < n; i++) cin >> q[i];

        auto depth = 0;
        while(depth < 5 && !dfs(0, depth)) {
            depth++;
        }
        if(depth >= 5) {
            cout << "5 or more" << endl;
        } else{
            cout << depth << endl;
        }
    }

    return 0;
}
