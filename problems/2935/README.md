## [2935\. 信用卡凸包](https://www.acwing.com/problem/content/2938/)

### 题目

信用卡是一个矩形，唯四个角作了圆滑处理，使它们都是与矩形的两边相切的 $\\frac{1}{4}$ 圆，如下图所示。

![QQ截图20201204094116.png](https://cdn.acwing.com/media/article/image/2020/12/04/19_d45f3cd235-QQ截图20201204094116.png)

现在平面上有一些规格相同的信用卡，试求其凸包的周长。

注意凸包未必是多边形，因为它可能包含若干段圆弧。

### 输入格式

第一行是一个正整数 $n$，表示信用卡的张数。

第二行包含三个实数 $a, b, r$，分别表示信用卡（圆滑处理前）竖直方向的长度、水平方向的长度，以及 $\\frac{1}{4}$ 圆的半径。

之后 $n$ 行，每行包含三个实数 $x, y, θ$，分别表示一张信用卡中心（即对角线交点）的横、纵坐标，以及绕中心 **逆时针** 旋转的 **弧度**。

### 输出格式

输出只有一行，包含一个实数，表示凸包的周长，四舍五入精确到小数点后 $2$ 位。

### 数据范围

$1 \\le n \\le 10000$,

$0.1 \\le a,b \\le 1000000.0$,

$0.0 \\le r < \\min \\{ a/4,b/4 \\}$,

$\|x\|,\|y\| \\le 1000000.0$,

$0 \\le θ < 2 \\pi$

### 输入样例1：

```
2
6.0 2.0 0.0
0.0 0.0 0.0
2.0 -2.0 1.5707963268
```

### 输出样例1：

```
21.66
```

### 样例1解释

![1.jpg](https://cdn.acwing.com/media/article/image/2020/12/04/19_f97ba0f435-1.jpg)

本样例中的 $2$ 张信用卡的轮廓在上图中用实线标出，如果视 $1.5707963268$ 为 $\\pi/2$( $\\pi$ 为圆周率)，则其凸包的周长为 $16+4 \\times \\sqrt 2$。

### 输入样例2：

```
3
6.0 6.0 1.0
4.0 4.0 0.0
0.0 8.0 0.0
0.0 0.0 0.0
```

### 输出样例2：

```
41.60
```

### 样例2解释

![2.png](https://cdn.acwing.com/media/article/image/2020/12/04/19_2f0a3d8635-2.png)

### 输入样例3：

```
3
6.0 6.0 1.0
4.0 4.0 0.1745329252
0.0 8.0 0.3490658504
0.0 0.0 0.5235987756
```

### 输出样例3：

```
41.63
```

### 样例3解释

![3.png](https://cdn.acwing.com/media/article/image/2020/12/04/19_6557c7a035-3.png)

其凸包的周长约为 $41.628267652$。

### 题解
