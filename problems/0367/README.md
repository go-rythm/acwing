## [367\. 学校网络](https://www.acwing.com/problem/content/369/)

### 题目

一些学校连接在一个计算机网络上，学校之间存在软件支援协议，每个学校都有它应支援的学校名单（学校 $A$ 支援学校 $B$，并不表示学校 $B$ 一定要支援学校 $A$）。

当某校获得一个新软件时，无论是直接获得还是通过网络获得，该校都应立即将这个软件通过网络传送给它应支援的学校。

因此，一个新软件若想让所有学校都能使用，只需将其提供给一些学校即可。

现在请问最少需要将一个新软件直接提供给多少个学校，才能使软件能够通过网络被传送到所有学校？

最少需要添加几条新的支援关系，使得将一个新软件提供给任何一个学校，其他所有学校就都可以通过网络获得该软件？

### 输入格式

第 $1$ 行包含整数 $N$，表示学校数量。

第 $2..N+1$ 行，每行包含一个或多个整数，第 $i+1$ 行表示学校 $i$ 应该支援的学校名单，每行最后都有一个 $0$ 表示名单结束（只有一个 $0$ 即表示该学校没有需要支援的学校）。

### 输出格式

输出两个问题的结果，每个结果占一行。

### 数据范围

$2 \\le N \\le 100$

### 输入样例：

```
5
2 4 3 0
4 5 0
0
0
1 0
```

### 输出样例：

```
1
2
```

### 题解
