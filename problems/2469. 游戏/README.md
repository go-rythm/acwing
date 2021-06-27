## [2469. 游戏](https://www.acwing.com/problem/content/2471/)

### 题目

小木木和小凳子是两个聪明的孩子，他们五岁的时候就开始学习英语了。

英语老师教他们玩一个很简单的游戏。

老师给他们一张全小写并无特殊符号的英语单词表，单词表如下：

```
ab
arc
arco
bar
bran
carbon
carbons
cobra
crab
crayon
narc
```

然后让他们从单词表里找词语接龙。

接龙的规则如下:

1. 前一个单词拥有的所有字母，在后一个单词里必须出现，而且字母出现次数不少于前一单词。
2. 后一个单词的长度比前一个单词的长度恰好多 *1*。

对于以上例子，一合法的接龙为:

```
ab
bar
crab
cobra
carbon
carbons
```

他们之中，谁接龙的长度长，谁就赢了。

小木木肯定不想输，所以找到你来帮忙。

### 输入格式

共 *n* 行，每行一个长度不超过 *100* 的单词。

### 输出格式

第一行输出接龙的最大长度 *ans*。

接下来 *ans* 行，按顺序每行输出一个单词表示你的接龙方案。

输出任意一种合法方案即可。

### 数据范围

*1 ≤ n ≤ 10000*

### 输入样例：

```
ab
arc
arco
bar
bran
carbon
carbons
cobra
crab
crayon
narc
```

### 输出样例：

```
6
ab
bar
crab
cobra
carbon
carbons
```
