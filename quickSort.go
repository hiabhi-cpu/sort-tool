package main

func quickSort(list []string) {
	quick(list, 0, len(list)-1)
}

func quick(list []string, low, high int) {
	// fmt.Println(high)
	if low < high {
		pivot := partition(list, low, high)
		quick(list, low, pivot-1)
		quick(list, pivot+1, high)
	}
}

func partition(list []string, low, high int) int {
	pivot := list[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if list[j] < pivot {
			i++
			swap(list, i, j)
		}
	}

	swap(list, i+1, high)
	return i + 1
}

func swap(list []string, i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}
