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


struct node {
    int to, val;
};

int n, m;
int a, b, l;
int s, t, k;

typedef pair<int, PII> PIII;

vector<vector<node>> g1;    // 正向图
vector<vector<node>> g2;    // 反向图
vi dist;

int cnt[1010];

int astar() {
    priority_queue<PIII, vector<PIII>, greater<PIII>> heap;
    heap.push({dist[s], {0, s}});   // 谁的d[u]+f[u]更小 谁先出队列
    
    while(heap.size()) {
        auto tt = heap.top(); heap.pop();
        int ver = tt.second.second, distance = tt.second.first;
        cnt[ver]++;
        if(cnt[t] == k) return distance;    // 终点被访问过k次 那此时的ver就是终点t 返回答案

        for(auto e : g1[ver]) { // 所有孩子
            if(cnt[e.to] < k) {
                heap.push({distance+e.val+dist[e.to], {distance+e.val, e.to}});
            }
        }
    }

    return -1;
}

int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> m;

    // 以t为起点 跑一遍dijkstra到各个点的最短路径
    g1 = vector<vector<node>>(n+1);
    g2 = vector<vector<node>>(n+1);
    dist = vi(n+1, INT_MAX);

    while(m--) {
        cin >> a >> b >> l;
        g1[a].push_back(node{b, l});
        g2[b].push_back(node{a, l});
    }

    cin >> s >> t >> k;
    if(s == t) k++; // 起点==终点 则d[S->S] = 0这种情况要舍去 总数第k大变为总数第k-1大

    // 在逆向图跑dijkstra 作为估计函数f[u]
    auto cmp = [](node a, node b){ return a.val > b.val; };
    priority_queue<node, vector<node>, decltype(cmp)> hp(cmp);

    hp.push(node{t, 0});
    dist[t] = 0;

    while(hp.size()) {
        auto cur = hp.top(); hp.pop();

        if(dist[cur.to] < cur.val) continue;

        for(auto e : g2[cur.to]) {
            int d = dist[cur.to] + e.val;
            if(d < dist[e.to]) {
                dist[e.to] = d;
                hp.push(node{e.to, d});
            }
        }
    }

    // astar
    cout << astar() << endl;

    return 0;
}
