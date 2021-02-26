package sort

// QuickSort 快速排序
func QuickSort(list []int) {
	partition(list, 0, len(list)-1)
}

func partition(list []int, left, right int) {
	if left >= right {
		return
	}
	tempValue := list[left]
	start := left
	end := right
	for left != right {
		for left < right && list[right] >= tempValue {
			right--
		}
		for left < right && list[left] <= tempValue {
			left++
		}
		if left < right {
			list[left], list[right] = list[right], list[left]
		}
	}
	list[right], list[start] = tempValue, list[left]
	partition(list, start, left)
	partition(list, left+1, end)
}
