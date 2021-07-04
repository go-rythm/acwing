## [847. 图中点的层次](https://www.acwing.com/problem/content/solution/849/1/)

### 题目

给定一个n个点m条边的有向图，图中可能存在重边和自环。

所有边的长度都是1，点的编号为1~n。

请你求出1号点到n号点的最短距离，如果从1号点无法走到n号点，输出-1。

### 输入格式

第一行包含两个整数n和m。

接下来m行，每行包含两个整数a和b，表示存在一条从a走到b的长度为1的边。

### 输出格式

输出一个整数，表示1号点到n号点的最短距离。

### 数据范围

1≤ n,m ≤10^5

### 输入样例：

```
4 5
1 2
2 3
3 4
1 3
1 4
```

### 输出样例：

```
1
```

### 题解

树的bfs

#### 树与图的遍历

**宽度优先遍历**

```go
const N = 100010

var (
	h, e, ne, d, q [N]int
    st [N]bool
	idx            int
)

// 添加一条边a->b
func add(a int, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

// 初始化
func init() {
    idx = 0;
    for i := 0; i < len(h); i++ {
		h[i] = -1
	}
}

func bfs() {
    hh, tt := 0, 0
    q[0] = 1
    for hh <= tt {
		t := q[hh]
		hh++
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if !st[j] {
				st[j] = true
				tt++
				q[tt] = j
			}
		}
	}
}
```

