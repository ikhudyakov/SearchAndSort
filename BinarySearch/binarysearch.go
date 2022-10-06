package binarysearch

import (
	"sort"
)

func BinarySearch(arr []int, n int) bool {
	sort.Ints(arr)       // сортируем полученный массив
	first := 0           // индекс первого элемента в массиве
	last := len(arr) - 1 // индекс последнего элемента в массиве

	if n < arr[first] || n > arr[last] { // если искомое число выходит из диапазона чисел, сразу false
		return false
	}
	for first <= last {
		mid := (first + last) / 2 // находим индекс середины массива
		if n == arr[mid] {        // проверяем, если искомое число равно числу, которое находится в массиве по этому индексу, возвращаем true
			return true
		}
		if arr[mid] < n { // проверяем, если искомое число больше, чем число, которое находится в массиве по этому индексу mid
			first = mid + 1 // то в переменную first записываем следующий индекс
		} else if arr[mid] > n { // проверяем, если искомое число меньше, чем число, которое находится в массиве по этому индексу mid
			last = mid - 1 // то в переменную last записываем предыдущий индекс
		}
	}

	return false // если ничего не нашли, возвращаем false
}
