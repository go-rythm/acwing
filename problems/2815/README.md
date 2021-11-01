## [2815\. 三维偏序](https://www.acwing.com/problem/content/2817/)

### 题目

给定 $n$ 个元素（编号 $1 \\sim n$），其中第 $i$ 个元素具有 $a\_i,b\_i,c\_i$ 三种属性。

设 $f(i)$ 表示满足以下 $4$ 个条件：

1. $a\_j \\le a\_i$
2. $b\_j \\le b\_i$
3. $c\_j \\le c\_i$
4. $j \\neq i$

的 $j$ 的数量。

对于 $d \\in [0,n)$，求满足 $f(i) = d$ 的 $i$ 的数量。

### 输入格式

第一行两个整数 $n,k$，表示元素数量和最大属性值。

接下来 $n$ 行，其中第 $i$ 行包含三个整数 $a\_i ,b\_i,c\_i$，分别表示第 $i$ 个元素的三个属性值。

### 输出格式

共 $n$ 行，每行输出一个整数，其中第 $d+1$ 行的整数表示满足 $f(i) = d$ 的 $i$ 的数量。

### 数据范围

$1 \\le n \\le 10^5$,

$1 \\le a\_i,b\_i,c\_i \\le k \\le 2 \\times 10^5$

### 输入样例：

```
10 3
3 3 3
2 3 3
2 3 1
3 1 1
3 1 2
1 3 1
1 1 2
1 2 2
1 3 2
1 2 1
```

### 输出样例：

```
3
1
3
0
1
0
1
0
0
1
```

### 题解
