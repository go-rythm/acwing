## [2189\. 有源汇上下界最大流](https://www.acwing.com/problem/content/2191/)

### 题目

给定一个包含 $n$ 个点 $m$ 条边的 **有向** 图，每条边都有一个流量下界和流量上界。

给定源点 $S$ 和汇点 $T$，求源点到汇点的最大流。

### 输入格式

第一行包含四个整数 $n,m,S,T$。

接下来 $m$ 行，每行包含四个整数 $a,b,c,d$ 表示点 $a$ 和 $b$ 之间存在一条有向边，该边的流量下界为 $c$，流量上界为 $d$。

点编号从 $1$ 到 $n$。

### 输出格式

输出一个整数表示最大流。

如果无解，则输出 `No Solution`。

### 数据范围

$1 \\le n \\le 202$,

$1 \\le m \\le 9999$,

$1 \\le a,b \\le n$,

$0 \\le c \\le d \\le 10^5$

### 输入样例：

```
10 15 9 10
9 1 17 18
9 2 12 13
9 3 11 12
1 5 3 4
1 6 6 7
1 7 7 8
2 5 9 10
2 6 2 3
2 7 0 1
3 5 3 4
3 6 1 2
3 7 6 7
5 10 16 17
6 10 10 11
7 10 14 15
```

### 输出样例：

```
43
```

### 题解

