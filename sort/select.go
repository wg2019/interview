package sort

// SelectSort 选择排序
func SelectSort(list []int) {
	l := len(list)
	if l < 2 {
		return
	}

	for i := 0; i < l; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		list[i], list[minIndex] = list[minIndex], list[i]
	}
}
