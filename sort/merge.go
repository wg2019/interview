package sort

// MergeSort 归并排序
func MergeSort(list []int) []int {
	l := len(list)
	if l < 2 {
		return list
	}
	half := l / 2
	return merge(MergeSort(list[0:half]), MergeSort(list[half:]))
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
		} else {
			result = append(result, left[0])
			left = left[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}
