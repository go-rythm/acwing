package oj

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

// Swap 交换两个int指针指向的值
func Swap(a, b *int) {
	*a, *b = *b, *a
}
