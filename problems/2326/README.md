## [2326\. 王者之剑](https://www.acwing.com/problem/content/2328/)

### 题目

给出一个 $n \\times m$ 网格，每个格子上有一个价值 $v\_{i,j}$ 的宝石。

Amber 可以自己决定起点，开始时刻为第 $0$ 秒。

以下操作，在每秒内按顺序执行。

1. 若第 $i$ 秒开始时，Amber 在 $(x,y)$，则 Amber 可以拿走 $(x,y)$ 上的宝石。
2. 在偶数秒时（$i$ 为偶数），则 Amber 周围 $4$ 格的宝石将会消失。
3. 若第 $i$ 秒开始时，Amber 在 $(x,y)$，则在第 $(i+1)$ 秒开始前，Amber 可以马上移动到相邻的格子 $(x+1,y),(x-1,y),(x,y+1),(x,y-1)$ 或原地不动 $(x,y)$。

求 Amber 最多能得到多大总价值的宝石。

![aaa.png](https://cdn.acwing.com/media/article/image/2020/08/14/19_10570ff8dd-aaa.png)

上图给出了一个 $2 \\times 2$ 的网格的例子。

在第 $0$ 秒，首先选择 $B2$ 进入，取走宝石 $3$；由于是偶数秒，周围的格子 $A2,B1$ 的宝石 $1,2$ 消失；向 $A2$ 走去。

在第 $1$ 秒，由于 $A2$ 的宝石已消失，无宝石可取；向 $A1$ 走去。

在第 $2$ 秒，取走 $A1$ 的宝石 $4$。

全程共取得 $2$ 块宝石：宝石 $3$ 和宝石 $4$。

### 输入格式

第一行包含两个整数 $n,m$。

接下来 $n$ 行，每行包含 $m$ 个整数，用来描述宝石价值矩阵。其中第 $i$ 行第 $j$ 列的整数表示 $v\_{i,j}$。

### 输出格式

输出可拿走的宝石最大总价值。

### 数据范围

$1 \\le n,m \\le 100$,

$1 \\le v\_{i,j} \\le 1000$

### 输入样例：

```
2 2
1 2
2 1
```

### 输出样例：

```
4
```

### 题解

