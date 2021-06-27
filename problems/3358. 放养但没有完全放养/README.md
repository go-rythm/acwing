## [3358. 放养但没有完全放养](https://www.acwing.com/problem/content/3361/)

### 题目

一个鲜为人知的事实是，奶牛拥有自己的文字：「牛文」。

牛文由 *26* 个字母 `a` 到 `z` 组成，但是当奶牛说牛文时，可能与我们所熟悉的 `abcdefghijklmnopqrstuvwxyz` 不同，她会按某种特定的顺序排列字母。

为了打发时间，奶牛 Bessie 在反复哼唱牛文字母歌，而 Farmer John 好奇她唱了多少遍。

给定一个小写字母组成的字符串，为 Farmer John 听到 Bessie 唱的字母，计算 Bessie 至少唱了几遍完整的牛文字母歌，使得 Farmer John 能够听到给定的字符串。

Farmer John 并不始终注意 Bessie 所唱的内容，所以他可能会漏听 Bessie 唱过的一些字母。

给定的字符串仅包含他记得他所听到的字母。

### 输入格式

输入的第一行包含 *26* 个小写字母 `a` 到 `z` 的牛文字母表顺序。

下一行包含一个小写字母组成的字符串，为 Farmer John 听到 Bessie 唱的字母。

### 输出格式

输出 Bessie 所唱的完整的牛文字母歌的最小次数。

### 数据范围

字符串的长度不小于 *1* 且不大于 *1000*。

### 输入样例：

```
abcdefghijklmnopqrstuvwxyz
mood
```

### 输出样例：

```
3
```

### 样例解释

在这个样例中，牛文字母表与日常的字母表的排列一致。

Bessie 至少唱了三遍牛文字母歌。

有可能 Bessie 只唱了三遍牛文字母歌，而 Farmer John 听到了以下被标记为大写的字母。

```
abcdefghijklMnOpqrstuvwxyz
abcdefghijklmnOpqrstuvwxyz
abcDefghijklmnopqrstuvwxyz
```
