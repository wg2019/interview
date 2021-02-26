package sort

// BubbleSort 冒泡排序
func BubbleSort(list []int) {
	l := len(list)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if list[j-1] > list[j] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}
