## [2566. 立体图](https://www.acwing.com/problem/content/2568/)

### 题目

小渊是个聪明的孩子，他经常会给周围的小朋友们讲些自己认为有趣的内容。

最近，他准备给小朋友们讲解彩色水平光源照射下的立体图，并请你帮他在平面上画出来。

小渊有一块面积为 *m × n* 的矩形区域，上面有 *m × n* 个边长为 *1* 的格子，每个格子上堆了一些同样大小的积木（积木的长宽高都是 *1*）。

为了方便阐述，我们假设这块区域是坐北朝南的，下面我们给出一个例子。

 ![1.png](https://cdn.acwing.com/media/article/image/2020/10/09/19_5e397b4c09-1.png)

小渊想请你打印出这些格子的立体图。

我们定义每个积木为如下格式，并且不会做任何翻转旋转，只会严格以这一种形式摆放（左侧是应该打印出来的图样，右侧为对应每一个位置符号的十进制 ASCII，其中 ASCII 为 *32* 的符号为空）。

 ![2.png](https://cdn.acwing.com/media/article/image/2020/10/09/19_d5b6570809-2.png)

在良好的光学环境下，小渊将 *T* 束平行光同时照射在这些积木上。

这些平行光首先满足一定是红绿蓝三基色之一，其次入射角度满足：与 *x* 轴和 *y* 轴的夹角度数均为 *45* 的倍数；且与 *z* 轴正方向的夹角或为 *45* 度，或为 *0* 度，或为 *315* 度。

具体来说，我们最多会考虑 *9* 个方向的不同平行光，它们的入射方向可以被描述为：

```
西北方45度仰角  正北方45度仰角  东北方45度仰角

正西方45度仰角  垂直从上入射光  正东方45度仰角

西南方45度仰角  正南方45度仰角  东南方45度仰角
```

对于每一个单位积木来说，可以打印出来的三个表面被分为 *12* 个小三角形。

红绿蓝三基色分别用字母 RGB 来表示。

而二次叠加后的三种颜色青黄紫，分别用 YCP 来表示。

对于三次叠加后的颜色，也就是白色，用 W 来表示。

 ![3.png](https://cdn.acwing.com/media/article/image/2020/10/09/19_c48703d609-3.png)

### 输入格式

第一行有两个正整数 *m* 和 *n*，表示区域有 *m* 行 *n* 列。

之后 *m* 行，依次由远及近描述了每一行的情况。每一行给出 *n* 个正整数，表示第 *i* 行第 *j* 列中有堆放了多少积木。

之后 *3* 行，每行三个字符，描述了 *9* 个对应方向（与地图描述方向相同）的光照颜色。其中每一个字符或者为 RGB 中之一，表示对应的颜色。或者为 `*`，表示没有照射光。

这九个方向依次是：

```
西北方45度仰角  正北方45度仰角  东北方45度仰角
正西方45度仰角  垂直从上入射光  正东方45度仰角
西南方45度仰角  正南方45度仰角  东南方45度仰角
```

### 输出格式

输出给出了打印后的效果。其中要求输出结果不含前导空行，结尾也没有额外空行。

输出的第一列不能全是空格，且每一行末尾也没有额外空格。

### 数据范围

*1 ≤ n,m ≤ 100*,

每一个位置堆放的积木总数不超过 *100*,

入射光颜色可能是 RGB 中的任何一种颜色,

最多可以有 *9* 束入射光。

### 输入样例1：

```
2 2
2 1
1 1
R**
***
**G
```

### 输出样例1：

```
        +-------+
       /YYYYY’/|
      /YY.*’YY/G|
     /.YYYYY/G/|
    +-------+G.G|
    |GGGGG/|:G|
    |GGGG/G|G*G|
    |GGG/GG|G:|
    |GGGXGGG|G’G+-------+
    |GG/GGG|/G/GYYYY’/|
    |G/GGGG|G/GG.*’YY/G|
    |/GGGGG|/.GGGGY/G/|
    +-------+-------+G.G|
/YGGGG’/GGGGG’/|:G|
/YY.*’GG/GG.*’GG/G|G*G|
/.YYYYG/.GGGGG/G/|G:|
+-------+-------+G.G|G’G+
|GGGGG/|GGGGG/|:G|/G/
|GGGG/G|GGGG/G|G*G|G/
|GGG/GG|GGG/GG|G:|/
|GGGXGGG|GGGXGGG|G’G+
|GG/GGG|GG/GGG|/G/
|G/GGGG|G/GGGG|G/
|/GGGGG|/GGGGG|/
+-------+-------+
```

### 输入样例2：

```
3 4
1 1 2 1
1 2 1 2
2 1 2 1
**B
***
R*G
```

### 输出样例2：

```
                    +-------+
                   /WWWWW’/|
                  /WW.*’WW/C|
                 /.WWWWW/C/|
        +-------+-------+-------+
       /WWWWW’/|YYYYY/WWWWW’/|
      /WW.*’WW/C|GYYY/WW.*’WW/C|
     /.WWWWW/C/|GGY/.WWWWW/C/|
    +-------+-------+-------+-------+C.C|---+
/WWWWW’/|YYYYY/WWWWW’/|YYYYY/|:C|C’/|
/WW.*’WW/C|GYYY/WW.*’WW/C|GYYY/Y|C*C|C/C|
/.WWWWW/C/|GGY/.WWWWW/C/|GGY/YY|C:|/C/|
+-------+C.G|GGG+-------+C.G|GGGXYYY|C’C+C.C|
|YYYYY/|:G|GG/|YYYYY/|:G|GG/GYY|/C/|:C|
|YYYY/Y|C*G|G/K|YYYY/Y|C*G|G/GGGY|C/C|C*C|
|YYY/YY|C:|/KK|YYY/YY|C:|/GGGGG|/C/|C:|
|YYYXYYY|C’G+---|YYYXYYY|C'G+-------+C.C|C’C+
|YY/YYY|/G/GKK|YY/YYY|/G/GGGGG’/|:C|/C/
|Y/YYYY|G/GG.*’|Y/YYYY|G/GG.*’WW/C|C*C|C/
|/YYYYY|/.YYYY|/YYYYY|/.WWWWW/C/|C:|/
+-------+-------+-------+-------+C.C|C’C+
|YYYYY/|YYYYY/|YYYYY/|YYYYY/|:C|/C/
|YYYY/Y|YYYY/Y|YYYY/Y|YYYY/Y|C*C|C/
|YYY/YY|YYY/YY|YYY/YY|YYY/YY|C:|/
|YYYXYYY|YYYXYYY|YYYXYYY|YYYXYYY|C’C+
|YY/YYY|YY/YYY|YY/YYY|YY/YYY|/C/
|Y/YYYY|Y/YYYY|Y/YYYY|Y/YYYY|C/
|/YYYYY|/YYYYY|/YYYYY|/YYYYY|/
+-------+-------+-------+-------+
```

### 样例2提示

 ![5.png](https://cdn.acwing.com/media/article/image/2020/10/09/19_0b49404e09-5.png)
