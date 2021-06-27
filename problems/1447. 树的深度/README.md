## [1447. 树的深度](https://www.acwing.com/problem/content/1449/)

### 题目

对于新的一年，农夫约翰打算给他的奶牛们一个节日二叉搜索树(BST)。

为了生成这个二叉搜索树，约翰首先选定一个 *1 ~ N* 的排列 *a=lbrace a_1,a_2,…,a_N rbrace*，然后将 *1* 和 *N* 作为参数，运行以下伪代码：

```
generate(l,r):
if l > r, return empty subtree;
x = argmin_{l <= i <= r} a_i; // index of min a_i in {a_l,...,a_r}
return a BST with x as the root,
    generate(l,x-1) as the left subtree,
    generate(x+1,r) as the right subtree;
```

例如，排列 *lbrace 3,2,5,1,4 rbrace* 生成以下BST：

```
    4
/ 
2   5
/ 
1   3
```

用 *d_i(a)* 表示排列 *a* 生成的二叉搜索树中，节点 *i* 的深度。

一个节点的深度定义为该节点到根节点的路径上的节点数。

在以上示例中，*d_4(a)= 1,d_2(a)= d_5(a)= 2,d_1(a)= d_3(a)= 3*。

排列 *a* 中，逆序对的数量等于满足 *1 ≤ i < j ≤ N* 且 *a_i > a_j* 的数对 *(i,j)* 的个数。

已知用来生成二叉搜索树的排列 *a* 中，共有 *K* 个逆序对。

对于每个整数 *i*(*1 ≤ i ≤ N*)，计算 *i* 节点在所有满足条件的排列所生成的二叉搜索树中的深度之和对 *M* 取模后的结果，也就是：

**sum_ad_i(a)  (bmod M)**

### 输入格式

共一行，包含三个整数 *N,K,M*。

### 输出格式

输出一行 *N* 个整数，其中第 *i* 个数表示**sum_ad_i(a) (bmod M)**的结果。

### 数据范围

*1 ≤ N ≤ 300*,

*0 ≤ K ≤ frac{N(N-1)}{2}*,

*M* 是一个质数，范围在 *[10^8,10^9+9]*。

### 输入样例1：

```
3 0 192603497
```

### 输出样例1：

```
1 2 3
```

### 样例1解释

满足条件的排列只有一种，*a = lbrace 1,2,3 rbrace*。

### 输入样例2：

```
3 1 144408983
```

### 输出样例2：

```
3 4 4
```

### 样例2解释

满足条件的排列有两种，*a = lbrace 1,3,2 rbrace* 和 *a = lbrace 2,1,3 rbrace*。
