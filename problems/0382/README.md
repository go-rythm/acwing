## [382\. K取方格数](https://www.acwing.com/problem/content/384/)

### 题目

在一个 $N \\times N$ 的矩形网格中，每个格子里都写着一个非负整数。

可以从左上角到右下角安排 $K$ 条路线，每一步只能往下或往右，沿途经过的格子中的整数会被取走。

若多条路线重复经过一个格子，只取一次。

求能取得的整数的和最大是多少。

### 输入格式

第一行包含两个整数 $N$ 和 $K$。

接下来 $N$ 行，每行包含 $N$ 个不超过 $1000$ 的整数，用来描述整个矩形网格。

### 输出格式

输出一个整数，表示能取得的最大和。

### 数据范围

$1 \\le N \\le 50$,

$0 \\le K \\le 10$

### 输入样例：

```
3 2
1 2 3
0 2 1
1 4 2
```

### 输出样例：

```
15
```

### 题解

