## [1345. 序号命名](https://www.acwing.com/problem/content/1347/)

### 题目

威斯康星州的大型牧场的主人们喜欢用连续的数字编号给奶牛们命名。

但是奶牛们并不喜欢这种看似非常方便的命名方式。

它们希望能够用它们喜欢的名字来互相称呼，而不是像这样“交个朋友吧，4734”。

现在，请你帮助可怜的牧牛人将这些奶牛的编号转换为一个与其编号有所关联的名字。

因为这些奶牛们都配有电话座机，因此请使用如下所示的电话的标准按键映射，来将数字转换为可能的字母（注意，没有 *Q* 和 *Z*）：

```
2: A,B,C     5: J,K,L    8: T,U,V
3: D,E,F     6: M,N,O    9: W,X,Y
4: G,H,I     7: P,R,S
```

现在，我们统计了一个牛可以接受的名字列表，列表中共有不到 *5000* 个奶牛可以接受的名字。

对于一个拥有某个编号的奶牛，它的编号通过数字与字母的映射，可以得到若干个可能的名字，请你找出这些名字中，奶牛可以接受的名字（即在名字列表中的名字）。

例如，编号 *4734* 可以对应如下 *81* 个可能的名字：

```
GPDG GPDH GPDI GPEG GPEH GPEI GPFG GPFH GPFI GRDG GRDH GRDI
GREG GREH GREI GRFG GRFH GRFI GSDG GSDH GSDI GSEG GSEH GSEI
GSFG GSFH GSFI HPDG HPDH HPDI HPEG HPEH HPEI HPFG HPFH HPFI
HRDG HRDH HRDI HREG HREH HREI HRFG HRFH HRFI HSDG HSDH HSDI
HSEG HSEH HSEI HSFG HSFH HSFI IPDG IPDH IPDI IPEG IPEH IPEI
IPFG IPFH IPFI IRDG IRDH IRDI IREG IREH IREI IRFG IRFH IRFI
ISDG ISDH ISDI ISEG ISEH ISEI ISFG ISFH ISFI
```

我们可以从这些名字中，找到在名字列表中出现过的名字，并字典顺序输出即可。

### 输入格式

第一行包含一个数字编号，编号的可能长度为 *1 ~ 12*。

接下来若干行，每行包含一个由大写字母构成的字符串，表示可接受名字名单中的一个名字。

数据保证，这些名字都是按照字典序排列的。

### 输出格式

输出数字编号对应的所有名字中，在名单中出现过的名字。

注意，按照字典顺序输出，每个名字占一行。

如果没有对应名字，请输出 NONE。

### 数据范围

数据保证名单中的名字不超过5000个。

奶牛编号中不会出现 *1* 和 *0*。

### 输入样例：

```
4734
ABCD
GREG
XYZ
```

### 输出样例：

```
GREG
```
