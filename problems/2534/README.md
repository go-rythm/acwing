## [2534\. 树上计数2](https://www.acwing.com/problem/content/2536/)

### 题目

给定一棵 $N$ 个节点的树，节点编号从 $1$ 到 $N$，每个节点都有一个整数权值。

现在，我们要进行 $M$ 次询问，格式为 `u v`，对于每个询问你需要回答从 $u$ 到 $v$ 的路径上（包括两端点）共有多少种不同的点权值。

### 输入格式

第一行包含两个整数 $N,M$。

第二行包含 $N$ 个整数，其中第 $i$ 个整数表示点 $i$ 的权值。

接下来 $N-1$ 行，每行包含两个整数 $x,y$，表示点 $x$ 和点 $y$ 之间存在一条边。

最后 $M$ 行，每行包含两个整数 $u,v$，表示一个询问。

### 输出格式

共 $M$ 行，每行输出一个询问的答案。

### 数据范围

$1 \\le N \\le 40000$,

$1 \\le M \\le 10^5$,

$1 \\le x,y,u,v \\le N$,

各点权值均在 $int$ 范围内。

### 输入样例：

```
8 2
105 2 9 3 8 5 7 7
1 2
1 3
1 4
3 5
3 6
3 7
4 8
2 5
3 8
```

### 输出样例：

```
4
4
```

### 题解

