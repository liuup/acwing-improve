
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
#define pb push_back

#define inf 0x3f3f3f3f

auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};
void printb(int a) { cout << bitset<sizeof(a)*8>(a) << endl; }  // 打印数字对应的二进制


// fenwick tree 树状数组
class Fenwick {
public:
    vector<int> nums;   // 原始数组
    vector<int> c;  // 下标从1开始的

    Fenwick();  // 默认构造函数 防止报错
    Fenwick(vector<int> nums);  // 数据的下标从1开始的
    void update(int x, int k); // x下标 k修改的数值    
    int getsum(int x); // [c[1], c[x]]累计 闭区间
    int rangesum(int left, int right);   // [c[left], c[right]]累计 闭区间

    int lowbit(int x);

    int ask(int x);
};

Fenwick::Fenwick() {}

Fenwick::Fenwick(vector<int> nums) {
    this->nums = nums;
    this->c = vector<int>(nums.size()+1, 0); 
}

void Fenwick::update(int x, int k) {
    while(x < this->c.size()) {
        c[x] += k;
        x += this->lowbit(x);
    }
}

int Fenwick::getsum(int x) {
    int ans = nums[x];
    while(x > 0) {
        ans += c[x];
        x -= this->lowbit(x);
    }
    return ans;
}

int Fenwick::rangesum(int left, int right) {
    return this->getsum(right) - this->getsum(left-1);
}

int Fenwick::lowbit(int x) {
    return x & (-x);
}

int n, m;
string q;
int a;
int l, r, d;

// 差分树状数组

int main(void) {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> m;
    vector<int> nums(n+1);
    for(int i = 1; i <= n; i++) {
        cin >> nums[i];
    }

    auto fw = Fenwick(nums);

    for(int i = 0; i < m; i++) {
        cin >> q;
        if(q == "Q") {
            cin >> a;
            cout << fw.getsum(a) << endl;
        } else if(q == "C") {
            cin >> l >> r >> d;
            fw.update(l, d);
            fw.update(r+1, -d);
        }
    }

    return 0;
}