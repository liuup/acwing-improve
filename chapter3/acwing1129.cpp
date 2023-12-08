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
auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};


struct node {
    int to, val;
};


int t, c, ts, te;

int rs, re, ci;

vector<vector<node>> g;
vector<int> dist;


int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> t >> c >> ts >> te;
    g = vector<vector<node>>(t+1);
    dist = vector<int>(t+1, INT_MAX);
    while(c--) {
        cin >> rs >> re >> ci;
        g[rs].push_back(node{re, ci});
        g[re].push_back(node{rs, ci});
    }

    auto cmp = [](node a, node b){ return a.val > b.val; };
    priority_queue<node, vector<node>, decltype(cmp)> q(cmp);

    q.push(node{ts, 0});
    dist[ts] = 0;

    while(q.size()) {
        auto cur = q.top(); q.pop();
        
        if(dist[cur.to] < cur.val) continue;

        for(auto e : g[cur.to]) {
            int d = dist[cur.to] + e.val;
            if(d < dist[e.to]) {
                dist[e.to] = d;
                q.push(node{e.to, d});
            }
        }
    }

    cout << dist[te] << endl;

    return 0;
}
