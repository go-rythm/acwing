## [178\. 第K短路](https://www.acwing.com/problem/content/180/)

### 题目

给定一张 $N$ 个点（编号 $1,2…N$），$M$ 条边的有向图，求从起点 $S$ 到终点 $T$ 的第 $K$ 短路的长度，路径允许重复经过点或边。

**注意：** 每条最短路中至少要包含一条边。

### 输入格式

第一行包含两个整数 $N$ 和 $M$。

接下来 $M$ 行，每行包含三个整数 $A,B$ 和 $L$，表示点 $A$ 与点 $B$ 之间存在有向边，且边长为 $L$。

最后一行包含三个整数 $S,T$ 和 $K$，分别表示起点 $S$，终点 $T$ 和第 $K$ 短路。

### 输出格式

输出占一行，包含一个整数，表示第 $K$ 短路的长度，如果第 $K$ 短路不存在，则输出 $-1$。

### 数据范围

$1 \\le S,T \\le N \\le 1000$,

$0 \\le M \\le 10^5$,

$1 \\le K \\le 1000$,

$1 \\le L \\le 100$

### 输入样例：

```
2 2
1 2 5
2 1 4
1 2 2
```

### 输出样例：

```
14
```

### 题解
