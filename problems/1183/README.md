## [1183\. 电力](https://www.acwing.com/problem/content/1185/)

### 题目

给定一个由 $n$ 个点 $m$ 条边构成的 **无向** 图，请你求出该图删除一个点之后，连通块最多有多少。

### 输入格式

输入包含多组数据。

每组数据第一行包含两个整数 $n,m$。

接下来 $m$ 行，每行包含两个整数 $a,b$，表示 $a,b$ 两点之间有边连接。

数据保证无重边。

点的编号从 $0$ 到 $n-1$。

读入以一行 $0\ 0$ 结束。

### 输出格式

每组数据输出一个结果，占一行，表示连通块的最大数量。

### 数据范围

$1 \\le n \\le 10000$,

$0 \\le m \\le 15000$,

$0 \\le a,b < n$

### 输入样例：

```
3 3
0 1
0 2
2 1
4 2
0 1
2 3
3 1
1 0
0 0
```

### 输出样例：

```
1
2
2
```

### 题解

