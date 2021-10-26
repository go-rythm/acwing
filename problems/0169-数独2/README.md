## [169\. 数独2](https://www.acwing.com/problem/content/171/)

### 题目

请你将一个 $16 \\times 16$ 的数独填写完整，使得每行、每列、每个 $4 \\times 4$ 十六宫格内字母 $A \\sim P$ 均恰好出现一次。

保证每个输入只有唯一解决方案。

![数独2.jpg](/media/article/image/2019/01/16/19_cabce58018-数独2.jpg)

### 输入格式

输入包含多组测试用例。

每组测试用例包括 $16$ 行，每行一组字符串，共 $16$ 个字符串。

第 $i$ 个字符串表示数独的第 $i$ 行。

字符串包含字符可能为字母 $A \\sim P$ 或 `-`（表示等待填充）。

测试用例之间用单个空行分隔，输入至文件结尾处终止。

### 输出格式

对于每个测试用例，均要求保持与输入相同的格式，将填充完成后的数独输出。

每个测试用例输出结束后，输出一个空行。

### 输入样例：

```
--A----C-----O-I
-J--A-B-P-CGF-H-
--D--F-I-E----P-
-G-EL-H----M-J--
----E----C--G---
-I--K-GA-B---E-J
D-GP--J-F----A--
-E---C-B--DP--O-
E--F-M--D--L-K-A
-C--------O-I-L-
H-P-C--F-A--B---
---G-OD---J----H
K---J----H-A-P-L
--B--P--E--K--A-
-H--B--K--FI-C--
--F---C--D--H-N-
```

### 输出样例：

```
FPAHMJECNLBDKOGI
OJMIANBDPKCGFLHE
LNDKGFOIJEAHMBPC
BGCELKHPOFIMAJDN
MFHBELPOACKJGNID
CILNKDGAHBMOPEFJ
DOGPIHJMFNLECAKB
JEKAFCNBGIDPLHOM
EBOFPMIJDGHLNKCA
NCJDHBAEKMOFIGLP
HMPLCGKFIAENBDJO
AKIGNODLBPJCEFMH
KDEMJIFNCHGAOPBL
GLBCDPMHEONKJIAF
PHNOBALKMJFIDCEG
IAFJOECGLDPBHMNK
```

### 题解

