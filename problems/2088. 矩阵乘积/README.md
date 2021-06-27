## [2088. 矩阵乘积](https://www.acwing.com/problem/content/2090/)

### 题目

已知矩阵：

**A_{mtimes n}=begin{Bmatrix}a_{11}  a_{12}    a_{1n}newline a_{21}  a_{22}    a_{2n} newline newline a_{m1}  a_{m2}    a_{mn} end{Bmatrix} ,B_{ntimes p}=begin{Bmatrix}b_{11}  b_{12}    b_{1p}newline b_{21}  b_{22}    b_{2p} newline newline b_{n1}  b_{n2}    b_{nq} end{Bmatrix}**

当矩阵 *A* 的列数与矩阵 *B* 的行数相同时，则 *A* 与 *B* 可以相乘，其乘积为一个 *mtimes p* 的矩阵 *D*：

**D_{mtimes p}=begin{Bmatrix} d_{11}  d_{12}    d_{1p}newline d_{21}  d_{22}    d_{2p} newline newline d_{m1}  d_{m2}    d_{mp}end{Bmatrix}**

其中 *d_{ij}=sumlimits^n_{k=1} a_{ik} × b_{kj}*，简记为 *D=Atimes B*。

现已知三个矩阵 *A,B,C*，这三个矩阵大多数元素为 *0*，我们把这种矩阵称为稀疏矩阵。因此，我们采用三元组 *i,j,a* 来表示矩阵的第 *i* 行第 *j* 列的值为 *a* 其余未列出的元素均为 *0*；在计算机中，我们仅给出非零元素的三元组，而且使用行优先法给出稀疏矩阵的三元组，首先是第 *1* 行按列给出，然后是第 *2* 行按列给出……

例如，矩阵：*begin{Bmatrix}1   0   0   0newline0   0   2 -1newline0   1   2   3newline0   0   0   0end{Bmatrix}* 那么，矩阵的三元组表示为：

```
1 1 1
2 3 2
2 4 -1
3 2 1
3 3 2
3 4 3
```

你的任务就是：编程完成计算 *D=Atimes Btimes C*。

### 输入格式

第 *1* 行为两个正整数：*x,y*，分别表示输出结果所在的行和列。

第 *2* 行为四个正整数 *m,n,o,p*，表明 *A* 为 *m × n* 矩阵，*B* 为 *n × o* 矩阵，*C* 为 *o × p* 矩阵。

第 *3* 行以后的每一行有三个整数分别是矩阵的三元组表示法中的一个元素的值，每个矩阵之间有一个空行。表示的顺序是矩阵 *A、B* 和 *C*。

### 输出格式

输出 *D=Atimes Btimes C* 的第 *x* 行第 *y* 列元素的值。

### 数据范围

*1 ≤ m,n,o,p ≤ 6000*,

三元数组的总个数不大于 *6000*。

### 输入样例：

```
1 2
3 4 2 3
1 1 3
1 4 5
2 2 1
3 1 2

1 2 2
2 1 1
3 1 2
3 2 4

1 2 2
1 3 3
2 1 1
2 2 2
```

### 输出样例：

```
12
```
