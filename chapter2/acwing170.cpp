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



int n, d;   // d为搜索深度

int a[10005];   // 存储加成序列


bool dfs(int u) {   // 搜索第u层
    if(u == d) return a[u-1] == n;
    for(int i = u-1; i >= 0; i--) { // 逆序搜索
        int t = a[u-1] + a[i];
        if(t > n) continue; // 越界剪枝
        a[u] = t;
        for(int j = u+1; j <= d; j++) t *= 2;
        if(t < n) return false; // 估价未来
        if(dfs(u+1)) return true;
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

    a[0] = 1;

    while(scanf("%d", &n), n) {
        d = 1;
        while(!dfs(1)) d++;    // 失败则增加一层
        for(int i = 0; i < d; i++) cout << a[i] << " ";
        cout << endl;
    }

    return 0;
}
