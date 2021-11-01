## [1137\. 选择最佳线路](https://www.acwing.com/problem/content/1139/)

### 题目

有一天，琪琪想乘坐公交车去拜访她的一位朋友。

由于琪琪非常容易晕车，所以她想尽快到达朋友家。

现在给定你一张城市交通路线图，上面包含城市的公交站台以及公交线路的具体分布。

已知城市中共包含 $n$ 个车站（编号$1$~$n$）以及 $m$ 条公交线路。

每条公交线路都是 **单向的**，从一个车站出发直接到达另一个车站，两个车站之间可能存在多条公交线路。

琪琪的朋友住在 $s$ 号车站附近。

琪琪可以在任何车站选择换乘其它公共汽车。

请找出琪琪到达她的朋友家（附近的公交车站）需要花费的最少时间。

### 输入格式

输入包含多组测试数据。

每组测试数据第一行包含三个整数 $n,m,s$，分别表示车站数量，公交线路数量以及朋友家附近车站的编号。

接下来 $m$ 行，每行包含三个整数 $p,q,t$，表示存在一条线路从车站 $p$ 到达车站 $q$，用时为 $t$。

接下来一行，包含一个整数 $w$，表示琪琪家附近共有 $w$ 个车站，她可以在这 $w$ 个车站中选择一个车站作为始发站。

再一行，包含 $w$ 个整数，表示琪琪家附近的 $w$ 个车站的编号。

### 输出格式

每个测试数据输出一个整数作为结果，表示所需花费的最少时间。

如果无法达到朋友家的车站，则输出 -1。

每个结果占一行。

### 数据范围

$n \\le 1000, m \\le 20000$,

$1 \\le s \\le n$,

$0 < w < n$,

$0 < t \\le 1000$

### 输入样例：

```
5 8 5
1 2 2
1 5 3
1 3 4
2 4 7
2 5 6
2 3 5
3 5 1
4 5 1
2
2 3
4 3 4
1 2 3
1 3 4
2 3 2
1
1
```

### 输出样例：

```
1
-1
```

### 题解
