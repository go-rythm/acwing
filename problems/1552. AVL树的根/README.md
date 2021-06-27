## [1552. AVL树的根](https://www.acwing.com/problem/content/1554/)

### 题目

AVL树是一种自平衡二叉搜索树。

在AVL树中，任何节点的两个子树的高度最多相差 *1* 个。

如果某个时间，某节点的两个子树之间的高度差超过 *1*，则将通过树旋转进行重新平衡以恢复此属性。

图 *1-4* 说明了旋转规则。

 ![31.jpg](https://cdn.acwing.com/media/article/image/2020/03/30/19_3a941ca472-31.jpg) ![32.jpg](https://cdn.acwing.com/media/article/image/2020/03/30/19_3dd0488472-32.jpg)

 ![33.jpg](https://cdn.acwing.com/media/article/image/2020/03/30/19_411c0b5472-33.jpg) ![34.jpg](https://cdn.acwing.com/media/article/image/2020/03/30/19_43ea12f472-34.jpg)

现在，给定插入序列，请你求出 AVL 树的根是多少。

### 输入格式

第一行包含整数 *N*，表示总插入值数量。

第二行包含 *N* 个不同的整数，表示每个插入值。

### 输出格式

输出得到的 AVL 树的根是多少。

### 数据范围

*1 ≤ N ≤ 20*

### 输入样例1：

```
5
88 70 61 96 120
```

### 输出样例1：

```
70
```

### 输入样例2：

```
7
88 70 61 96 120 90 65
```

### 输出样例2：

```
88
```
