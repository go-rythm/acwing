## [3370. 牛年](https://www.acwing.com/problem/content/3373/)

### 题目

Farmer John 的奶牛们得知最近正在庆祝牛年的到来时十分兴奋。

牛年总是奶牛们的最爱。

我们知道，中国历法中每一年所对应的生肖遵循 *12* 年的周期： `Ox, Tiger, Rabbit, Dragon, Snake, Horse, Goat, Monkey, Rooster, Dog, Pig, Rat`（牛、虎、兔、龙、蛇、马、羊、猴、鸡、狗、猪、鼠），然后回到牛。

奶牛 Bessie 自豪地说她是在许多年前的一个牛年出生的。

她的朋友 Elsie 想要知道她与 Bessie 出生相差多少年，并且希望你能够通过查看农场上若干奶牛出生年份之间的关系来帮助她推算。

### 输入格式

输入的第一行包含一个整数 *N*。

以下 *N* 行每行包含一个 *8* 个单词的短语，指定了两头奶牛的出生年份之间的关系，格式为 `Mildred born in previous Dragon year from Bessie`（Mildred 在 Bessie 出生的前一个龙年出生），或 `Mildred born in next Dragon year from Bessie`（Mildred 在 Bessie 出生的后一个龙年出生）。

最后一个单词是农场上某一头奶牛的名字，为 “Bessie” 或一头已经在之前的输入中出现过的奶牛。

第一个单词是农场上某一头奶牛的名字，不为 “Bessie” 且未在之前的输入中出现过。

所有的奶牛名字不超过 *10* 个字符，且仅包含字符 *a..z* 或 *A..Z*。

第 *5* 个单词是上述十二生肖之一。

第 *4* 个单词是 `previous`（之前）或 `next`（之后）之一。

例如，如果短语为 `Mildred born in previous Dragon year from Bessie`，则 Mildred 的出生年份为最为接近且严格处于 Bessie 的出生年份之前（不等于）的龙年。

### 输出格式

输出 Bessie 和 Elsie 的出生年份之间相差的年数。输入保证可以通过给定的信息求出结果。

### 数据范围

*1 ≤ N ≤ 100*

### 输入样例：

```
4
Mildred born in previous Dragon year from Bessie
Gretta born in previous Monkey year from Mildred
Elsie born in next Ox year from Gretta
Paulina born in next Dog year from Bessie
```

### 输出样例：

```
12
```

### 样例解释

在以上的输入中，

- Elsie 在 Bessie 之前 *12* 年出生。
- Mildred 在 Bessie 之前 *9* 年出生。
- Gretta 在 Bessie 之前 *17* 年出生。
- Paulina 在 Bessie 之前 *9* 年出生。
