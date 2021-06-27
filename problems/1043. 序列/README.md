## [1043. 序列](https://www.acwing.com/problem/content/1045/)

### 题目

给定两个长度为 *n* 的正整数序列 *lbrace a_i rbrace* 与 *lbrace b_i rbrace*，序列的下标为 *1, 2, ⋯ , n*。

现在你需要分别对两个序列各指定恰好 *K* 个下标，要求至少有 *L* 个下标在两个序列中都被指定，使得这 *2K* 个下标在序列中对应的元素的总和最大。

形式化地说，你需要确定两个长度为 *K* 的序列 *lbrace c_i rbrace, lbrace d_i rbrace*，其中

**1 ≤ c_1 < c_2 < ⋯ < c_K ≤ n , 1 ≤ d_1 < d_2 < ⋯ < d_K ≤ n**

并要求

**|lbrace c_1, c_2, ⋯ , c_K rbrace cap lbrace d_1, d_2, · · · , d_K rbrace| ≥ L**

目标是最大化

**sum^{K}_{i=1} a_{c_i} +sum^{K}_{i=1} b_{d_i}**

### 输入格式

本题输入文件包含多组数据。

第一行一个正整数 *T* 表示数据组数。接下来每三行表示一组数据。

每组数据第一行三个整数 *n, K, L*，变量意义见题目描述。

每组数据第二行 *n* 个整数表示序列 *lbrace a_i rbrace*。

每组数据第三行 *n* 个整数表示序列 *lbrace b_i rbrace*。

### 输出格式

对于每组数据输出一行一个整数表示答案。

### 数据范围

对于所有测试点：*T ≤ 10 , 1 ≤ sum n ≤ 10^6, 1 ≤ L ≤ K ≤ n ≤ 2 × 10^5, 1 ≤ a_i, b_i ≤ 10^9*。

每个测试点的具体限制见下表：

 ![QQ截图20190920174007.png](https://cdn.acwing.com/media/article/image/2019/09/20/19_ac2ff5b4db-QQ截图20190920174007.png)

### 输入样例：

```
5
1 1 1
7
7
3 2 1
4 1 2
1 4 2
5 2 1
4 5 5 8 4
2 1 7 2 7
6 4 1
1 5 8 3 2 4
2 6 9 3 1 7
7 5 4
1 6 6 6 5 9 1
9 5 3 9 1 4 2
```

### 输出样例：

```
14
12
27
45
62
```

### 样例解释

第一组数据选择的下标为：*lbrace c_i rbrace = lbrace 1 rbrace , lbrace d_i rbrace = lbrace 1 rbrace*。

第二组数据选择的下标为：*lbrace c_i rbrace = lbrace 1, 3 rbrace , lbrace d_i rbrace = lbrace 2, 3 rbrace*

第三组数据选择的下标为：*lbrace c_i rbrace = lbrace 3, 4 rbrace , lbrace d_i rbrace = lbrace 3, 5 rbrace*。

第四组数据选择的下标为：*lbrace c_i rbrace = lbrace 2, 3, 4, 6 rbrace , lbrace d_i rbrace = lbrace 2, 3, 4, 6 rbrace*。

第五组数据选择的下标为：*lbrace c_i rbrace = lbrace 2, 3, 4, 5, 6 rbrace , lbrace d_i rbrace = lbrace 1, 2, 3, 4, 6 rbrace*。
