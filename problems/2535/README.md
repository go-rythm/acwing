## [2535\. 二次离线莫队](https://www.acwing.com/problem/content/2537/)

### 题目

给定一个长度为 $n$ 的序列 $a\_1,a\_2,…,a\_n$ 以及一个整数 $k$。

现在要进行 $m$ 次询问，每次询问给定一个区间 $[l,r]$，请你求出共有多少个数对 $(i,j)$ 满足 $l \\le i < j \\le r$ 且 $a\_i \\oplus a\_j$ 的结果在二进制表示下恰好有 $k$ 个 $1$。

注：$\\oplus$ 表示按位异或操作。

### 输入格式

第一行包含三个整数 $n,m,k$。

第二行包含 $n$ 个整数表示 $a\_1,a\_2,…,a\_n$。

接下来 $m$ 行，每行包含两个整数 $l,r$，表示一次询问。

### 输出格式

共 $m$ 行，每行输出一个查询的结果。

### 数据范围

$1 \\le n,m \\le 10^5$,

$0 \\le k \\le 14$,

$0 \\le a\_i < 2^{14}$,

$1 \\le l < r \\le n$

### 输入样例：

```
5 3 2
1 1 4 7 7
1 5
1 3
3 5
```

### 输出样例：

```
8
2
2
```

### 题解

