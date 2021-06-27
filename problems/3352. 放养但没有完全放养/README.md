## [3352. 放养但没有完全放养](https://www.acwing.com/problem/content/3355/)

### 题目

一个鲜为人知的事实是，奶牛拥有自己的文字：「牛文」。

牛文由 *26* 个字母 `a` 到 `z` 组成，但是当奶牛说牛文时，可能与我们所熟悉的 `abcdefghijklmnopqrstuvwxyz` 不同，她会按某种特定的顺序排列字母。

为了打发时间，Bessie 的表妹 Mildred 在反复哼唱牛文字母歌，而 Farmer Nhoj 好奇她唱了多少遍。

给定一个小写字母组成的字符串，为 Farmer Nhoj 听到 Mildred 唱的字母，计算 Mildred 至少唱了几遍完整的牛文字母歌，使得 Farmer Nhoj 能够听到给定的字符串。

Farmer Nhoj 并不始终注意 Mildred 所唱的内容，所以他可能会漏听 Mildred 唱过的一些字母。

给定的字符串仅包含他记得他所听到的字母。

### 输入格式

输入仅一行，包含一个小写字母组成的字符串，为 Farmer Nhoj 听到 Mildred 唱的字母。

### 输出格式

输出 Mildred 所唱的完整的牛文字母歌的最小次数。

### 数据范围

字符串的长度不小于 *1* 且不大于 *10^5*。

### 输入样例：

```
mildredree
```

### 输出样例：

```
3
```

### 样例解释

Mildred 至少唱了三遍牛文字母歌。

有可能 Mildred 只唱了三遍牛文字母歌，如果牛文字母表以 `mildre` 开头，并且 Farmer Nhoj 听到了以下被标记为大写的字母。

```
MILDREabcfghjknopqstuvwxyz
milDREabcfghjknopqstuvwxyz
mildrEabcfghjknopqstuvwxyz
```
