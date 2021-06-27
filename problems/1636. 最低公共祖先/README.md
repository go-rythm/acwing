## [1636. 最低公共祖先](https://www.acwing.com/problem/content/1638/)

### 题目

树中两个结点 *U* 和 *V* 的最低公共祖先（LCA）是指同时具有 *U* 和 *V* 作为后代的最深结点。

二叉搜索树 (BST) 递归定义为具有以下属性的二叉树：

- 若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值
- 若它的右子树不空，则右子树上所有结点的值均大于或等于它的根结点的值
- 它的左、右子树也分别为二叉搜索树

现在给定 BST 中的任意两个结点，请你找出它们的最低公共祖先。

### 输入格式

第一行包含两个整数 *M* 和 *N*，分别表示询问结点对数以及二叉搜索树中的结点数量。

第二行包含 *N* 个不同整数，表示该二叉搜索树的前序遍历序列。

接下来 *M* 行，每行包含两个整数 *U* 和 *V*，表示一组询问。

所有结点权值均在 **int** 范围内。

### 输出格式

对于每对给定的 *U* 和 *V*，输出一行结果。

如果 *U* 和 *V* 的 LCA 是 *A*，且 *A* 不是 *U* 或 *V*，则输出 `LCA of U and V is A.`。

如果 *U* 和 *V* 的 LCA 是 *A*，且 *A* 是 *U* 或 *V* 中的一个，则输出 `X is an ancestor of Y.`，其中 *X* 表示 *A*，*Y* 表示另一个结点。

如果 *U* 或 *V* 没有在 BST 中找到，则输出 `ERROR: U is not found.` 或 `ERROR: V is not found.` 或 `ERROR: U and V are not found.`。

### 数据范围

*1 ≤ M ≤ 1000*,

*1 ≤ N ≤ 10000*

### 输入样例：

```
6 8
6 3 1 2 5 4 8 7
2 5
8 7
1 9
12 -3
0 8
99 99
```

### 输出样例：

```
LCA of 2 and 5 is 3.
8 is an ancestor of 7.
ERROR: 9 is not found.
ERROR: 12 and -3 are not found.
ERROR: 0 is not found.
ERROR: 99 and 99 are not found.
```
