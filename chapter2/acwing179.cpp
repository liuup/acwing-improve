#pragma GCC optimize(2)

#include <bits/stdc++.h>
using namespace std;

// 提交记得取消注释!

// #define ACM_DEBUG

// 提交记得取消注释!

typedef long long LL;
typedef unsigned long long ULL;

#define PII pair<int,int>
#define all(a) a.begin(), a.end()

#define umap unordered_map
#define pq priority_queue

#define vi vector<int>
#define vvi vector<vector<int>>

auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};


// acwing 179

int dts[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

string start;
string ed = "12345678x";

string ops = "udlr";    // 和dts对应

bool findfg = false;    // 是否找到的标志

struct node {
    string s;
    int val;
};

struct op {
    string last;
    char p;
};

auto cmp = [](node a, node b){ return a.val > b.val; };

pq<node, vector<node>, decltype(cmp)> q(cmp);
umap<string, int> dist;
umap<string, op> path;

int distance(string s) {
    int ans = 0;
    for(int i = 0; i < s.size(); i++) {
        if(s[i] != ed[i]) ans++;
    }
    return ans;
}


int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(0);
    std::cout.tie(0);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif // ACM_DEBUG
    
    string tt;
    for(int i = 0; i < 9; i++) {
        cin >> tt;
        start += tt;
    }

    q.push(node{start, 0 + distance(start)});
    dist[start] = 0;
    path[start] = op{"", ' '};

    while(q.size()) {
        string cur = q.top().s;
        q.pop();

        if(cur == ed) {
            findfg = true;
            // std::cout << dist[cur] << endl;
            // return 0;
            break;
        }

        // 找到0的位置
        int idx = -1;
        for(int i = 0; i < cur.size(); i++) {
            if(cur[i] == 'x') {
                idx = i;
                break;
            }
        }

        int x = idx/3;
        int y = idx%3;

        for(int i = 0; i < 4; i++) {
            int dx = x + dts[i][0];
            int dy = y + dts[i][1];

            if(dx<0||dy<0||dx>=3||dy>=3) continue;

            string r = cur;
            int idxnew = dx*3+dy;
            swap(r[idx], r[idxnew]);

            if(!dist.count(r)) {
                dist[r] = dist[cur]+1;
                q.push(node{r, dist[r]+distance(r)});
                path[r] = op{cur, ops[i]};  
            }
        }
    }
    if(!findfg) {
        cout << "unsolvable" << endl;
        return 0;
    }

    string final;
    string ss = ed;
    while(path[ss].last.size() != 0) {
        final += path[ss].p;
        ss = path[ss].last;
    }

    // 再反转一下
    for(int i = 0; i < final.size()/2; i++) {
        swap(final[i], final[final.size()-i-1]);
    }

    cout << final << endl;
    
    return 0;
}
