package oj

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Memset(a []int, v int) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}

// Unique 使用双指针算法去除一个有序切片中的重复值，返回单调有序的切片
func Unique(a []int) []int {
	j := 0
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] != a[i-1] {
			a[j] = a[i]
			j++
		}
	}
	return a[:j]
}
