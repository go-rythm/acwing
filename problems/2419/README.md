## [2419\. prufer序列](https://www.acwing.com/problem/content/2421/)

### 题目

本题需要你实现prufer序列与无根树之间的相互转化。

假设本题涉及的无根树共有 $n$ 个节点，编号 $1 \\sim n$。

为了更加简单明了的描述无根树的结构，我们不妨在输入和输出时将该无根树描述为一个以 $n$ 号节点为根的有根树。

这样就可以设这棵无根树的 **父亲序列** 为 $f\_1,f\_2,…,f\_{n-1}$，其中 $f\_i$ 表示将该树看作以 $n$ 号节点为根的有根树时，$i$ 号节点的父节点编号。

同时，设这棵无根树的 **[prufer序列](https://baike.baidu.com/item/prufer%E6%95%B0%E5%88%97/2182091?fr=aladdin)** 为 $p\_1,p\_2,…,p\_{n-2}$。

现在，给定一棵由 $n$ 个节点构成的无根树，以及该无根树的一个序列（父亲序列或prufer序列），请你求出另一个序列。

### 输入格式

输入共两行。

第一行包含两个整数 $n,m$，表示无根树的节点数以及给定序列类型。

如果 $m = 1$，则第二行包含 $n-1$ 个整数，表示给定序列为父亲序列。

如果 $m = 2$，则第二行包含 $n-2$ 个整数，表示给定序列为prufer序列。

### 输出格式

共一行，输出另一个序列，整数之间用单个空格隔开。

### 数据范围

$2 \\le n \\le 10^5$

### 输入样例1：

```
6 1
3 5 4 5 6
```

### 输出样例1：

```
3 5 4 5
```

### 输入样例2：

```
6 2
3 5 4 5
```

### 输出样例2：

```
3 5 4 5 6
```

### 题解

