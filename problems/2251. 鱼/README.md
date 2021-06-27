## [2251. 鱼](https://www.acwing.com/problem/content/2253/)

### 题目

在平面坐标系上给定 *n* 个不同的整点（也即横坐标与纵坐标皆为整数的点）。

我们称从这 *n* 个点中选择 *6* 个不同的点所组成的有序六元组 *langle A,B,C,D,E,F rangle* 是一条“鱼”，当且仅当：*AB=AC,BD=CD,DE=DF*（身形要对称），并且 *angle BAD,angle BDA* 与 *angle CAD,angle CDA* 都是锐角（脑袋和屁股显然不能是凹的），*angle ADE,angle ADF* 大于*90^circ*（也即为钝角或平角，为了使尾巴不至于翘那么别扭）。

下图就是一个合法的鱼的例子：

 ![a.png](https://cdn.acwing.com/media/article/image/2020/08/04/19_a0883a30d5-a.png)

其中点的组成相同，但顺序不同的鱼视为不同的鱼，即 *langle A,B,C,D,E,F rangle* 和 *langle A,C,B,D,E,F rangle* 视为不同的两条鱼（毕竟鱼也有背和肚子的两面），同理*langle A,B,C,D,E,F rangle* 和 *langle A,B,C,D,F,E rangle* 也可以视为不同的两条鱼（假设鱼尾巴可以打结）。

问给定的 *n* 个点可以构成多少条鱼。

特别的，数据保证 *n* 个点互不重复。

### 输入格式

第一行一个正整数 *n*，代表平面上点的个数。

接下来 *n* 行每行两个整数 *x,y*，代表点的横纵坐标。

### 输出格式

输出一行一个非负整数，代表鱼的个数。

### 数据范围

*6 ≤ n ≤ 1000*,

*0 ≤ |x|,|y| ≤ 10^9*,

### 输入样例：

```
8
-2 0
-1 0
0 1
0 -1
1 0
2 0
3 1
3 -1
```

### 输出样例：

```
16
```
