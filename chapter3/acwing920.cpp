
// #define ACM_DEBUG   // comment this line when upload !!!

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

#define INF 0x3f3f3f3f

auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};


struct node {
    int to;
    double val;
};

int m, n;
string line;


int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> m >> n;

    vector<vector<int>> g(n+1);

    getline(cin, line);

    while(m--) {
        getline(cin, line);
        stringstream ss(line);

        int p;
        vector<int> tmp;
        while(ss >> p) {
            tmp.push_back(p);
        }

        // 各个点之间都要链接上
        for(int i = 0; i < tmp.size(); i++) {
            for(int j = i+1; j < tmp.size(); j++) {
                g[tmp[i]].push_back(tmp[j]);
            }
        }
    }

    // 求一下bfs深度
    queue<int> q;
    vector<bool> vis(n+1, false);
    
    q.push(1);
    vis[1] = true;

    int step = 0;
    while(q.size()) {
        int size = q.size();
        for(int i = 0; i < size; i++) {
            int cur = q.front(); q.pop();
            vis[cur] = true;

            if(cur == n) {
                cout << step-1 << endl;
                return 0;
            }
            
            for(auto e : g[cur]) {
                if(vis[e]) {
                    continue;
                }
                q.push(e);
            }
        }
        step++;
    }

    cout << "NO" << endl;

    return 0;
}
