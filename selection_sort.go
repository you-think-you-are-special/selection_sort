package selection_sort

func intSort(arr []int, order int8) []int {
	length := len(arr)
	min := 0

	for i := 0; i < length-1; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if order == 0 && arr[j] < arr[min] {
				min = j
			}

			if order == 1 && arr[j] > arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

func Asc(arr []int) []int {
	return intSort(arr, 0)
}

func Desc(arr []int) []int {
	return intSort(arr, 1)
}
