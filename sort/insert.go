package sort

// InsertSort 插入排序
func InsertSort(list []int) {
	l := len(list)
	for i := 1; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if list[j] < list[j+1] {
				break
			}
			list[j], list[j+1] = list[j+1], list[j]
		}
	}
}
