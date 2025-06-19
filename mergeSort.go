package main

func mergeSort(list []string) []string {
	if len(list) <= 1 {
		return list
	}
	list1 := mergeSort(list[0 : len(list)/2])

	list2 := mergeSort(list[(len(list) / 2):])
	return merge(list1, list2)
}

func merge(list1, list2 []string) []string {
	resList := make([]string, 0)
	i, j := 0, 0

	for i < len(list1) || j < len(list2) {
		if i == len(list1) {
			resList = append(resList, list2[j])
			j++
		} else if j == len(list2) {
			resList = append(resList, list1[i])
			i++
		} else if list1[i] < list2[j] {
			resList = append(resList, list1[i])
			i++
		} else {
			resList = append(resList, list2[j])
			j++
		}
	}
	// fmt.Println(resList)
	return resList
}
