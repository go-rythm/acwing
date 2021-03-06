## [4\. 多重背包问题 I](https://www.acwing.com/problem/content/4/)

### 题目

有 $N$ 种物品和一个容量是 $V$ 的背包。

第 $i$ 种物品最多有 $s\_i$ 件，每件体积是 $v\_i$，价值是 $w\_i$。

求解将哪些物品装入背包，可使物品体积总和不超过背包容量，且价值总和最大。

输出最大价值。

### 输入格式

第一行两个整数，$N，V$，用空格隔开，分别表示物品种数和背包容积。

接下来有 $N$ 行，每行三个整数 $v\_i, w\_i, s\_i$，用空格隔开，分别表示第 $i$ 种物品的体积、价值和数量。

### 输出格式

输出一个整数，表示最大价值。

### 数据范围

$0 \\lt N, V \\le 100$

$0 \\lt v\_i, w\_i, s\_i \\le 100$

### 输入样例

```
4 5
1 2 3
2 4 1
3 4 3
4 5 2
```

### 输出样例：

```
10
```

### 题解

前置题目：0003

前置知识：完全背包

本题知识：动态规划-背包问题

#### 题目分析

本题和完全背包问题的区别只有每件物品的数量是有限的。

且本题的数据范围较小，三重循环处理的规模是 10^6，故可以直接通过暴力算法解决问题

```go
for i := 1; i <= n; i++ {
    for j := 0; j <= m; j++ {
        for k := 0; k <= s[i] && k*v[i] <= j; k++ {
            f[i][j] = oj.Max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
        }
    }
}
```

