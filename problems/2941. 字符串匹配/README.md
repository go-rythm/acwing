## [2941. 字符串匹配](https://www.acwing.com/problem/content/2944/)

### 题目

小 *C* 学习完了字符串匹配的相关内容，现在他正在做一道习题。

对于一个字符串 *S*，题目要求他找到 *S* 的所有具有下列形式的拆分方案数：

*S = ABC*，*S = ABABC*，*S = ABAB ⋯ ABC*，其中 *A*，*B*，*C* 均是非空字符串，且 *A* 中出现奇数次的字符数量不超过 *C* 中出现奇数次的字符数量。

更具体地，我们可以定义 *AB* 表示两个字符串 *A*，*B* 相连接，例如 *A = texttt{aab}*，*B = texttt{ab}*，则 *AB = texttt{aabab}*。

并递归地定义 *A^1=A*，*A^n = A^{n - 1} A*（*n ≥ 2* 且为正整数），例如 *A = texttt{abb}*，则 *A^3=texttt{abbabbabb}*。

则小 *C* 的习题是求 *S = {(AB)}^iC* 的方案数，其中 *F(A) ≤ F(C)*，*F(S)* 表示字符串 *S* 中出现奇数次的字符的数量。

两种方案不同当且仅当拆分出的 *A*、*B*、*C* 中有至少一个字符串不同。

小 C 并不会做这道题，只好向你求助，请你帮帮他。

### 输入格式

本题有多组数据，输入文件第一行一个正整数 *T* 表示数据组数。

每组数据仅一行一个字符串 *S*，意义见题目描述。*S* 仅由英文小写字母构成。

### 输出格式

对于每组数据输出一行一个整数表示答案。

### 数据范围

测试点编号*| S | ≤*特殊性质*1 ~ 4**10*无*5 ~ 8**100*无*9 ~ 12**1000*无*13 ~ 14**2^{15}**S* 中只包含一种字符*15 ~ 17**2^{16}**S* 中只包含两种字符*18 ~ 21**2^{17}*无*22 ~ 25**2^{20}*无

对于所有测试点，保证 *1 ≤ T ≤ 5*，*1 ≤ |S| ≤ 2^{20}*。

### 输入样例1：

```
3
nnrnnr
zzzaab
mmlmmlo
```

### 输出样例1：

```
8
9
16
```

### 样例1解释

对于第一组数据，所有的方案为

1. *A=texttt{n}*，*B=texttt{nr}*，*C=texttt{nnr}*。
2. *A=texttt{n}*，*B=texttt{nrn}*，*C=texttt{nr}*。
3. *A=texttt{n}*，*B=texttt{nrnn}*，*C=texttt{r}*。
4. *A=texttt{nn}*，*B=texttt{r}*，*C=texttt{nnr}*。
5. *A=texttt{nn}*，*B=texttt{rn}*，*C=texttt{nr}*。
6. *A=texttt{nn}*，*B=texttt{rnn}*，*C=texttt{r}*。
7. *A=texttt{nnr}*，*B=texttt{n}*，*C=texttt{nr}*。
8. *A=texttt{nnr}*，*B=texttt{nn}*，*C=texttt{r}*。

### 输入样例2：

```
5
kkkkkkkkkkkkkkkkkkkk
lllllllllllllrrlllrr
cccccccccccccxcxxxcc
ccccccccccccccaababa
ggggggggggggggbaabab
```

### 输出样例2：

```
156
138
138
147
194
```