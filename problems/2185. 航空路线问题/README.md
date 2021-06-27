## [2185. 航空路线问题](https://www.acwing.com/problem/content/2187/)

### 题目

给定一张航空图，图中顶点代表城市，边代表 *2* 城市间的直通航线。

现要求找出一条满足下述限制条件的且途经城市最多的旅行路线。

1. 从 **最西端** 城市出发，单向从西向东途经若干城市到达 **最东端** 城市，然后再单向从东向西飞回起点(可途经若干城市)。
2. 除起点城市外，任何城市只能访问 *1* 次。

对于给定的航空图，试设计一个算法找出一条满足要求的最佳航空旅行路线。

### 输入格式

第 *1* 行有 *2* 个正整数 *N* 和 *V*，*N* 表示城市数，*V* 表示直飞航线数。

接下来的 *N* 行中每一行是一个城市名，可乘飞机访问这些城市。

城市名出现的顺序是从西向东。也就是说，设 *i,j* 是城市表列中城市出现的顺序，当 *i>j* 时，表示城市 *i* 在城市 *j* 的东边，而且不会有 *2* 个城市在同一条经线上。城市名是一个长度不超过 *15* 的字符串，串中的字符可以是字母或阿拉伯数字。例如，*AGR34* 或 *BEL4*。

再接下来的 *V* 行中，每行有 *2* 个城市名，中间用空格隔开，如 *city1* *city2* 表示 *city1* 到 *city2* 有一条直通航线，从 *city2* 到 *city1* 也有一条直通航线。

### 输出格式

第 *1* 行是旅行路线中所访问的城市总数 *M*。

接下来的 *M＋1* 行是旅行路线的城市名，每行写 *1* 个城市名。

首先是出发城市名，然后按访问顺序列出其它城市名。

注意，最后 *1* 行（终点城市）的城市名必然是出发城市名。

如果问题有多个解，则输出任意一个即可。

如果问题无解，则输出 `No Solution!`。

### 数据范围

*1 ≤ N ≤ 30*,

*1 ≤ V ≤ 250*

### 输入样例：

```
8 9
Vancouver
Yellowknife
Edmonton
Calgary
Winnipeg
Toronto
Montreal
Halifax
Vancouver Edmonton
Vancouver Calgary
Calgary Winnipeg
Winnipeg Toronto
Toronto Halifax
Montreal Halifax
Edmonton Montreal
Edmonton Yellowknife
Edmonton Calgary
```

### 输出样例：

```
7
Vancouver
Edmonton
Montreal
Halifax
Toronto
Winnipeg
Calgary
Vancouver
```
