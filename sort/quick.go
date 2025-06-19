package sort

func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}

	left := 0
	right := len(data) - 1
	index := (left + right) / 2
	pivot := data[index]

	for left < right {
		for data[left] < pivot {
			left++
		}
		for data[right] > pivot {
			right--
		}
		tmp := data[left]
		data[left] = data[right]
		data[right] = tmp
		left++
		right--
	}
	//元データをソート後、pivotの位置が変更する可能性があるので、
	//あらためてpivotの位置を取得して、indexに指定。
	for i, num := range data {
		if num == pivot {
			index = i
		}
	}
	quickSort(data[:index])
	quickSort(data[index+1:])
}
