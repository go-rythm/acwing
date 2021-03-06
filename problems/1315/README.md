## [1315\. 网格](https://www.acwing.com/problem/content/1317/)

### 题目

某城市的街道呈网格状，左下角坐标为 $A(0,0)$，右上角坐标为 $B(n,m)$，其中 $n \\ge m$。

现在从 $A(0,0)$ 点出发，只能沿着街道向正右方或者正上方行走，且不能经过图示中直线左上方的点，即任何途径的点 $(x,y)$ 都要满足 $x \\ge y$，请问在这些前提下，到达 $B(n,m)$ 有多少种走法。

![1.png](https://cdn.acwing.com/media/article/image/2020/01/03/19_90c130462d-1.png)

### 输入格式

仅有一行，包含两个整数 $n$ 和 $m$，表示城市街区的规模。

### 输出格式

输出一个整数，表示不同的方案总数。

### 数据范围

$1 \\le m \\le n \\le 5000$

### 输入样例：

```
6 6
```

### 输出样例：

```
132
```

### 题解

