## [638. AlbocedeDNA](https://www.acwing.com/problem/content/640/)

### 题目

Albocede 外来物种的 DNA 由 *4* 种类型的核苷酸组成：*a，b，c* 和 *d*。

不同的 Albocede 物种可能具有这些核苷酸的不同序列，但任何 Albocede 物种的 DNA 序列都遵循以下所有规则：

- 序列中 *a，b，c* 和 *d*，每种核苷酸都至少包含一个。
- 序列中所有 *a* 都在 *b* 之前，所有 *b* 都在 *c* 之前，所有 *c* 都在 *d* 之前。
- 序列中 *a* 和 *c* 的数量相同。
- 序列中 *b* 和 *d* 的数量相同。

例如， `abcd` 和 `aabbbccddd` 是有效的 Albocede DNA 序列，而 `acbd`， `abc` 和 `abbccd` 不是。

Albocede-n 是 Albocede的进化物种。

Albocede-n 的 DNA 序列由一个或多个有效的 Albocede DNA 序列端对端连接在一起组成。

例如， `abcd` 和 `aaabcccdaabbbccdddabcd` 都是有效的 Albocede-n DNA 序列。 （可以发现有效的 Albocede-n DNA 序列不一定也是有效的 Albocede DNA 序列。）

在你的一次外星探险中，你找到了一个有趣的 DNA 序列，它只由 *a，b，c* 和 *d* 组成。

你感兴趣的是该序列中有多少不同的子序列是有效的 Albocede-n DNA 序列。

由于结果可能非常大，请将结果对 *1000000007* 取模后输出。

### 输入格式

第一行包含整数 *T*，表示共有 *T* 组测试数据。

每组数据占一行，包含一个由 *a,b,c,d* 构成的字符串 *S*。

### 输出格式

每组数据输出一个结果，每个结果占一行。

结果表示为 `Case #x: y`，其中 *x* 是组别编号（从 *1* 开始），*y* 是对 *1000000007* 取模后的结果。

### 数据范围

*1 ≤ T ≤ 20*,

*1 ≤ S的长度 ≤ 500*

### 输入样例：

```
5
abcd
aaaabcd
aaaabbccd
abcdabcdaabccd
b
```

### 输出样例：

```
Case #1: 1
Case #2: 4
Case #3: 28
Case #4: 71
Case #5: 0
```
