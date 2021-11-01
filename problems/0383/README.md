## [383\. 观光](https://www.acwing.com/problem/content/385/)

### 题目

“您的个人假期”旅行社组织了一次比荷卢经济联盟的巴士之旅。

比荷卢经济联盟有很多公交线路。

每天公共汽车都会从一座城市开往另一座城市。

沿途汽车可能会在一些城市（零或更多）停靠。

旅行社计划旅途从 $S$ 城市出发，到 $F$ 城市结束。

由于不同旅客的景点偏好不同，所以为了迎合更多旅客，旅行社将为客户提供多种不同线路。

游客可以选择的行进路线有所限制，要么满足所选路线总路程为 $S$ 到 $F$ 的最小路程，要么满足所选路线总路程仅比最小路程多一个单位长度。

![3463_1.png](/media/article/image/2019/02/26/19_75361c2839-3463_1.png)

如上图所示，如果 $S = 1，F = 5$，则这里有两条最短路线 $1 \\to 2 \\to 5,1 \\to 3 \\to 5$，长度为 $6$；有一条比最短路程多一个单位长度的路线 $1 \\to 3 \\to 4 \\to 5$，长度为 $7$。

现在给定比荷卢经济联盟的公交路线图以及两个城市 $S$ 和 $F$，请你求出旅行社最多可以为旅客提供多少种不同的满足限制条件的线路。

### 输入格式

第一行包含整数 $T$，表示共有 $T$ 组测试数据。

每组数据第一行包含两个整数 $N$ 和 $M$，分别表示总城市数量和道路数量。

接下来 $M$ 行，每行包含三个整数 $A,B,L$，表示有一条线路从城市 $A$ 通往城市 $B$，长度为 $L$。

需注意，线路是 **单向的**，存在从 $A$ 到 $B$ 的线路不代表一定存在从 $B$ 到 $A$ 的线路，另外从城市 $A$ 到城市 $B$ 可能存在多个不同的线路。

接下来一行，包含两个整数 $S$ 和 $F$，数据保证 $S$ 和 $F$ 不同，并且 $S、F$ 之间至少存在一条线路。

### 输出格式

每组数据输出一个结果，每个结果占一行。

数据保证结果不超过 $10^9$。

### 数据范围

$2 \\le N \\le 1000$,

$1 \\le M \\le 10000$,

$1 \\le L \\le 1000$，

$1 \\le A,B,S,F \\le N$

### 输入样例：

```
2
5 8
1 2 3
1 3 2
1 4 5
2 3 1
2 5 3
3 4 2
3 5 4
4 5 3
1 5
5 6
2 3 1
3 2 1
3 1 10
4 5 2
5 2 7
5 2 7
4 1
```

### 输出样例：

```
3
2
```

### 题解
