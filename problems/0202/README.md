## [202\. 最幸运的数字](https://www.acwing.com/problem/content/204/)

### 题目

$8$ 是中国的幸运数字，如果一个数字的每一位都由 $8$ 构成则该数字被称作是幸运数字。

现在给定一个正整数 $L$，请问至少多少个 $8$ 连在一起组成的正整数（即最小幸运数字）是 $L$ 的倍数。

### 输入格式

输入包含多组测试用例。

每组测试用例占一行，包含一个整数 $L$。

当输入用例 $L=0$ 时，表示输入终止，该用例无需处理。

### 输出格式

每组测试用例输出结果占一行。

结果为 `Case i:` +一个整数 $N$，$N$ 代表满足条件的最小幸运数字的位数。

如果满足条件的幸运数字不存在，则 $N=0$。

### 数据范围

$1 \\le L \\le 2 \\times 10^9$

### 输入样例：

```
8
11
16
0
```

### 输出样例：

```
Case 1: 1
Case 2: 2
Case 3: 0
```

### 题解
