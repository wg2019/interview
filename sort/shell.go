package sort

// ShellSort 希尔排序
func ShellSort(list []int) {
	l := len(list)
	for gap := l / 2; gap > 0; gap /= 2 {
		for i := gap; i < l; i++ {
			for j := i - 1; j >= 0; j-- {
				if list[j] < list[j+1] {
					break
				}
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
