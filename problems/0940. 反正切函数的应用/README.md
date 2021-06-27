## [940. 反正切函数的应用](https://www.acwing.com/problem/content/942/)

### 题目

反正切函数可展开成无穷级数，有如下公式

 ![QQ截图20190901201054.png](https://cdn.acwing.com/media/article/image/2019/09/01/19_94ae70eccc-QQ截图20190901201054.png) 公式 *(1)*

使用反正切函数计算 *π* 是一种常用的方法。例如，最简单的计算 *π* 的方法：

 ![QQ截图20190901201146.png](https://cdn.acwing.com/media/article/image/2019/09/01/19_b5508632cc-QQ截图20190901201146.png) 公式 *(2)*

然而，这种方法的效率很低，但我们可以根据角度和的正切函数公式：

 ![QQ截图20190901201223.png](https://cdn.acwing.com/media/article/image/2019/09/01/19_c83cfd7acc-QQ截图20190901201223.png) 公式 *(3)*

通过简单的变换得到：

 ![QQ截图20190901201246.png](https://cdn.acwing.com/media/article/image/2019/09/01/19_d64e3910cc-QQ截图20190901201246.png) 公式 *(4)*

利用这个公式，令 *p = frac{1}{2}, q = frac{1}{3}*，则 *frac{p+q}{1-pq} = 1*，有

 ![QQ截图20190901201558.png](https://cdn.acwing.com/media/article/image/2019/09/01/19_473c73dacc-QQ截图20190901201558.png)

使用 *frac{1}{2}* 和 *frac{1}{3}* 的反正切来计算 *arctan(1)*，速度就快多了。

我们将公式 *(4)* 写成如下形式

 ![QQ截图20190901201723.png](https://cdn.acwing.com/media/article/image/2019/09/01/19_7b6ca904cc-QQ截图20190901201723.png)

其中 *a、b* 和 *c* 均为正整数。

我们的问题是：对于每一个给定的 *a*，求 *b+c* 的值。

我们保证对于任意的 *a* 都存在整数解。

如果有多个解，要求你给出 *b+c* 最小的解。

### 输入格式

只包含一个整数 *a*。

### 输出格式

只包含一个整数，为 *b+c* 的值。

### 数据范围

*1 ≤ a ≤ 60000*

### 输入样例：

```
1
```

### 输出样例：

```
5
```
