## [2863\. 最短路](https://www.acwing.com/problem/content/2866/)

### 题目

给一个 $N$ 个点 $M$ 条边的仙人掌。仙人掌定义如下：

> 任意一条边至多只出现在一条简单回路的无向连通图称为仙人掌。

有 $Q$ 组询问，每次询问两点之间的最短路径。

### 输入格式

第一行包含三个整数，分别表示 $N,M,Q$。

下接 $M$ 行，每行三个整数 $v,u,w$ 表示一条无向边 $v-u$，长度为 $w$。

最后 $Q$ 行，每行两个整数 $v,u$ 表示一组询问。

### 输出格式

输出 $Q$ 行，每行一个整数表示询问的答案。

### 数据范围

$1 \\le N \\le 10000$,

$1 \\le M \\le 12000$,

$1 \\le Q \\le 10000$

### 输入样例：

```
9 10 2
1 2 1
1 4 1
3 4 1
2 3 1
3 7 1
7 8 2
7 9 2
1 5 3
1 6 4
5 6 1
1 9
5 7
```

### 输出样例：

```
5
6
```

### 题解

